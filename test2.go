package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

  "github.com/dharmit009/gopass/passutil"
)

func main() {
	a := app.New()
	w := a.NewWindow("Password Manager")

	// Create Password Manager
	pm := &PasswordManager{}

	// Create Tabs
	viewTab := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("View Entries"),
	)

	addTab := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Add Entry"),
	)

	addTab.Add(widget.NewLabel("Website"))
	addTab.Add(widget.NewEntry())
	addTab.Add(widget.NewLabel("Username"))
	addTab.Add(widget.NewEntry())
	addTab.Add(widget.NewLabel("Password"))
	addTab.Add(widget.NewPasswordEntry())

	addTab.Add(widget.NewButtonWithIcon("Generate Password", theme.ViewRefreshIcon(), func() {
		genpass := passutil.GeneratePassword() // generate a 16 character password using passutil package
		addTab.Objects[5].(*widget.Entry).SetText(genpass)
	}))

	addTab.Add(widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func() {
		web := addTab.Objects[1].(*widget.Entry).Text
		uname := addTab.Objects[3].(*widget.Entry).Text
		pass := addTab.Objects[5].(*widget.Entry).Text

		// Ask for master password
		mpasswordEntry := widget.NewPasswordEntry()
		dialog := widget.NewModalPopUp(container.New(layout.NewVBoxLayout(),
			widget.NewLabel("Enter Master Password:"),
			mpasswordEntry,
			widget.NewButton("Submit", func() {
				if err := pm.addEntry(web, uname, pass, mpasswordEntry.Text); err != nil {
					log.Println(err)
					return
				} else{
          widget.NewLabel("Saving new Password")
        }
				w.Close()
			}),
		), w.Canvas())
		dialog.Show()
	}),
	)

	updateTab := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Update Entry"),
	)
	removeTab := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Remove Entry"),
	)

	// Create App Tabs
	tabs := container.NewAppTabs(
		container.NewTabItem("View", viewTab),
		container.NewTabItem("Add", addTab),
		container.NewTabItem("Update", updateTab),
		container.NewTabItem("Remove", removeTab),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// Set Content
	w.SetContent(tabs)

	// Show Window
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}

type PasswordEntry struct {
	Website  string
	Username string
	Password string
}

type PasswordManager struct {
	entries []PasswordEntry
}

func (pm *PasswordManager) addEntry(website, username, password, masterPassword string) error {
	// Check if master password is correct
	if masterPassword != "mypassword" {
		return fmt.Errorf("incorrect master password")
	}

	// Add entry
	pm.entries = append(pm.entries, PasswordEntry{
		Website:  website,
		Username: username,
		Password: password,
	})

	// Save entries to file
	if err := pm.saveEntries(); err != nil {
		return fmt.Errorf("failed to save entries: %v", err)
	}

	return nil
}

func (pm *PasswordManager) saveEntries() error {
	// TODO: Implement save entries to file
	return nil
}
