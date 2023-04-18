package gui

import (
  "fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

// Used to call Confirm dialog with two options: Confirm and Cancel
func ShowConfirmationDialog(window fyne.Window, title string, message string, callback func(bool)) {
	confirmDialog := dialog.NewConfirm(title, message, func(response bool) {
		callback(response)
	}, window)
	confirmDialog.Show()
}

// Used to display error.
func ShowErrorDialog(window fyne.Window, title string, message string) {
    errorDialog := dialog.NewError(fmt.Errorf(message), window)
    errorDialog.Show()
}


