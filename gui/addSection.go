package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/dharmit009/gopass/passutil"
	"github.com/dharmit009/gopass/watchman"
)

// checks input fields
func fieldChecker(form *widget.Form, numberOfItems int) bool {
	var out bool = true
	for x := 0; x < numberOfItems; x++ {
		if form.Items[x].Widget.(*widget.Entry).Text == "" {
			out = false
		}
	}
	return out
}

func AddForm(w fyne.Window) *widget.Form {
	form := widget.NewForm(
		&widget.FormItem{Text: "Website Name", Widget: widget.NewEntry()},
		&widget.FormItem{Text: "Username", Widget: widget.NewEntry()},
		&widget.FormItem{Text: "Password", Widget: widget.NewPasswordEntry()},
		&widget.FormItem{Text: "Master Password", Widget: widget.NewPasswordEntry()},
		&widget.FormItem{Text: "Strength: ", Widget: widget.NewProgressBar()},
	)

	form.Append("", widget.NewButtonWithIcon("Generate Password", theme.ViewRefreshIcon(), func() {
		password := passutil.GeneratePassword() // generate a 16 character password using passutil package
		form.Items[2].Widget.(*widget.Entry).SetText(password)
	}))

	form.Append("", widget.NewButtonWithIcon("Check Password Strength", theme.ConfirmIcon(), func() {
		temp := float64(passutil.StrengthCheck(form.Items[2].Widget.(*widget.Entry).Text))
		strengthBar := form.Items[4].Widget.(*widget.ProgressBar)
		strengthBar.Min = float64(-3)
		strengthBar.Max = float64(11)
		strengthBar.SetValue(temp)
	}))

	form.Append("", widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		if fieldChecker(form, 4) == false {
			ShowErrorDialog(w, "Error", "Field is empty!")
		} else {
			pass1 := form.Items[3].Widget.(*widget.Entry).Text
			if out := watchman.CheckPassEqualToMP(pass1); out == true {
				ShowConfirmationDialog(w, "Confirm Action", "Are you sure you want to do this?", func(response bool) {
					if response {
						err := watchman.AddEntry(
							form.Items[0].Widget.(*widget.Entry).Text,
							form.Items[1].Widget.(*widget.Entry).Text,
							form.Items[2].Widget.(*widget.Entry).Text,
							form.Items[3].Widget.(*widget.Entry).Text,
						)
						if err != nil {
							dialog.NewError(err, w)
						}
					}
				})
			} else {
				ShowErrorDialog(w, "Error", "Error: Incorrect Master Password!")
			}
		}
	}))

	form.Hide() // hide the form initially
	return form
}
