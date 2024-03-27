package main

import (
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
	terminalUI.DrawUI()
	p := parser.NewParser()
	p.PsdamGia.GetSdamGiaEx()
	p.PDoc.DocxFileParse("../../internal/parser/Pdoc/documents/1.docx", "")
	p.CompareExersices()
	//compareExersices(&parsedExSdamGia, &parsedExFromDoc)
}
