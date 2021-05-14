package snacks

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var Active bool
var Snack string

func fetchDefaultObjs() []fyne.CanvasObject{
	var objs []fyne.CanvasObject

	objs = append(objs, widget.NewLabel("Time for a healthy snack!"))

	return objs
}

func Load(app fyne.App, content *fyne.Container) []fyne.CanvasObject {
	objs := fetchDefaultObjs()

	//uncomment for Snack Demo
	//Active = true
	//Snack = "a Baby Carrot"

	if Snack != "" {
		message := fmt.Sprint("How about ", Snack)
		objs = append(objs, widget.NewLabel(message))
		content.Objects = objs
		content.Layout.Layout(content.Objects, content.Size())
		Snack = ""
	}

	if Active {
		objs = append(objs, widget.NewButton("Done!", func() {
			//add a point to the user and reset the panel
			Active = false
			objs = fetchDefaultObjs()
			content.Objects = objs
			content.Layout.Layout(content.Objects, content.Size())
		}))
		objs = append(objs, widget.NewButton("Skip", func() {
			//dismiss the action by resetting the panel
			Active = false
			objs = fetchDefaultObjs()
			content.Objects = objs
			content.Layout.Layout(content.Objects, content.Size())
		}))
	}

	return objs
}