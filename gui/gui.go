package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/dharmit009/gopass/watchman"
)

const (
	fileName    = "./test.enc"
	windowTitle = "Gopass"
)

var registered bool

func Register(w fyne.Window) {
	passwordEntry := widget.NewPasswordEntry()
	passwordConfirmEntry := widget.NewPasswordEntry()
	passwordEntry.PlaceHolder = "Enter your password"
	passwordConfirmEntry.PlaceHolder = "Re-enter your password"

	form := widget.NewForm()
	form.Append("", passwordEntry)
	form.Append("", passwordConfirmEntry)
	form.Append("", widget.NewButton("Create Master Key", func() {
		flag := watchman.CheckPassEqual(passwordEntry.Text, passwordConfirmEntry.Text)
		if !flag {
			passwordEntry.SetText("")
			passwordConfirmEntry.SetText("")
			w.SetContent(widget.NewLabel("Passwords are not the same. Try Again!"))
			return
		}
		if watchman.CreateMasterKey(fileName, passwordConfirmEntry.Text, passwordEntry.Text) {
			registered = true
			form.Hide()
			showLoginForm(w)
		} else {
			passwordEntry.SetText("")
			passwordConfirmEntry.SetText("")
			w.SetContent(widget.NewLabel("Registration Error"))
		}
	}))

	container := fyne.NewContainerWithLayout(layout.NewMaxLayout(), form)
	w.SetContent(container)
}

func Login(w fyne.Window) {
	passwordConfirmEntry := widget.NewPasswordEntry()
	passwordConfirmEntry.PlaceHolder = "Enter Your Master Password"

	form := widget.NewForm()
	form.Append("", passwordConfirmEntry)
	form.Append("", widget.NewButton("Login", func() {
		comparer := watchman.CheckPassEqualToMP(passwordConfirmEntry.Text)
		if comparer {
			w.SetContent(widget.NewLabel("Login Successful!"))
		} else {
			w.SetContent(widget.NewLabel("Login Unsuccessful"))
		}
	}))

	container := fyne.NewContainerWithLayout(layout.NewMaxLayout(), form)
	w.SetContent(container)
}

func showLoginForm(w fyne.Window) {
	Login(w)
}

func RunGUI() {
	a := app.New()
	w := a.NewWindow(windowTitle)

	if success := watchman.CheckMasterKey(fileName); success && !registered {
		showLoginForm(w)
	} else {
		Register(w)
	}

	w.ShowAndRun()
	defer a.Quit()
}
