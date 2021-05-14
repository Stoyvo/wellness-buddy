package hydrate

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var Active bool

func Load(app fyne.App, content *fyne.Container) []fyne.CanvasObject {
	var objs []fyne.CanvasObject

	objs = append(objs, widget.NewLabel("Remember to Hydrate"))

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