package gui

import (
	// "strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/dharmit009/gopass/passutil"
)

func Dashboard(w fyne.Window) *fyne.Container{

	// Create a top canvas with an empty form
	// rhsContainer := container.NewHBox()

	form := widget.NewForm(
		&widget.FormItem{Text: "Website Name", Widget: widget.NewEntry()},
		&widget.FormItem{Text: "Username", Widget: widget.NewEntry()},
    &widget.FormItem{Text: "Password", Widget: widget.NewPasswordEntry()},
    // &widget.FormItem{Text: "Strength: ", Widget: widget.NewLabel("")},
    &widget.FormItem{Text: "Strength: ", Widget: widget.NewProgressBar()},
	)

	form.Append("", widget.NewButtonWithIcon("Generate Password", theme.ViewRefreshIcon(), func() {
		password := passutil.GeneratePassword() // generate a 16 character password using passutil package
		form.Items[2].Widget.(*widget.Entry).SetText(password)
	}))

	form.Append("", widget.NewButtonWithIcon("Check Password Strength", theme.ConfirmIcon(), func() {
		temp := float64(passutil.StrengthCheck(form.Items[2].Widget.(*widget.Entry).Text))
		// s := strconv.FormatFloat(float64(temp), 'f', 2, 64)
		// form.Items[3].Widget.(*widget.Label).SetText(s)
    strengthBar := form.Items[3].Widget.(*widget.ProgressBar)
    strengthBar.Min = float64(-3)
    strengthBar.Max = float64(11)
    strengthBar.SetValue(temp)
	}))

	form.Append("", widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
    form.Hide()
	}))

	form.Hide() // hide the form initially

	// Create a rhs container with four buttons
	addBtn := widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func() {
		// show the form when the button is clicked
		form.Show()
	})

	viewBtn := widget.NewButtonWithIcon("View", theme.SearchIcon(), func() {})
	delBtn := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {})
	updateBtn := widget.NewButtonWithIcon("Update", theme.UploadIcon(), func() {})

	lhsContainer := container.NewVBox(
		addBtn,
    delBtn,
    updateBtn,
		viewBtn,
	)
	// lhsContainer.Resize(fyne.NewSize(0, 100)) 

  rhsContainer := fyne.NewContainerWithLayout(layout.NewMaxLayout(), form)

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
