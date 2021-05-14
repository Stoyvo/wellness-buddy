package hydrate

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var Active bool

func fetchDefaultObjs() []fyne.CanvasObject{
	var objs []fyne.CanvasObject

	objs = append(objs, widget.NewLabel("Remember to Hydrate"))

	return objs
}

func Load(app fyne.App, content *fyne.Container) []fyne.CanvasObject {
	objs := fetchDefaultObjs()

	if Active {
		objs = append(objs, widget.NewButton("Done!", func() {
			//add a point to the user and reset the panel
			objs = fetchDefaultObjs()
			content.Objects = objs
			content.Layout.Layout(content.Objects, content.Size())
		}))
		objs = append(objs, widget.NewButton("Skip", func() {
			//dismiss the action by resetting the panel
			objs = fetchDefaultObjs()
			content.Objects = objs
			content.Layout.Layout(content.Objects, content.Size())
		}))
	}

	return objs
}