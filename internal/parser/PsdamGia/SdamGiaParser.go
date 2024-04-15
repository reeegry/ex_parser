package PsdamGia

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type Answer struct {
	text string
	imgs []string
}

type Exersice struct {
	ProblemId  string
	ExText     string
	imgs_src   []string
	answer_ptr *Answer
}

type PsdamGia struct {
	// Exs  map[string]*Exersice
	Exs  []*Exersice
	url  string
	Subj string
}

func (p *PsdamGia) SetUrl(url string) {
	p.url = url
}

func (p *PsdamGia) GetUrl() string {
	return p.url
}

func NewPsdamGia() *PsdamGia {
	return &PsdamGia{
		// Exs:  make(map[string]*Exersice),
		Exs:  make([]*Exersice, 0),
		url:  "",
		Subj: "",
	}
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

	// exrsices := make(map[string]*Exersice)

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
			ProblemId:  problemId,
			ExText:     txt,
			imgs_src:   imgs,
			answer_ptr: new(Answer),
		}
		// p.Exs[problemId] = exrsice
		p.Exs = append(p.Exs, exrsice)
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
		for _, ex := range p.Exs {
			if ex.ProblemId == id {
				ex.answer_ptr = ansStruct
			}
		}
		// if entry, ok := p.Exs[id]; ok {
		// 	entry.answer_ptr = ansStruct

		// }

	})

	fmt.Println("url:", p.GetUrl())
	err := c.Visit(p.GetUrl())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// p.Exs = exrsices
}
