package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

const (
	url = "https://hist-ege.sdamgia.ru" // TODO: прочекать все остальные предметы
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("div.nobreak", func(h *colly.HTMLElement) {
		txt := strings.ReplaceAll(h.Text, "\u00ad", "")
		fmt.Println(txt)
		fmt.Println()
		h.ForEach("img", func(_ int, el *colly.HTMLElement) {
			if el.Attr("class") != "briefcase" {
				img_src := el.Attr("src")
				if img_src[:4] != "http" {
					fmt.Println(url + img_src)
				} else {
					fmt.Println(img_src)
				}
			}

		})
	})

	c.OnHTML("tr.prob_answer", func(h *colly.HTMLElement) {
		h.ForEach("td", func(_ int, el *colly.HTMLElement) {
			if el.Attr("style") != "border:1px solid black;width:25px;text-align:center;padding:2px;" {
				fmt.Println(el.Text)
			}
		})
		h.ForEach("img", func(_ int, el *colly.HTMLElement) {
			fmt.Println(el.Attr("src"))
			fmt.Println(el.Text)
		})
	})

	c.Visit("https://hist-ege.sdamgia.ru/test?theme=42&print=true")
}
