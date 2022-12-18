package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func DialogInfo(window *fyne.Window, title string, msg string) dialog.Dialog {
	diaInfo := dialog.NewInformation(title, msg, *window)
	return diaInfo
}
