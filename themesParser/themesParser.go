package themesParser

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func ThemesParser(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ->", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited ->", r.Request.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("div[class=Catalog]", func(h *colly.HTMLElement) {
		fmt.Println(1)
	})

	err := c.Visit("https://math-ege.sdamgia.ru/prob-catalog")
	if err != nil {
		return
	}
}
