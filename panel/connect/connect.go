package connect

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/stoyvo/wellness-buddy/assets"
)

var Active bool
var Joke string

func fetchDefaultObjs() []fyne.CanvasObject{
	var objs []fyne.CanvasObject

	objs = append(objs, widget.NewLabel("Time to connect with a loved one!"))

	jokeImg := canvas.NewImageFromResource(assets.ResourceJokePng)
	jokeImg.FillMode = canvas.ImageFillOriginal
	objs = append(objs, jokeImg)

	return objs
}

func Load(app fyne.App, content *fyne.Container) []fyne.CanvasObject {
	objs := fetchDefaultObjs()

	//uncomment for Connect Demo
	Active = true
	Joke = "I couldn't figure out how the seat belt worked. Then it just clicked."

	if Joke != "" {
		message := fmt.Sprint("Perhaps start with a joke? Like: \"", Joke, "\"")
		objs = append(objs, widget.NewLabel(message))
		content.Objects = objs
		content.Layout.Layout(content.Objects, content.Size())
	}

	if Active {
		objs = append(objs, widget.NewButton("Done!", func() {
			app.Preferences().SetInt(
				"connect",
				app.Preferences().Int("connect") + 1,
			)
			objs = fetchDefaultObjs()
			content.Objects = objs
			content.Layout.Layout(content.Objects, content.Size())
			Joke = ""
			Active = false
		}))
		objs = append(objs, widget.NewButton("Skip", func() {
			//dismiss the action by resetting the panel
			objs = fetchDefaultObjs()
			content.Objects = objs
			content.Layout.Layout(content.Objects, content.Size())
			Joke = ""
			Active = false
		}))
	}

	return objs
}