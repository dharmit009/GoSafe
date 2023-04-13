package gui

import (
	"fyne.io/fyne/v2/widget"
)

func DelForm() *widget.Form {

	form := widget.NewForm(
		&widget.FormItem{Text: "ViewForm", Widget: widget.NewLabel("")},
	)
	form.Hide() // hide the form initially

	return form
}
