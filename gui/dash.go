package gui

import (
	// "strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/dharmit009/gopass/watchman"

	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

)

func reloadDB(){
  db := watchman.LoadPasswordDB(dbfile)
  watchman.SavePasswordDB(db, dbfile)
}


// Outputs Dashboard to the screen also works as master form and loads other forms
// based on triggers.
func Dashboard(w fyne.Window) *fyne.Container {

	// ------VVVVVVVVV------------UPDATE FORM-------VVVVVVVVV---------------------- //
	updform := widget.NewForm(
		&widget.FormItem{Text: "updform", Widget: widget.NewLabel("")},
	)
	updform.Hide() // hide the form initially

	// ------------------------------------------------------------------------- //
	addform := AddForm(w)

	viewform, list := ViewForm(w)
	listc := container.NewHScroll(list)
	viewc := container.NewGridWithColumns(2, listc, viewform)

  delform, dlist := DelForm(w)
  listd := container.NewHScroll(dlist)
	delec := container.NewHSplit(listd, delform)

	viewc.Hide()
  delec.Hide()

	addBtn := widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func() {
    reloadDB()

		addform.Show()
		delform.Hide()
		updform.Hide()
		viewform.Hide()

		viewc.Hide()
    delec.Hide()

	})

	viewBtn := widget.NewButtonWithIcon("View", theme.SearchIcon(), func() {
    reloadDB()

		addform.Hide()
		delform.Hide()
		updform.Hide()
		viewform.Show()

		viewc.Show()
    delec.Hide()

	})

	delBtn := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {

    reloadDB()

		addform.Hide()
		delform.Show()
		updform.Hide()
		viewform.Hide()

		viewc.Hide()
    delec.Show()

	})

	updateBtn := widget.NewButtonWithIcon("Update", theme.UploadIcon(), func() {
    reloadDB()

		addform.Hide()
		delform.Hide()
		updform.Show()
		viewform.Hide()

		viewc.Hide()
    delec.Hide()

	})

	lhsContainer := container.NewVBox(
		addBtn,
		delBtn,
		updateBtn,
		viewBtn,
	)
	// lhsContainer.Resize(fyne.NewSize(0, 100))

	rhsContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), addform, delec, updform, viewc)

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
