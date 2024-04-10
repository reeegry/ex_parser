package Pdoc

import (
	"io"
	"os"
	"strings"
	"unicode"

	"code.sajari.com/docconv"
)

type Exersize struct {
	num    uint
	text   string
	answer string
}

type PDoc struct {
	Exs []string
}

func NewPDoc() *PDoc {
	return &PDoc{
		Exs: make([]string, 0),
	}
}

func (p *PDoc) rusParseVariant(text *string) {
	//exercise := make([]string, 0)
	// p.Exs = make([]string, 0)
	indexes := make([][2]int, 0)
	runesText := []rune(*text)
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
						p.Exs = append(p.Exs, string(runesText[j[0]:i-6]))
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

	tmpl, _, err := docconv.ConvertODT(r)
	if err != nil {
		panic(err)
	}

	p.rusParseVariant(&tmpl)
	//fmt.Println(tmpl)
}
