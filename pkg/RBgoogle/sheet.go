package RBgoogle

import (
	"context"
	"os"

	"github.com/RB-PRO/HFLabs/pkg/bases"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

type RBsheet struct {
	SH *spreadsheet.Sheet
}

// Создать экземпляр листа, с которым будет работать скрипт
func NewSheets() (*spreadsheet.Sheet, error) {
	data, err := os.ReadFile("client_secret.json")
	if err != nil {
		return nil, err
	}
	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	if err != nil {
		return nil, err
	}
	client := conf.Client(context.TODO())
	service := spreadsheet.NewServiceWithClient(client)
	spreadsheet, err := service.FetchSpreadsheet("1aDy5lhQV8B1ZRio_HNk02xL8qZ7g_5EqL5Q5cPj-MMU")
	if err != nil {
		return nil, err
	}
	sheet, err := spreadsheet.SheetByIndex(0)
	if err != nil {
		return nil, err
	}
	return sheet, nil
}

// Загрузить информацию
func Push(sheet *spreadsheet.Sheet, datas []bases.HttpHFLabs) error {
	for indexData, valData := range datas {
		sheet.Update(indexData, 0, valData.Code)
		sheet.Update(indexData, 1, valData.Description)
	}

	// Синхронизация с Google Sheets
	err := sheet.Synchronize()
	if err != nil {
		return err
	}
	return err
}

// Загрузить информацию
func Cell(sheet *spreadsheet.Sheet) string {
	/*for indexData, valData := range datas {
		sheet.Update(indexData, 0, valData.Code)
		sheet.Update(indexData, 1, valData.Description)
	}*/
	return sheet.Columns[0][0].Value
}

// Удалить колонки
func DelCol(sheet *spreadsheet.Sheet) error {
	sheet.DeleteColumns(0, 2)

	// Синхронизация с Google Sheets
	err := sheet.Synchronize()
	if err != nil {
		return err
	}
	return err
}
