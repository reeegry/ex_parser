package terminalUI

import (
	"fmt"
	"github.com/rivo/tview"
	"strconv"
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

	fmt.Println(modes)
	return &modes
}

func DrawUI() {
	app := tview.NewApplication()

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
	//x, y, width, height := subjectsList.GetRect()
	//subjectsList.Set
	modes := *CreateModes(subjects)
	//x, y, width, height := subjectsList.GetRect()
	modeList := tview.NewList()
	//modeList.SetRect(x+width+indent, y, width, height)
	SelectedHandler := func(subject string) func() {
		return func() {
			modeList.SetBorder(true)
			subjectsList.SetBorder(false)
			modeList.Clear()

			for _, value := range modes[subject] {
				modeList.AddItem(value, "", 0, func() {})
			}

			app.SetFocus(modeList)
			modeList.SetDoneFunc(func() {
				modeList.SetBorder(false)
				subjectsList.SetBorder(true)
				app.SetFocus(subjectsList)
			}).SetSelectedFunc(func(int, string, string, rune) {
				app.SetFocus(subjectsList)
				modeList.SetBorder(false)
				subjectsList.SetBorder(true)
			})
		}
	}

	subjectsList = subjectsList.
		AddItem("Русский язык", "", 0, SelectedHandler("Русский язык")).
		AddItem("Математика профиль", "", 0, SelectedHandler("Математика")).
		AddItem("Обществознание", "", 0, SelectedHandler("Обществознание")).
		AddItem("Биология", "", 0, SelectedHandler("Биология")).
		AddItem("Химия", "", 0, SelectedHandler("Химия")).
		AddItem("Информатика", "", 0, SelectedHandler("Информатика")).
		AddItem("История", "", 0, SelectedHandler("История")).
		AddItem("Английский язык", "", 0, SelectedHandler("Английский язык")).
		AddItem("Физика", "", 0, SelectedHandler("Физика"))

	subjectsList.SetBorderPadding(1, 1, 2, 2)
	modeList.SetBorderPadding(1, 10, 2, 2)

	fmt.Println(subjects)
	flex := tview.NewFlex().
		AddItem(tview.NewFlex().
			AddItem(subjectsList, 30, 1, true).
			AddItem(modeList, 30, 1, false), 0, 1, true)
	if err := app.SetRoot(flex, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}
