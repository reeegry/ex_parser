package parser

import (
	"fmt"
	"strings"

	"github.com/reeegry/ex_parser/internal/levenshteinDistance"
	"github.com/reeegry/ex_parser/internal/parser/Pdoc"
	"github.com/reeegry/ex_parser/internal/parser/PsdamGia"
	"github.com/rivo/tview"
)

type Parser struct {
	PsdamGia *PsdamGia.PsdamGia
	PDoc     *Pdoc.PDoc
	// url       []string
}

func NewParser() *Parser {
	return &Parser{
		PsdamGia: PsdamGia.NewPsdamGia(),
		PDoc:     Pdoc.NewPDoc(),
	}
}

func (p *Parser) CompareExersices(tw *tview.TextView) {
	toReplace := []string{" ", ".", ",", "\n"}
	for _, docProblem := range p.PDoc.Exs {
		for _, sdamGiaProblem := range p.PsdamGia.Exs {

			fmt.Println(docProblem)

			var docTxtReplaced = docProblem.Text
			var sdamGiaProblemReplaced = sdamGiaProblem.ExText

			for _, chr := range toReplace {
				docTxtReplaced = strings.ReplaceAll(docTxtReplaced, chr, "")
				sdamGiaProblemReplaced = strings.ReplaceAll(sdamGiaProblemReplaced, chr, "")
			}

			if levenshteinDistance.FindDistance(&docTxtReplaced, &sdamGiaProblemReplaced) < min(len(docTxtReplaced), len(sdamGiaProblemReplaced))/100*5 {
				// fmt.Printf("Возможен спиздинг\n %s\n", sdamGiaProblem.ExText) // добавить пробелы и нормальное форматирование
				fmt.Fprintf(tw, "[red]Возможен спиздинг: [white][Задача [green]%v[white]] [ID на решу ЕГЭ: [green]%s]\n[white]%s\n\n\n",
					docProblem.Num, sdamGiaProblem.ProblemId, sdamGiaProblem.ExText)

			}
		}
	}
}
