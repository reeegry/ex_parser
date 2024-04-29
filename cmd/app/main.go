package main

import (
	"fmt"

	"github.com/reeegry/ex_parser/internal/categoryParser/jsonDecode"
	"github.com/reeegry/ex_parser/internal/parser"
	"github.com/reeegry/ex_parser/internal/terminalUI"
)

const (
	inf      = "inf"
	hist     = "hist"
	math     = "math"
	mathBase = "mathb"
	chem     = "chem"
	rus      = "rus"
	bio      = "bio"
	eng      = "en"
	geo      = "geo"
	de       = "de"
	soc      = "soc"
	fr       = "fr"
	lit      = "lit"
	sp       = "sp"
	url      = "https://math-ege.sdamgia.ru"
)

func main() {
	p := parser.NewParser()
	decoder := jsonDecode.New("math", "../../internal/categoryParser/jsonData/math.json")
	fmt.Printf("%v", decoder)

	p.PsdamGia.SetUrl("https://math-ege.sdamgia.ru/test?theme=172&print=true")
	p.PsdamGia.GetSdamGiaEx()

	terminalUI.DrawUI(p)

	// p.PDoc.DocxFileParse("../../internal/parser/Pdoc/documents/testod.odt", "")
	// p.CompareExersices()
	//compareExersices(&parsedExSdamGia, &parsedExFromDoc)
}
