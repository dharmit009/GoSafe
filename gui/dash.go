package gui

import (
	// "strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Dashboard(w fyne.Window) *fyne.Container {

	// ------VVVVVVVVV------------DEL FORM-------VVVVVVVVV---------------------- //

	delform := widget.NewForm(
		&widget.FormItem{Text: "delForm", Widget: widget.NewLabel("")},
	)
	delform.Hide() // hide the form initially
	// Create a rhs container with four buttons

	// ------VVVVVVVVV------------UPDATE FORM-------VVVVVVVVV---------------------- //
	updform := widget.NewForm(
		&widget.FormItem{Text: "updform", Widget: widget.NewLabel("")},
	)
	updform.Hide() // hide the form initially

	// ------------------------------------------------------------------------- //

	addform := AddForm(w)
	viewform := ViewForm()
	addBtn := widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func() {
		addform.Show()
		viewform.Hide()
		delform.Hide()
		updform.Hide()
	})

	viewBtn := widget.NewButtonWithIcon("View", theme.SearchIcon(), func() {
		addform.Hide()
		viewform.Show()
		delform.Hide()
		updform.Hide()
	})

	delBtn := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
		addform.Hide()
		viewform.Hide()
		delform.Show()
		updform.Hide()
	})

	updateBtn := widget.NewButtonWithIcon("Update", theme.UploadIcon(), func() {
		addform.Hide()
		viewform.Hide()
		delform.Hide()
		updform.Show()
	})

	lhsContainer := container.NewVBox(
		addBtn,
		delBtn,
		updateBtn,
		viewBtn,
	)
	// lhsContainer.Resize(fyne.NewSize(0, 100))

	rhsContainer := fyne.NewContainerWithLayout(layout.NewMaxLayout(), addform, delform, updform, viewform)

	// lhsContainer.Layout()
	// Combine the top and bottom canvases in a VBox
	containera := container.NewHBox(
		lhsContainer,
		// form,
		rhsContainer,
	)

	return containera
	// Set the window content to the VBox
	// w.SetContent(containera)
	// w.Resize(fyne.NewSize(700, 400))
}
