package app

import (
	"fmt"

	"github.com/RB-PRO/HFLabs/pkg/RBgoogle"
	"github.com/RB-PRO/HFLabs/pkg/parsing"
)

func Run() {

	// Парсим данные
	dataTable := parsing.Parse()

	fmt.Println(dataTable)

	sheetLogin, _ := RBgoogle.NewSheets()

	RBgoogle.Push(sheetLogin, dataTable)

	fmt.Println(RBgoogle.Cell(sheetLogin))

	RBgoogle.DelCol(sheetLogin)
}
