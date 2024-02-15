package docxParse

import (
	"fmt"
	"io"
	"os"

	"code.sajari.com/docconv"
)

type Exersize struct {
	num    uint
	text   string
	answer string
}

func rusParse() {

}

func DocxFileParse(path string, sybject string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var r io.Reader
	r = f

	tmpl, _, err := docconv.ConvertDocx(r)
	if err != nil {
		return
	}

	fmt.Println(tmpl)
}
