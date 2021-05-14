package dailychallenge

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var ChallengeAction string

func Load() []fyne.CanvasObject {
	var objs []fyne.CanvasObject

	objs = append(objs, widget.NewLabel("Daily Challenge"))

	objs = append(objs, widget.NewLabel("Come back here when you have a cool challenge!"))

	if ChallengeAction != "" {
		objs = append(objs, widget.NewLabel(ChallengeAction))
		objs = append(objs, widget.NewButton("Done!", func() {
			//add a point to the user and reset the panel
		}))
		objs = append(objs, widget.NewButton("Skip", func() {
			//dismiss the action by reseting the panel
		}))
	}

	return objs
}