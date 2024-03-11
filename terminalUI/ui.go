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
		title = "A[red]n[yellow]t[green]i[blue]s[darkmagenta]p[red]i[yellow]z[white]d[:yellow]i[:green]n[:darkcyan]g"
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
	modes := *CreateModes(subjects)
	modeList := tview.NewList()

	RusSelected := func() {
		modeList.Clear()

		for _, value := range modes["Русский язык"] {
			modeList.AddItem(value, "", ' ', func() {})
		}

		app.SetFocus(modeList)
		modeList.SetDoneFunc(func() {
			app.SetFocus(subjectsList)
		}).SetSelectedFunc(func(int, string, string, rune) {
			app.SetFocus(subjectsList)
		})
	}

	MathSelected := func() {
		modeList.Clear()

		for _, value := range modes["Математика профиль"] {
			modeList.AddItem(value, "", ' ', func() {})
		}

		app.SetFocus(modeList)
		modeList.SetDoneFunc(func() {
			app.SetFocus(subjectsList)
		}).SetSelectedFunc(func(int, string, string, rune) {
			app.SetFocus(subjectsList)
		})
	}

	SocSelected := func() {
		modeList.Clear()

		for _, value := range modes["Обществознание"] {
			modeList.AddItem(value, "", ' ', func() {})
		}

		app.SetFocus(modeList)
		modeList.SetDoneFunc(func() {
			app.SetFocus(subjectsList)
		}).SetSelectedFunc(func(int, string, string, rune) {
			app.SetFocus(subjectsList)
		})
	}

	BioSelected := func() {
		modeList.Clear()

		for _, value := range modes["Биология"] {
			modeList.AddItem(value, "", ' ', func() {})
		}

		app.SetFocus(modeList)
		modeList.SetDoneFunc(func() {
			app.SetFocus(subjectsList)
		}).SetSelectedFunc(func(int, string, string, rune) {
			app.SetFocus(subjectsList)
		})
	}

	ChemSelected := func() {
		modeList.Clear()

		for _, value := range modes["Химия"] {
			modeList.AddItem(value, "", ' ', func() {})
		}

		app.SetFocus(modeList)
		modeList.SetDoneFunc(func() {
			app.SetFocus(subjectsList)
		}).SetSelectedFunc(func(int, string, string, rune) {
			app.SetFocus(subjectsList)
		})
	}

	InfoSelected := func() {
		modeList.Clear()

		for _, value := range modes["Информатика"] {
			modeList.AddItem(value, "", ' ', func() {})
		}

		app.SetFocus(modeList)
		modeList.SetDoneFunc(func() {
			app.SetFocus(subjectsList)
		}).SetSelectedFunc(func(int, string, string, rune) {
			app.SetFocus(subjectsList)
		})
	}

	HistSelected := func() {
		modeList.Clear()

		for _, value := range modes["История"] {
			modeList.AddItem(value, "", ' ', func() {})
		}

		app.SetFocus(modeList)
		modeList.SetDoneFunc(func() {
			app.SetFocus(subjectsList)
		}).SetSelectedFunc(func(int, string, string, rune) {
			app.SetFocus(subjectsList)
		})
	}

	EngSelected := func() {
		modeList.Clear()

		for _, value := range modes["Английский язык"] {
			modeList.AddItem(value, "", ' ', func() {})
		}

		app.SetFocus(modeList)
		modeList.SetDoneFunc(func() {
			app.SetFocus(subjectsList)
		}).SetSelectedFunc(func(int, string, string, rune) {
			app.SetFocus(subjectsList)
		})
	}

	PhysSelected := func() {
		modeList.Clear()

		for _, value := range modes["Физика"] {
			modeList.AddItem(value, "", ' ', func() {})
		}

		app.SetFocus(modeList)
		modeList.SetDoneFunc(func() {
			app.SetFocus(subjectsList)
		}).SetSelectedFunc(func(int, string, string, rune) {
			app.SetFocus(subjectsList)
		})
	}

	subjectsList = subjectsList.
		AddItem("Русский язык", "", ' ', RusSelected).
		AddItem("Математика профиль", "", ' ', MathSelected).
		AddItem("Обществознание", "", ' ', SocSelected).
		AddItem("Биология", "", ' ', BioSelected).
		AddItem("Химия", "", ' ', ChemSelected).
		AddItem("Информатика", "", ' ', InfoSelected).
		AddItem("История", "", ' ', HistSelected).
		AddItem("Английский язык", "", ' ', EngSelected).
		AddItem("Физика", "", ' ', PhysSelected)

	subjectsList.SetBorderPadding(1, 1, 2, 2)

	fmt.Println(subjects)
	flex := tview.NewFlex().
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(subjectsList, 10, 1, true).
			AddItem(modeList, 0, 1, false), 0, 1, true)
	if err := app.SetRoot(flex, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}
