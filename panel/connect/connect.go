package connect

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var Active bool
var Joke string

func Load() []fyne.CanvasObject {
	var objs []fyne.CanvasObject

	objs = append(objs, widget.NewLabel("Time to connect with a loved one!"))

	if Joke != "" {
		message := fmt.Sprint("Perhaps start with a joke? Like: \"", Joke, "\"")
		objs = append(objs, widget.NewLabel(message))
	}

	if Active {
		objs = append(objs, widget.NewButton("Done!", func() {
			//add a point to the user and reset the panel
		}))
		objs = append(objs, widget.NewButton("Skip", func() {
			//dismiss the action by resetting the panel
		}))
	}

	return objs
}