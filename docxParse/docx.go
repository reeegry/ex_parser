package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gonfva/docxlib"
)

var fileLocation *string

func init() {
	fileLocation = flag.String("file", "./14.docx", "file location")
	flag.Parse()
}

type Exersize struct {
	num    uint
	text   string
	answer string
}

func main() {
	readFile, err := os.Open(*fileLocation)
	if err != nil {
		panic(err)
	}
	fileinfo, err := readFile.Stat()
	if err != nil {
		panic(err)
	}
	size := fileinfo.Size()
	doc, err := docxlib.Parse(readFile, int64(size))
	if err != nil {
		panic(err)
	}
	for _, para := range doc.Paragraphs() {
		for _, child := range para.Children() {
			if child.Run != nil {
				fmt.Printf("%s", child.Run.Text.Text)
			}
		}
		fmt.Println()
	}
	fmt.Println("End of main")
}
