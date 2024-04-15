package terminalUI

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/reeegry/ex_parser/internal/parser"
	"github.com/rivo/tview"
)

func CreateModes(subjects []string) *map[string][]string {
	counts := map[string]int{
		"Русский язык":       27,
		"Математика профиль": 19,
		"Обществознание":     25,
		"Биология":           28,
		"Химия":              34,
		"Информатика":        27,
		"Литература":         11,
		"История":            21,
		"Английский язык":    42,
		"Физика":             26,
	}

	modes := make(map[string][]string)

	for _, subj := range subjects {
		modes[subj] = append(modes[subj], "Пробник")
		for i := 1; i <= counts[subj]; i++ {
			modes[subj] = append(modes[subj], strconv.FormatInt(int64(i), 10))
		}
	}

	// fmt.Println(modes)
	return &modes
}

func FindDocs() []string {
	files, _ := filepath.Glob("../../internal/parser/Pdoc/documents/*.odt") // все остальные форматы сделать
	return files
}

func DrawUI(p *parser.Parser) {
	app := tview.NewApplication()

	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	const (
		title  = "A[red]n[yellow]t[green]i[blue]s[darkmagenta]p[red]i[yellow]z[white]d[:yellow]i[:green]n[:darkcyan]g"
		indent = 20
	)

	subjects := []string{
		"Русский язык",
		"Математика профиль",
		"Обществознание",
		"Биология",
		"Химия",
		"Информатика",
		"Литература",
		"История",
		"Английский язык",
		"Физика",
	}

	subjectsList := tview.NewList()
	subjectsList.SetBorder(true)
	modes := *CreateModes(subjects)
	modeList := tview.NewList()
	filesList := tview.NewList()

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().
			AddItem(subjectsList, 0, 2, true).
			AddItem(modeList, 0, 1, false).
			AddItem(filesList, 0, 4, false), 0, 1, true)

	filesListDoneFunc := func() {
		app.SetFocus(subjectsList)
		filesList.SetBorder(false)
		subjectsList.SetBorder(true)
	}

	filesListSelectedFunc := func(int, string, string, rune) {
		app.SetFocus(subjectsList)
		filesList.SetBorder(false)
		subjectsList.SetBorder(true)
		path, _ := filesList.GetItemText(filesList.GetCurrentItem())

		//Disable lists
		app.SetRoot(textView, true).EnableMouse(true)

		p.PDoc.DocxFileParse(path, "")
		p.CompareExersices(textView)

		fmt.Print(path) // <- Путь до файла
	}

	modeListSelectedFunc := func(int, string, string, rune) {
		files := FindDocs()
		filesList.Clear()

		for _, el := range files {
			splitedPath := strings.Split(el, "\\")
			fileName := splitedPath[len(splitedPath)-1]
			filesList.AddItem(fileName, "", 0, func() {})
		}
		app.SetFocus(filesList)
		modeList.SetBorder(false)
		filesList.SetBorder(true)
	}

	modeListDoneFunc := func() {
		modeList.SetBorder(false)
		subjectsList.SetBorder(true)
		app.SetFocus(subjectsList)
	}

	SelectedHandler := func(subject string) func() {
		return func() {
			modeList.SetBorder(true)
			subjectsList.SetBorder(false)
			modeList.Clear()

			for _, value := range modes[subject] {
				modeList.AddItem(value, "", 0, func() {})
			}

			var url_subject string

			switch subject {
			case "Русский язык":
				url_subject = "rus"
			case "Математика профиль":
				url_subject = "math"
			case "Обществознание":
				url_subject = "soc"
			case "Биология":
				url_subject = "bio"
			case "Химия":
				url_subject = "chem"
			case "Информатика":
				url_subject = "info"
			case "Литература":
				url_subject = "lit"
			case "История":
				url_subject = "hist"
			case "Английский язык":
				url_subject = "eng"
			case "Физика":
				url_subject = "phys"
			}

			p.PsdamGia.Subj = url_subject

			app.SetFocus(modeList)
		}
	}

	modeList.SetDoneFunc(modeListDoneFunc).SetSelectedFunc(modeListSelectedFunc)
	filesList.SetSelectedFunc(filesListSelectedFunc).SetDoneFunc(filesListDoneFunc)

	subjectsList = subjectsList.
		AddItem("Русский язык", "", 0, SelectedHandler("Русский язык")).
		AddItem("Математика профиль", "", 0, SelectedHandler("Математика профиль")).
		AddItem("Обществознание", "", 0, SelectedHandler("Обществознание")).
		AddItem("Биология", "", 0, SelectedHandler("Биология")).
		AddItem("Химия", "", 0, SelectedHandler("Химия")).
		AddItem("Информатика", "", 0, SelectedHandler("Информатика")).
		AddItem("История", "", 0, SelectedHandler("История")).
		AddItem("Английский язык", "", 0, SelectedHandler("Английский язык")).
		AddItem("Физика", "", 0, SelectedHandler("Физика"))

	subjectsList.SetBorderPadding(1, 1, 2, 2)
	modeList.SetBorderPadding(1, 10, 2, 2)

	// fmt.Println(subjects)
	if err := app.SetRoot(flex, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}
