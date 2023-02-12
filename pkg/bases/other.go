package bases

/*
Этот пакет содержит общую структуру и некоторые дополнительные функции
*/
import (
	"io"
	"os"
)

// Структура данных таблицы, с которой нужно работать.
type HttpHFLabs struct {
	Code        string `db:"code"`        // HTTP-код
	Description string `db:"description"` // Описание
}

// Получение значение из файла
func DataFile(filename string) (string, error) {
	// Открыть файл
	fileToken, errorToken := os.Open(filename)
	if errorToken != nil {
		return "", errorToken
	}

	// Прочитать значение файла
	data := make([]byte, 512)
	n, err := fileToken.Read(data)
	if err == io.EOF { // если конец файла
		return "", errorToken
	}
	fileToken.Close() // Закрытие файла

	return string(data[:n]), nil
}
