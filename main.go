package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dharmit009/gopass/passutil"
)

const (
	windowTitle = "Password Manager"
	fileName    = "./test.enc"
)

func main() {
	a := app.New()
	w := a.NewWindow(windowTitle)

	// Check if the encrypted file exists
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// If the file doesn't exist, prompt the user to create a new password
		fmt.Println("Master password has not been set")
		passwordEntry := widget.NewPasswordEntry()
		passwordConfirmEntry := widget.NewPasswordEntry()

		form := &widget.Form{
			Items: []*widget.FormItem{
				{Text: "Enter a new password", Widget: passwordEntry},
				{Text: "Confirm password", Widget: passwordConfirmEntry},
			},
			OnSubmit: func() {
				// When the user submits the form, check if the passwords match
				if passwordEntry.Text != passwordConfirmEntry.Text {
					// If the passwords don't match, show an error message and reset the form
					widget.NewLabel("Passwords do not match. Please try again.").Show()
					passwordEntry.SetText("")
					passwordConfirmEntry.SetText("")
					return
				}

				// If the passwords match, encrypt the password and save it to a file
				if err := passutil.EncryptFile(fileName, []byte(passwordConfirmEntry.Text), passwordEntry.Text); err != nil {
					// If there was an error encrypting the password, show an error message and reset the form
					widget.NewLabel(fmt.Sprintf("Error saving password: %v", err)).Show()
					passwordEntry.SetText("")
					passwordConfirmEntry.SetText("")
					return
				}
				// If the password was encrypted successfully, close the form
				//				form.Hide()
			},
		}
		// Set the form as the content of the window
		w.SetContent(container.NewCenter(form))
	} else {
		// If the file exists, then take login form
		passwordConfirmEntry := widget.NewPasswordEntry()

		lform := &widget.Form{
			Items: []*widget.FormItem{
				{Text: "Enter a your master password", Widget: passwordConfirmEntry},
			},
			OnSubmit: func() {
				// When the user submits the form, check if the passwords match
				comparer, err := passutil.DecryptFile(fileName, passwordConfirmEntry.Text)
				if err != nil {
					fmt.Println("Error: masterpassword file has been corrupted !!")
				}
				if passwordConfirmEntry.Text == string(comparer) {
					fmt.Println("Entered Password: ", passwordConfirmEntry.Text)
					fmt.Println("Decrypted Password: ", string(comparer))
					w.SetContent(widget.NewLabel("LOGIN SUCCESSFUL!!"))
					return
				}

				// If the password was encrypted successfully, close the form
				// form.Hide()
			},
		}
		// Set the form as the content of the window
		w.SetContent(container.NewCenter(lform))
	}

	// Show the window
	w.ShowAndRun()
}
