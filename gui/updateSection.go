package gui

import (
	"fyne.io/fyne/v2/widget"
)


// UpdForm() used to generate output for dashboard
func UpdForm() *widget.Form {

	form := widget.NewForm(
		&widget.FormItem{Text: "ViewForm", Widget: widget.NewLabel("")},
	)
	form.Hide() // hide the form initially

	return form
}
