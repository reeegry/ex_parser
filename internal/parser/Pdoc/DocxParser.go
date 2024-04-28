package Pdoc

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"

	"code.sajari.com/docconv"
)

type Exersize struct {
	Num    int
	Text   string
	Answer string
}

type PDoc struct {
	// Exs []string
	Exs []*Exersize
}

func NewPDoc() *PDoc {
	return &PDoc{
		Exs: make([]*Exersize, 0),
	}
}

func (p *PDoc) rusParseVariant(text *string) {
	//exercise := make([]string, 0)
	// p.Exs = make([]string, 0)
	indexes := make([][2]int, 0)
	runesText := []rune(*text)
	runesText = append(runesText, rune('$'))
	i := 0
	for i < len(runesText) {
		numIndex := i
		strValue := ""
		for numIndex < len(runesText) && unicode.IsDigit(runesText[numIndex]) {
			strValue += string(runesText[numIndex])
			numIndex++
		}

		if strValue != "" {
			numIndex--
			pair := [2]int{i, numIndex}
			indexes = append(indexes, pair)
		}

		// Поменять на проверку с концом числа
		if i >= 6 && i < len(runesText) {
			endCheck := string(runesText[i-5 : i])
			if strings.ToLower(endCheck) == "ответ" {
				for _, j := range indexes {
					if j[0] < i-6 {

						// p.Exs = append(p.Exs, string(runesText[j[0]:i-6]))
						exNum, _ := strconv.Atoi(string(runesText[j[0] : j[1]+1]))
						p.Exs = append(p.Exs, &Exersize{
							Num:    exNum,
							Text:   string(runesText[j[1]+1 : i-6]),
							Answer: "",
						})
					}
				}
			}
		}

		i = numIndex + 1
	}

}

func (p *PDoc) DocxFileParse(path string, subject string) {

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var r io.Reader
	r = f

	pathSplited := strings.Split(path, "/")
	ext := strings.Split(pathSplited[len(pathSplited)-1], ".")[1]
	fmt.Println(ext)
	fmt.Println(1)

	var tmpl string

	switch ext {
	case "odt":
		tmpl, _, err = docconv.ConvertODT(r)
	case "doc":
		tmpl, _, err = docconv.ConvertDoc(r)
	case "docx":
		tmpl, _, err = docconv.ConvertDocx(r)
	}
	if err != nil {
		panic(err)
	}

	p.rusParseVariant(&tmpl)
}
