package parser

import (
	"fmt"
	"github.com/reeegry/ex_parser/internal/levenshteinDistance"
	"github.com/reeegry/ex_parser/internal/parser/Pdoc"
	"github.com/reeegry/ex_parser/internal/parser/PsdamGia"
	"strings"
)

type Parser struct {
	PsdamGia *PsdamGia.PsdamGia
	PDoc     *Pdoc.PDoc
	url      []string
}

func NewParser() *Parser {
	return &Parser{
		PsdamGia: PsdamGia.NewPsdamGia(),
		PDoc:     Pdoc.NewPDoc(),
	}
}

func (p *Parser) CompareExersices() {
	toReplace := []string{" ", ".", ",", "\n"}
	for _, sdamGiaProblem := range p.PsdamGia.Exs {
		for _, docProblem := range p.PDoc.Exs {

			for _, chr := range toReplace {
				docProblem = strings.ReplaceAll(docProblem, chr, "")
				sdamGiaProblem.ExText = strings.ReplaceAll(sdamGiaProblem.ExText, chr, "")
			}

			if levenshteinDistance.FindDistance(&docProblem, &sdamGiaProblem.ExText) < min(len(docProblem), len(sdamGiaProblem.ExText))/100*5 {
				fmt.Println("Возможен спиздинг\n")
			}
		}
	}
}
