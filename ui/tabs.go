package main

import (
	 "fmt"
  "os"
	"log"
	"io/ioutil"
	"encoding/json"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/dharmit009/gopass/passutil"
	"github.com/dharmit009/gopass/ui/jman"
)

const jsfile = "./passwords.json"

// type PasswordEntry struct {
// 	CredId   string
// 	Website  string
// 	Username string
// 	Password string
// }
//
// type PasswordManager struct {
// 	entries []PasswordEntry
// }

var (
	webee = widget.NewEntry()
	uname = widget.NewEntry()
	passe = widget.NewPasswordEntry()
	mpass = widget.NewPasswordEntry()

	webeel = widget.NewLabel("Website : ")
	unamel = widget.NewLabel("Username: ")
	passel = widget.NewLabel("Password: ")

	autoGenButton *widget.Button
)

func autoGen() {
	genpass := passutil.GeneratePassword()
	passe.SetText(genpass)
	passe.Refresh()
}

func OnSelect(selected string) {
	webeel.SetText("Website : " + selected)
	unamel.SetText("Username: " + selected)
	passel.SetText("Password: " + selected)
}

func main() {
	a := app.New()
	w := a.NewWindow("Password Manager")

	// Create Password Manager
	// pm := PasswordManager{}
	manager := jman.PasswordManager{}

  if _, err := os.Stat("passwords.json"); err == nil {
		data, err := ioutil.ReadFile("passwords.json")
		if err != nil {
			fmt.Println("Error loading passwords:", err)
			return
		}else{
      fmt.Println("pass read !!")
    }

		err = json.Unmarshal(data, &manager)
		if err != nil {
			fmt.Println("Error parsing passwords:", err)
			return
		} else{
      fmt.Println("pass not read !!")
    }
	}

	items := []string{"Item0", "Item1", "Item2", "Item3", "Item4",
		"Item5", "Item6", "Item7", "Item8", "Item9"}

	dropdown := widget.NewSelect(items, OnSelect)
	autoGenButton = widget.NewButtonWithIcon("Generate Password", theme.ViewRefreshIcon(), autoGen)

	webee.SetPlaceHolder("Enter Website Name ")
	uname.SetPlaceHolder("Enter Username")
	passe.SetPlaceHolder("Create or Generate New Password")
	mpass.SetPlaceHolder("Enter Master Password")

	// Create Tabs
	viewTab := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("View Entries"),
	)
	viewTab.Add(dropdown)
	viewTab.Add(webeel)
	viewTab.Add(unamel)
	viewTab.Add(passel)
	viewTab.Add(widget.NewButtonWithIcon("View", theme.ZoomInIcon(), func() {}))

	addTab := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Add Entry"),
	)
	addTab.Add(webee)
	addTab.Add(uname)
	addTab.Add(passe)
	addTab.Add(autoGenButton)
	addTab.Add(mpass)

	addTab.Add(widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func() {

    newEntry, _ := jman.GetNewEntry(webee.Text, uname.Text, passe.Text)
    fmt.Println("New Entry: ", newEntry)
    manager.AddEntry(newEntry)


	}))

	// 	addTab.Add(widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func() {
	// 		web := addTab.Objects[1].(*widget.Entry).Text
	// 		uname := addTab.Objects[3].(*widget.Entry).Text
	// 		var pass string
	// 		if passe.Text == "" {
	// 			pass = addTab.Objects[5].(*widget.Entry).Text
	// 		} else {
	// 			pass = passe.Text
	// 		}

	updateTab := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Update Entry"),
	)

	updateTab.Add(dropdown)
	updateTab.Add(webee)
	updateTab.Add(uname)
	updateTab.Add(passe)
	updateTab.Add(autoGenButton)
	updateTab.Add(mpass)

	updateTab.Add(widget.NewButtonWithIcon("Update", theme.ContentAddIcon(), func() {}))

	removeTab := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Remove Entry"),
	)

	removeTab.Add(dropdown)
	removeTab.Add(webeel)
	removeTab.Add(unamel)
	removeTab.Add(passel)
	removeTab.Add(widget.NewButtonWithIcon("Remove", theme.ContentRemoveIcon(), func() {}))

	// Create App Tabs
	tabs := container.NewAppTabs(
		container.NewTabItem("View", viewTab),
		container.NewTabItem("Add", addTab),
		container.NewTabItem("Remove", removeTab),
		container.NewTabItem("Update", updateTab),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	tabs.OnChanged = func(tab *container.TabItem) {
		go reloadJSONFile(tab, jsfile)
	}

	// Set Content
	w.SetContent(tabs)

	// Show Window
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}

// func (pm *PasswordManager) addEntry(website, username, password, masterPassword string) error {
// 	// Check if master password is correct
// 	if masterPassword != "test" {
// 		return fmt.Errorf("incorrect master password")
// 	}
// 
// 	// Add entry
// 	pm.entries = append(pm.entries, PasswordEntry{
// 		Website:  website,
// 		Username: username,
// 		Password: password,
// 	})
// 
// 	// Save entries to file
// 	if err := pm.saveEntries(); err != nil {
// 		return fmt.Errorf("failed to save entries: %w", err)
// 	}
// 	return nil
// 
// }
// 
// func (pm *PasswordManager) saveEntries() error {
// 	// TODO: Implement saving entries to a file
// 	return nil
// }
// 
// func (pm *PasswordManager) loadEntries() error {
// 	// TODO: Implement loading entries from a file
// 	return nil
// }

func reloadJSONFile(tab *container.TabItem, filename string) {
	if tab == nil || tab.Content == nil {
		return
	}

	// Check if the tab content is a widget that supports updating its content
	if content, ok := tab.Content.(fyne.CanvasObject); ok {
		// Read the JSON file and update the content of the widget
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Println("Error loading JSON file:", err)
			return
		}

		// Assuming that the content widget is a label
		if label, ok := content.(*widget.Label); ok {
			label.SetText(string(data))
		}

		// for _, field := range entryFields {
		// 	field.SetText("")
		// }

	}
}
