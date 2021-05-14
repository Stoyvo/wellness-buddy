package exercise

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

var Active bool
var TakeWalk bool

func Load() []fyne.CanvasObject {
	var objs []fyne.CanvasObject

	objs = append(objs, widget.NewLabel("Time to exercise"))
	//add youtube video link
	u, _ := url.Parse("https://www.youtube.com/watch?v=ferw4VhbN54")
	objs = append(objs, widget.NewHyperlink("Watch the YouTube video", u))

	if Active {
		objs = append(objs, widget.NewButton("Done!", func() {
			//add a point to the user and reset the panel
		}))
		objs = append(objs, widget.NewButton("Skip", func() {
			//dismiss the action by reseting the panel
		}))
	}

	if TakeWalk {
		objs = append(objs, widget.NewButton("Start", func() {
			//start walk timer and show stop button
		}))
	}

	return objs
}