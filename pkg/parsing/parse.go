package parsing

import (
	"github.com/RB-PRO/HFLabs/pkg/bases"
	"github.com/gocolly/colly/v2"
)

// Парсинг с помощью gocolly
func Parse() (data []bases.HttpHFLabs) {
	c := colly.NewCollector() // Экземпляр GoColly

	// Срабатываем по правилу
	c.OnHTML("div[class=table-wrap] tbody tr", func(e *colly.HTMLElement) {
		data = append(data, bases.HttpHFLabs{ // Наполняем массив структур данными
			Code:        e.DOM.Find("td[class=confluenceTd]:nth-child(1)").Text(),
			Description: e.DOM.Find("td[class=confluenceTd]:nth-child(2)").Text(),
		})
	})

	c.Visit("https://confluence.hflabs.ru/pages/viewpage.action?pageId=1181220999")
	return data
}
