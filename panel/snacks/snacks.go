package snacks

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var Active bool
var Snack string

func Load() []fyne.CanvasObject {
	var objs []fyne.CanvasObject

	objs = append(objs, widget.NewLabel("Time for a healthy snack!"))

	if Snack != "" {
		message := fmt.Sprint("How about ", Snack)
		objs = append(objs, widget.NewLabel(message))
	}

	if Active {
		objs = append(objs, widget.NewButton("Done!", func() {
			//add a point to the user and reset the panel
		}))
		objs = append(objs, widget.NewButton("Skip", func() {
			//dismiss the action by reseting the panel
		}))
	}

	return objs
}