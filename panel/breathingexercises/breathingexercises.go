package breathingexercises

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

var Active bool

func fetchDefaultObjs() []fyne.CanvasObject{
	var objs []fyne.CanvasObject
	objs = append(objs, widget.NewLabel("Breathing Exercises"))
	//add youtube video link
	u, _ := url.Parse("https://www.youtube.com/watch?v=wfDTp2GogaQ")
	objs = append(objs, widget.NewHyperlink("Watch the \"Mindful Breathing Exercise\" video on YouTube", u))

	return objs
}

func Load(app fyne.App, content *fyne.Container) []fyne.CanvasObject {
	objs := fetchDefaultObjs()

	if Active {
		objs = append(objs, widget.NewButton("Done!", func() {
			//add a point to the user and reset the panel
			//make button unavailable for 2 hours
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

		content.Objects = objs
		content.Layout.Layout(content.Objects, content.Size())
	}

	return objs
}