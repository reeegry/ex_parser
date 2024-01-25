package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

const (
	inf  = "inf"
	hist = "his"
	url  = "https://*-ege.sdamgia.ru" // TODO: прочекать все остальные предметы
)

// func subjCheck(url string) string {
// 	var url_copy string = url
// 	fmt.Println(url_copy)

// 	return url_copy
// }

type Answer struct {
	text string
	imgs []string
}

type Exersice struct {
	exText     string
	imgs_src   []string
	answer_ptr *Answer
}

func main() {
	c := colly.NewCollector()
	exrsices := make(map[string]*Exersice)

	c.OnHTML("div.nobreak", func(h *colly.HTMLElement) {
		txt := strings.ReplaceAll(h.Text, "\u00ad", "")
		problem_id := h.DOM.Find("a").Text()

		imgs := []string{}
		h.ForEach("img", func(_ int, el *colly.HTMLElement) {
			if el.Attr("class") != "briefcase" {
				img_src := el.Attr("src")
				if img_src[:4] != "http" {
					img_src = url + img_src
				}
				imgs = append(imgs, img_src)
			}

		})
		exrsice := &Exersice{
			exText:     txt,
			imgs_src:   imgs,
			answer_ptr: new(Answer),
		}
		exrsices[problem_id] = exrsice
	})

	c.OnHTML("tr.prob_answer", func(h *colly.HTMLElement) {
		var ans_struct *Answer = new(Answer)
		var id string

		h.ForEach("td", func(_ int, el *colly.HTMLElement) {
			if el.Attr("style") == "border:1px solid black;text-align:left;padding:2px;text-indent:0" { // ans
				ans_struct.text = el.Text
			}
			id_selection := el.DOM.Find("a")
			_, ok := id_selection.Attr("href")
			if ok {
				id = id_selection.Text()
			}
		})

		ans_imgs := []string{}
		h.ForEach("img", func(_ int, el *colly.HTMLElement) {
			ans_imgs = append(ans_imgs, el.Attr("src"))
		})
		ans_struct.imgs = ans_imgs
		if entry, ok := exrsices[id]; ok {
			entry.answer_ptr = ans_struct

		}
		fmt.Printf("%p\n", ans_struct)

	})

	c.Visit("https://ege.sdamgia.ru/test?theme=205&print=true")

	for id, ex := range exrsices {
		for i := 0; i < 100; i++ {
			fmt.Printf("-")
		}
		fmt.Println()
		fmt.Println(id)
		fmt.Println(ex.exText)
		fmt.Println(ex.imgs_src)
		fmt.Println(ex.answer_ptr)
		fmt.Printf("%p\n", ex.answer_ptr)
	}

}
