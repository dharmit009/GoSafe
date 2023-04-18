package main

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/dharmit009/gopass/ui/jman"
	"github.com/dharmit009/gopass/ui/passutil"
)

const jsfile = "./password.json"

type dropdownItem struct {
	ID      int
	Website string
}

var (
	webee = widget.NewEntry()
	uname = widget.NewEntry()
	passe = widget.NewPasswordEntry()
	mpass = widget.NewPasswordEntry()

	idl    = widget.NewLabel("ID: ")
	webeel = widget.NewLabel("Website : ")
	unamel = widget.NewLabel("Username: ")
	passel = widget.NewLabel("Password: ")

	entryFields = []*widget.Entry{webee, uname, passe, mpass}
	labelFields = []*widget.Label{idl, webeel, unamel, passel}

	// <---------------------- JSON SECTION ------------------------>
	j, _          = jman.NewJman()
	entries, _    = j.GetEntries()
	autoGenButton *widget.Button
	strengthbar   *widget.ProgressBar
)

func main() {

	a := app.New()
	w := a.NewWindow("Password Manager")

	// 	<---------------------- WIDGETS SECTION ------------------------>
	entries = updateEntries(*j)

	items := make([]string, len(entries))
	for i, entry := range entries {
		// Add the ID and website name to the items slice
		items[i] = fmt.Sprintf("%d: %s", entry.ID, entry.Website)
	}

	// Create the dropdown widget using the dropdownItems slice
	dropdown := widget.NewSelect(items, func(selected string) {

		// Parse the selected ID from the dropdown selection
		selectedID, err := strconv.Atoi(strings.Split(selected, ":")[0])
		if err != nil {
			return
		}

		// Find the selected entry by ID and populate the entry details
		for _, entry := range entries {
			if entry.ID == selectedID {

				webeel.SetText("Website: " + entry.Website)
				unamel.SetText("Username: " + entry.Username)
				passel.SetText("Password: " + entry.Password)

				idl.SetText("ID: " + strconv.Itoa(entry.ID))
				webee.SetText(entry.Website)
				uname.SetText(entry.Username)
				passe.SetText(entry.Password)
				break
			}
		}
	})

	autoGenButton = widget.NewButtonWithIcon("Generate Password", theme.ViewRefreshIcon(), autoGen)

	strengthbar = widget.NewProgressBar()
	strengthbar.Min = float64(0)
	strengthbar.Max = float64(10.2)

	webee.SetPlaceHolder("Enter Website Name")
	uname.SetPlaceHolder("Enter Username")
	passe.SetPlaceHolder("Create or Generate New Password")
	mpass.SetPlaceHolder("Enter Master Password")

	passe.OnChanged = func(text string) {
		strength := passutil.StrengthCheck(text)
		strengthbar.SetValue(float64(strength))
	}

	// <----------------------------------- VIEW TAB SECTION ----------------------------------->
	viewTab := container.New(layout.NewVBoxLayout())

	viewTab.Add(dropdown)
	viewTab.Add(idl)
	viewTab.Add(webeel)
	viewTab.Add(unamel)
	viewTab.Add(passel)
	viewTab.Add(widget.NewButtonWithIcon("View", theme.ZoomInIcon(), func() {}))

	// <---------------------- ADD TAB SECTION ------------------------>

	addTab := container.New(layout.NewVBoxLayout())

	addTab.Add(webee)
	addTab.Add(uname)
	addTab.Add(passe)
	addTab.Add(autoGenButton)
	addTab.Add(strengthbar)

	addTab.Add(mpass)

	addTab.Add(widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func() {

		we := webee.Text
		u := uname.Text
		p := passe.Text

		if webee.Text != "" && uname.Text != "" && passe.Text != "" && mpass.Text != "" {
			jman.ShowConfirmationDialog(w, "Add Entry?", "Are you sure you want to Add this entry?", func(response bool) {
				if response {
					err := j.AddEntry(we, u, p)
					resetFields(*j, entryFields, labelFields, dropdown)
					if err != nil {
						fmt.Println(err)
					}
				}
			})
		}

	}))

	// <---------------------- UPDATE TAB SECTION ------------------------>

	updateTab := container.New(layout.NewVBoxLayout())

	updateTab.Add(dropdown)
	updateTab.Add(idl)
	updateTab.Add(webee)
	updateTab.Add(uname)
	updateTab.Add(passe)
	updateTab.Add(strengthbar)
	updateTab.Add(autoGenButton)
	updateTab.Add(mpass)

	updateTab.Add(widget.NewButtonWithIcon("Update", theme.ContentAddIcon(), func() {

		id, err := strconv.Atoi(strings.Split(dropdown.Selected, ":")[0])

		we := webee.Text
		u := uname.Text
		p := passe.Text

		if id > 0 {
			jman.ShowConfirmationDialog(w, "Remove Entry?", "Are you sure you want to delete this entry?", func(response bool) {
				if response {
					err = j.UpdateEntry(id, we, u, p)
					resetFields(*j, entryFields, labelFields, dropdown)
					if err != nil {
						fmt.Println(err)
					}
				}
			})
		}

	}))

	// <---------------------- REMOVE TAB SECTION ------------------------>

	removeTab := container.New(layout.NewVBoxLayout())

	removeTab.Add(dropdown)
	removeTab.Add(idl)
	removeTab.Add(webeel)
	removeTab.Add(unamel)
	removeTab.Add(passel)
	removeTab.Add(widget.NewButtonWithIcon("Remove", theme.ContentRemoveIcon(), func() {

		id, err := strconv.Atoi(strings.Split(dropdown.Selected, ":")[0])
		if id > 0 {
			jman.ShowConfirmationDialog(w, "Remove Entry?", "Are you sure you want to delete this entry?", func(response bool) {
				if response {
					err = j.RemoveEntry(id)
					resetFields(*j, entryFields, labelFields, dropdown)
					if err != nil {
						fmt.Println(err)
					}
				}
			})
		}
	}))

	// <---------------------- GUI MANAGEMENT SECTION ------------------------>

	// Create App Tabs
	tabs := container.NewAppTabs(
		container.NewTabItem("Add", addTab),
		container.NewTabItem("View", viewTab),
		container.NewTabItem("Remove", removeTab),
		container.NewTabItem("Update", updateTab),
	)

	tabs.OnChanged = func(tab *container.TabItem) {
		// switchTab(*j, entries, vitems, vdropdown, dropdown, tabs)
		resetFields(*j, entryFields, labelFields, dropdown)
		refreshList(*j, dropdown)
	}

	tabs.SetTabLocation(container.TabLocationTop)

	// Set Content
	w.SetContent(tabs)

	// Show Window
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}

func autoGen() {
	genpass := passutil.GeneratePassword()
	passe.SetText(genpass)
	strength := passutil.StrengthCheck(passe.Text)
	strengthbar.SetValue(float64(strength))
	passe.Refresh()
	return
}

func resetFields(j jman.Jman, entryFields []*widget.Entry, labelFields []*widget.Label, drop *widget.Select) {
	// Reset text for all Entry widgets
	for _, entryy := range entryFields {
		entryy.SetText("")
	}

	// Reset text for all Label widgets
	strs := []string{"ID: ", "Website: ", "Username: ", "Password: "}
	for i, label := range labelFields {
		label.SetText(strs[i])
	}

	drop.ClearSelected()
	items := refreshList(j, drop)
	drop.Options = items

}

func refreshList(j jman.Jman, dropdown *widget.Select) []string {
	entries = updateEntries(j)

	items := make([]string, len(entries))
	for i, entry := range entries {
		items[i] = fmt.Sprintf("%d: %s", entry.ID, entry.Website)
	}

	return items

}

func updateEntries(j jman.Jman) []jman.Entry {
	err := j.Save()
	if err != nil {
		fmt.Println("Error (Save of Data): ", err)
	}
	err = j.Load()
	if err != nil {
		fmt.Println("Error (Loading of Data): ", err)
	}
	entries, _ = j.GetEntries()

	return entries
}
