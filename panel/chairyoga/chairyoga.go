package chairyoga

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

var Active bool

func Load(content *fyne.Container) []fyne.CanvasObject {
	var objs []fyne.CanvasObject

	objs = append(objs, widget.NewLabel("Chair Yoga"))
	//add youtube video link
	u, _ := url.Parse("https://www.youtube.com/watch?v=m4t9nCW3630")
	objs = append(objs, widget.NewHyperlink("Watch the YouTube video", u))

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