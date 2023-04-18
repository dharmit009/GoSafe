package main

import (
	"fyne.io/fyne/v2/app"

	"github.com/dharmit009/gopass/gui"
	"github.com/dharmit009/gopass/passutil"
)

const fileName = "./test.enc"

var registered bool = false
var logged bool = false


func main(){

  a := app.New()
  w := a.NewWindow("Gopass")
  keyexist := passutil.CheckMasterKey(fileName)

  if keyexist && !registered{
    // show login 
    w = gui.Login()
    w.ShowAndRun()
  } else if keyexist && !registered{
    // show Dashboard 
    w = gui.Tabs()
    w.ShowAndRun()
  } else {
    // show registration
    w = gui.Registration()
    w.ShowAndRun()
  }

  w.ShowAndRun()
  w.Close()
  a.Quit()

}
