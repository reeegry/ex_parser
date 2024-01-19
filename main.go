package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func printLines() {
	for i := 0; i < 50; i++ {
		fmt.Print("-")
	}

	fmt.Println()
}

func main() {
	c := colly.NewCollector()

	c.OnHTML("p.left_margin", func(h *colly.HTMLElement) {
		txt := strings.ReplaceAll(h.Text, "\u00ad", "")
		fmt.Println(txt)
		fmt.Println()
		selection := h.DOM
		fmt.Println(selection.Find("i").Text())
		h.ForEach("img", func(_ int, el *colly.HTMLElement) {
			fmt.Println(el.Attr("src"))

		})
	})

	c.Visit("https://ege.sdamgia.ru/test?theme=79&print=true")
}
