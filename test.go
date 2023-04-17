package main

import (
	"encoding/json"
	"io/ioutil"
  "fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Entry struct {
	Website  string `json:"website"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Entries struct {
	Entries []Entry `json:"entries"`
}

var entries []Entry

func main() {
	// Initialize the a and window
	a := app.New()
	window := a.NewWindow("Password Manager")

	// Load the entries from the JSON file
	loadEntries()

	// Create the list of entries
	list := widget.NewList(
		func() int {
			return len(entries)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, item fyne.CanvasObject) {
			label := item.(*widget.Label)
			label.SetText(entries[i].Website)
		},
	)

	// Create the delete button
	deleteBtn := widget.NewButton("Delete", func() {
		if len(list.Selected()) == 0 {
			dialog.ShowInformation("No Selection", "Please select an entry to delete.", window)
			return
		}
		selectedIndex := list.Selected()[0]
		entries = aend(entries[:selectedIndex], entries[selectedIndex+1:]...)
		saveEntries()
		list.Refresh()
	})

	// Create the layout
	content := container.New(layout.NewVBoxLayout(), list, deleteBtn)

	// Set the window content and show it
	window.SetContent(content)
	window.ShowAndRun()
}

func loadEntries() {
	file, err := ioutil.ReadFile("passwords.json")
	if err != nil {
		fmt.Println("Error")
	}

	var data Entries
	if err := json.Unmarshal(file, &data); err != nil {
		fmt.Println("Error")
	}

	entries = data.Entries
}

func saveEntries() {
	data := Entries{Entries: entries}
	file, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Error")
	}

	if err := ioutil.WriteFile("passwords.json", file, 0644); err != nil {
		fmt.Println("Error") 
	}
}
