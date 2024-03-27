package docxParse

import (
	"code.sajari.com/docconv"
	"io"
	"os"
	"strings"
	"unicode"
)

type Exersize struct {
	num    uint
	text   string
	answer string
}

func rusParseVariant(text *string) *[]string {
	exercise := make([]string, 0)
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
						//fmt.Println(j, i-6, string(runesText[j[0]:j[1]+1]))
						exercise = append(exercise, string(runesText[j[0]:i-6]))
					}
				}
			}
		}

		i = numIndex + 1
		//fmt.Println(i)
	}

	return &exercise
}

func DocxFileParse(path string, subject string) *[]string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var r io.Reader
	r = f

	tmpl, _, err := docconv.ConvertDocx(r)
	if err != nil {
		panic(err)
	}

	return rusParseVariant(&tmpl)
	//fmt.Println(tmpl)
}
