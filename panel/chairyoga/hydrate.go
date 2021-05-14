package chairyoga

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var Active bool

func Load() []fyne.CanvasObject {
	var objs []fyne.CanvasObject

	objs = append(objs, widget.NewLabel("Chair Yoga"))

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