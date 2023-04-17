package gui

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"encoding/json"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/dharmit009/gopass/passutil"
	"github.com/dharmit009/gopass/watchman"
)

// It is used to trigger view form ...
func ViewForm(w fyne.Window) (*widget.Form, *widget.List) {

	websiteEntry := widget.NewLabel("")
	usernameEntry := widget.NewLabel("")
	passwordEntry := widget.NewLabel("")

	copyUsernameButton := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		copyToClipboard(usernameEntry.Text)
	})

	copyPasswordButton := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		copyToClipboard(passwordEntry.Text)
	})

	var passwordEntries watchman.PasswordDB

	data, err := ioutil.ReadFile("passwords.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	err = json.Unmarshal(data, &passwordEntries)
	if err != nil {
		fmt.Println("Error unmarshalling data:", err)
	}

	list := widget.NewList(
		func() int { return len(passwordEntries.Entries) },
		func() fyne.CanvasObject { return widget.NewButton("", func() {}) },
		func(lli widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Button).SetText(passwordEntries.Entries[lli].Website + " - [ " + passwordEntries.Entries[lli].Username + " ] ")
			co.(*widget.Button).OnTapped = func() {
				websiteEntry.SetText(passwordEntries.Entries[lli].Website)
				usernameEntry.SetText(passwordEntries.Entries[lli].Username)
				mpassEntry := widget.NewPasswordEntry()
				passwordDialog := dialog.NewCustom("Enter Master Password", "Confirm", mpassEntry, w)
				passwordDialog.SetOnClosed(func() {
					if watchman.CheckPassEqualToMP(mpassEntry.Text) {
						decpasswd, _ := passutil.Decrypt(passwordEntries.Entries[lli].Password, "testtest")
						passwordEntry.SetText(string(decpasswd))
					} else {
						ShowErrorDialog(w, "Error: ", "Incorrect Master Password")
					}
				})
        passwordDialog.Show()
			}
		},
	)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Website: ", Widget: websiteEntry},
			{Text: "Username: ", Widget: container.New(layout.NewHBoxLayout(), usernameEntry, copyUsernameButton)},
			{Text: "Password: ", Widget: container.New(layout.NewHBoxLayout(), passwordEntry, copyPasswordButton)},
		},
	}
	form.SubmitText = ""

	form.Hide() // hide the form initially

	return form, list
}

func copyToClipboard(str string) {
	// Get the clipboard
	clip := os.Getenv("WAYLAND_DISPLAY")
	if clip == "" {
		clip = "X11"
	}

	switch strings.ToUpper(clip) {
	case "WAYLAND":
		// TODO: Implement copying to clipboard in Wayland
		fmt.Println("Copying to clipboard is not supported in Wayland")
	case "X11":
		// Copy to clipboard in X11
		cmd := exec.Command("xclip", "-selection", "clipboard")
		cmd.Stdin = strings.NewReader(str)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error copying to clipboard:", err)
		}
	default:
		fmt.Println("Unknown clipboard type:", clip)
	}
}
