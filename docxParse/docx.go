package docxParse

import (
	"code.sajari.com/docconv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Exersize struct {
	num    uint
	text   string
	answer string
}

func rusParseVariant(text *string) *[]string {
	exercise := make([]string, 0)
	indexes := make([]int, 0)
	runesText := []rune(*text)
	for i, char := range runesText {
		value, _ := strconv.Atoi(string(char))

		if 1 <= value && value <= 9 {
			indexes = append(indexes, i)
		}

		if i >= 6 && i < len(runesText) {
			endCheck := string(runesText[i-5 : i])
			if strings.ToLower(endCheck) == "ответ" {
				for _, j := range indexes {
					if j < i-6 {
						fmt.Println(j, i-6)
						exercise = append(exercise, string(runesText[j:i-6]))
					}
				}
			}
		}

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
