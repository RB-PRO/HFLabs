package app

import (
	"io"
	"log"
	"os"

	"github.com/RB-PRO/HFLabs/pkg/RBgoogle"
	"github.com/RB-PRO/HFLabs/pkg/parsing"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Run() {

	spreadsheetID := "1aDy5lhQV8B1ZRio_HNk02xL8qZ7g_5EqL5Q5cPj-MMU" // ID книги
	gid := "1303466946"                                             // ID листа

	// Получить токен для телеграм бота
	tokenTelegram, errorFile := dataFile("telegramToken")
	if errorFile != nil {
		log.Fatalln(errorFile)
	}

	// Подключаемся к Google Sheet
	sheetLogin, sheetLoginError := RBgoogle.NewSheets(spreadsheetID)
	if sheetLoginError != nil {
		log.Fatalln(sheetLoginError)
	}

	// Запускаем бота
	bot, errNewBot := tgbotapi.NewBotAPI((tokenTelegram))
	if errNewBot != nil {
		log.Fatalln(errNewBot)
	}
	bot.Debug = false
	log.Printf("Авторизовался %s", bot.Self.UserName)

	// Настройка бота
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// Опрашиваем обновления в боте
	for update := range updates {
		// Игнорирование НЕсообщения
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		// Игнорирование НЕкоманды
		if !update.Message.IsCommand() {
			continue
		}

		// Собираем сообщение для отправки
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Слушаем только команды
		switch update.Message.Command() {
		case "start":
			msg.Text = "Привет. Я бот, который обновит данные в таблице\n" +
				"https://docs.google.com/spreadsheets/d/" + spreadsheetID + "/edit#gid=" + gid + "\n" +
				"Отправь мне команду /update и будет магия :)"
		case "update":
			// Парсим данные
			dataTable, dataTableErr := parsing.Parse()
			if dataTableErr != nil {
				msg.Text = dataTableErr.Error()
				break
			}

			// Отчистка Колонок
			errorDel := sheetLogin.DelCol()
			if errorDel != nil {
				msg.Text = errorDel.Error()
				break
			}

			// Загружаем данные в Google Sheet
			errorPush := sheetLogin.Push(dataTable)
			if errorPush != nil {
				msg.Text = errorPush.Error()
				break
			}

			// Объединить сообщение вывода
			var dataMessage string
			for _, valueData := range dataTable {
				dataMessage += valueData.Code + " - " + valueData.Description + "\n"
			}

			msg.Text = "Я загрузил\n" +
				dataMessage +
				"\nв таблицу\n" +
				"https://docs.google.com/spreadsheets/d/" + spreadsheetID + "/edit#gid=" + gid
		default:
			msg.Text = "Попробуй команду /start"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

// Получение значение из файла
func dataFile(filename string) (string, error) {
	// Открыть файл
	fileToken, errorToken := os.Open(filename)
	if errorToken != nil {
		return "", errorToken
	}

	// Прочитать значение файла
	data := make([]byte, 64)
	n, err := fileToken.Read(data)
	if err == io.EOF { // если конец файла
		return "", errorToken
	}
	fileToken.Close() // Закрытие файла

	return string(data[:n]), nil
}
