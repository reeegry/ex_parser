package unloadDoc

import (
	"fmt"
	"os"
)

func Upload(parsedExPtr *[]string) {
	file, err := os.Create("parsed.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	parsedEx := *parsedExPtr
	for _, exText := range parsedEx {
		text := exText
		file.WriteString(text + "\n\n")
	}
	fmt.Println("Done.")
}
