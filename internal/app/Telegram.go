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

	RBgoogle.Push(dataTable)

}
