package PsdamGia

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type Answer struct {
	text string
	imgs []string
}

type Exersice struct {
	ExText     string
	imgs_src   []string
	answer_ptr *Answer
}

type PsdamGia struct {
	Exs map[string]*Exersice
	url string
}

func NewPsdamGia() *PsdamGia {
	return &PsdamGia{}
}

func (p *PsdamGia) ExPrint() {
	for id, ex := range p.Exs {
		for i := 0; i < 100; i++ {
			fmt.Printf("-")
		}
		fmt.Println()
		fmt.Println("ID: ", id)
		fmt.Println("EX TEXT: ", ex.ExText)
		fmt.Println("EX IMGS_SRC: ", ex.imgs_src)
		fmt.Println("ANSWER TEXT: ", ex.answer_ptr.text)
		fmt.Println("ANSWER IMGS: ", ex.answer_ptr.imgs)
	}
}

func (p *PsdamGia) GetSdamGiaEx() {
	c := colly.NewCollector()
	exrsices := make(map[string]*Exersice)

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnHTML("div.nobreak", func(h *colly.HTMLElement) {
		txt := h.DOM.Find("p").Text()
		txt = strings.ReplaceAll(txt, "\u00ad", "")
		problemId := h.DOM.Find("span.prob_nums").Find("a").Text()

		imgs := []string{}
		h.ForEach("img", func(_ int, el *colly.HTMLElement) {
			if el.Attr("class") != "briefcase" {
				img_src := el.Attr("src")
				if img_src[:4] != "http" {
					img_src = p.url + img_src
				}
				imgs = append(imgs, img_src)
			}

		})
		exrsice := &Exersice{
			ExText:     txt,
			imgs_src:   imgs,
			answer_ptr: new(Answer),
		}
		exrsices[problemId] = exrsice
	})

	c.OnHTML("tr.prob_answer", func(h *colly.HTMLElement) {
		var ansStruct *Answer = new(Answer)
		var id string

		h.ForEach("td", func(_ int, el *colly.HTMLElement) {
			if el.Attr("style") == "border:1px solid black;text-align:left;padding:2px;text-indent:0" { // ans
				ansStruct.text = el.Text
			}
			idSelection := el.DOM.Find("a")
			_, ok := idSelection.Attr("href")
			if ok {
				id = idSelection.Text()
			}
		})

		ansImgs := []string{}
		h.ForEach("img", func(_ int, el *colly.HTMLElement) {
			ansImgs = append(ansImgs, el.Attr("src"))
		})
		ansStruct.imgs = ansImgs
		if entry, ok := exrsices[id]; ok {
			entry.answer_ptr = ansStruct

		}

	})

	err := c.Visit("https://math-ege.sdamgia.ru/test?theme=182&print=true")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	p.Exs = exrsices
}
