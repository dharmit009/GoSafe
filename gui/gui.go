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

var registered bool = false
var logged bool = false

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
			w.SetContent(widget.NewLabel("Master Password must be more than 8 characters"))
		}
	}))

	container := fyne.NewContainerWithLayout(layout.NewMaxLayout(), form)
	w.SetContent(container)
	w.Resize(fyne.NewSize(700, 400))
	w.Close()
}

func Login(w fyne.Window) {
	passwordConfirmEntry := widget.NewPasswordEntry()
	passwordConfirmEntry.PlaceHolder = "Enter Your Master Password"

	form := widget.NewForm()
	form.Append("", passwordConfirmEntry)
	form.Append("", widget.NewButton("Login", func() {
		if comparer := watchman.CheckPassEqualToMP(passwordConfirmEntry.Text); comparer == true {
			logged = true
			container := Dashboard(w)
			w.SetContent(container)
		} else {
			w.SetContent(widget.NewLabel("Login Unsuccessful"))
		}
	}))

	container := fyne.NewContainerWithLayout(layout.NewMaxLayout(), form)
	w.SetContent(container)
	w.Resize(fyne.NewSize(700, 400))
}

func showLoginForm(w fyne.Window) {
	Login(w)
}

func RunGUI() {
	a := app.New()
	w := a.NewWindow(windowTitle)
	keyexist := watchman.CheckMasterKey(fileName)

	if keyexist && !registered {
		showLoginForm(w)
	} else if keyexist && logged {
		Dashboard(w)
	} else {
		Register(w)
	}

	w.ShowAndRun()
	w.Close()
	w.Resize(fyne.NewSize(700, 400))
	a.Quit()
}
