package gui

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/dharmit009/gopass/watchman"
)

// DelForm() function is used to trigger delete item form
func DelForm(w fyne.Window) (*widget.Form, *widget.List) {
	var passwordEntries watchman.PasswordDB
  passwordEntries = watchman.LoadPasswordDB("./password.json")

	data, err := ioutil.ReadFile("passwords.json")
	if err != nil {
		dialog.ShowError(fmt.Errorf("Error: While reading the json file!"), w)
	}

	err = json.Unmarshal(data, &passwordEntries)
	if err != nil {
		dialog.ShowError(fmt.Errorf("Error: While Unmarshalling the file!"), w)
	}

	var selectedEntry *watchman.PasswordEntry

	list := widget.NewList(
		func() int {
			return len(passwordEntries.Entries)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			e := passwordEntries.Entries[id]
			item.(*widget.Label).SetText(e.Website)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		selectedEntry = &passwordEntries.Entries[id]
	}

	delButton := widget.NewButton("Delete", func() {
		if selectedEntry == nil {
			return
		}
		var newEntries []watchman.PasswordEntry
		for _, e := range passwordEntries.Entries {
			if e.Website != selectedEntry.Website {
				newEntries = append(newEntries, e)
			}
		}
		passwordEntries.Entries = newEntries
		watchman.SavePasswordDB(passwordEntries, "passwords.json")
		selectedEntry = nil
		list.Refresh()
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Website: ", Widget: list},
			{Text: "Delete: ", Widget: delButton},
		},
	}

	return form, list
}
