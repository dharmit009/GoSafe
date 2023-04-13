package gui

import (
  "fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func ShowConfirmationDialog(window fyne.Window, title string, message string, callback func(bool)) {
	confirmDialog := dialog.NewConfirm(title, message, func(response bool) {
		callback(response)
	}, window)
	confirmDialog.Show()
}

func ShowErrorDialog(window fyne.Window, title string, message string) {
    errorDialog := dialog.NewError(fmt.Errorf(message), window)
    errorDialog.Show()
}


