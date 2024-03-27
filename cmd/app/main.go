package main

import (
	"github.com/reeegry/ex_parser/internal/parser"
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
	//terminalUI.DrawUI()
	//parsedExSdamGia := SdamGiaParse()
	//ExPrint(&parsedExSdamGia)
	//var parsedExFromDoc []string
	//parsedExFromDoc = *docxParse.DocxFileParse("./docxParse/1.docx", "")
	//fmt.Println(parsedExFromDoc)
	//unloadDoc.Upload(&parsedExFromDoc)

	p := parser.NewParser()
	p.PsdamGia.GetSdamGiaEx()
	p.PDoc.DocxFileParse("../../internal/parser/Pdoc/documents/1.docx", "")
	p.CompareExersices()
	//compareExersices(&parsedExSdamGia, &parsedExFromDoc)
}
