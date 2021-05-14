package dailychallenge

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var ChallengeAction string

func fetchDefaultObjs() []fyne.CanvasObject{
	var objs []fyne.CanvasObject

	objs = append(objs, widget.NewLabel("Daily Challenge"))

	objs = append(objs, widget.NewLabel("Come back here when you have a cool challenge!"))

	return objs
}

func Load(content *fyne.Container) []fyne.CanvasObject {
	objs := fetchDefaultObjs()

	if ChallengeAction != "" {
		objs = append(objs, widget.NewLabel(ChallengeAction))

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

		content.Objects = objs
		content.Layout.Layout(content.Objects, content.Size())
	}

	return objs
}