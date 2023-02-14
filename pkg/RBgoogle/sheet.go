package RBgoogle

import (
	"context"
	"os"

	"github.com/RB-PRO/HFLabs/pkg/bases"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

// Структура для работы с Google Sheet
type RBsheet struct {
	SH *spreadsheet.Sheet
}

// Создать экземпляр листа, с которым будет работать скрипт
func NewSheets() (RBsheet, error) {

	// Загружаем секретный ключ
	data, err := os.ReadFile("client_secret.json")
	if err != nil {
		return RBsheet{}, err
	}

	// Парсим ключ в библиотеку"gopkg.in/Iwark/spreadsheet.v2"
	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	if err != nil {
		return RBsheet{}, err
	}
	client := conf.Client(context.TODO())
	service := spreadsheet.NewServiceWithClient(client)

	// Получаем экземпляр книги
	spreadsheet, err := service.FetchSpreadsheet("1aDy5lhQV8B1ZRio_HNk02xL8qZ7g_5EqL5Q5cPj-MMU")
	if err != nil {
		return RBsheet{}, err
	}

	// Получаем экземпляр листа
	sheet, err := spreadsheet.SheetByIndex(0)
	if err != nil {
		return RBsheet{}, err
	}
	return RBsheet{sheet}, nil
}

// Загрузить информацию
func (sheet RBsheet) Push(datas []bases.HttpHFLabs) error {
	for indexData, valData := range datas {
		sheet.SH.Update(indexData, 0, valData.Code)
		sheet.SH.Update(indexData, 1, valData.Description)
	}

	// Синхронизация с Google Sheets
	err := sheet.SH.Synchronize()
	if err != nil {
		return err
	}
	return err
}

// Получить значение ячейки
func (sheet RBsheet) Cell(row, col int) string {
	return sheet.SH.Columns[row][col].Value
}

// Удалить колонки
func (sheet RBsheet) DelCol() error {
	sheet.SH.DeleteColumns(0, 2)

	// Синхронизация с Google Sheets
	err := sheet.SH.Synchronize()
	if err != nil {
		return err
	}
	return err
}
