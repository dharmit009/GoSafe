package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/dharmit009/gopass/gui"
	"github.com/dharmit009/gopass/passutil"
)

const fileName = "master.enc"

var registered bool = false
var logged bool = false


func showTabs(w fyne.Window) *fyne.Container{
    container := gui.Tabs(w)
    return container
}

func showLogin(w fyne.Window) *fyne.Container {
  container, flag := gui.Login(w)
  logged = flag
  return container
}

func showReg(w fyne.Window) *fyne.Container{
  container, flag := gui.Registration(w)
  registered = flag
  return container 
}


func main() {

	myapp := app.NewWithID("0")
	w := myapp.NewWindow("Gopass")
  var container *fyne.Container

	keyexist := passutil.CheckMasterKey(fileName)

	if keyexist && !registered {
    container = showLogin(w)
	} else if keyexist && logged {
    container = showTabs(w)
	} else {
    container = showReg(w)
	}

  w.SetContent(container)
	w.Resize(fyne.NewSize(400, 300))
	w.SetFixedSize(true)
  w.Show()
  myapp.Run()

}
