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

func Load(app fyne.App, content *fyne.Container) []fyne.CanvasObject {
	var objs []fyne.CanvasObject

	jokeImg := canvas.NewImageFromResource(assets.ResourceJokePng)
	jokeImg.FillMode = canvas.ImageFillOriginal
	objs = append(objs, jokeImg)

	objs = append(objs, widget.NewLabel("Time to connect with a loved one!"))

	if Joke != "" {
		message := fmt.Sprint("Perhaps start with a joke? Like: \"", Joke, "\"")
		objs = append(objs, widget.NewLabel(message))
	}

	if Active {
		objs = append(objs, widget.NewButton("Done!", func() {
			app.Preferences().SetInt(
				"connect",
				app.Preferences().Int("connect") + 1,
			)
		}))
		objs = append(objs, widget.NewButton("Skip", func() {
			//dismiss the action by resetting the panel
		}))
	}

	return objs
}