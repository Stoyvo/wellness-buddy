package hydrate

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/stoyvo/wellness-buddy/assets"
	"time"
)

var Active bool
var doneBtn *widget.Button
var skipBtn *widget.Button

func init() {
	doneBtn = widget.NewButton("Done!", hideBtns)
	doneBtn.Hidden = true
	skipBtn = widget.NewButton("Skip!", hideBtns)
	skipBtn.Hidden = true
}

func hideBtns() {
	Active = false
	doneBtn.Hide()
	skipBtn.Hide()
}

func Load(app fyne.App, content *fyne.Container) []fyne.CanvasObject {
	var objs []fyne.CanvasObject

	objs = append(objs, widget.NewLabel("Remember to Hydrate"))

	// Animate
	imgContainer := container.NewCenter()
	var imgFrames = []*canvas.Image{
		canvas.NewImageFromResource(assets.ResourceWater1Png),
		canvas.NewImageFromResource(assets.ResourceWater2Png),
	}
	totalFrames := len(imgFrames)
	go func() {
		var i int
		for {
			if i >= totalFrames {
				i=0
			}
			imgFrames[i].FillMode = canvas.ImageFillOriginal
			imgContainer.Objects = []fyne.CanvasObject{imgFrames[i]}
			imgContainer.Layout.Layout(imgContainer.Objects, imgContainer.Size())
			content.Refresh()
			time.Sleep(1 * time.Second)
			i++
		}
	}()
	objs = append(objs, imgContainer)

	// Buttons
	objs = append(objs, doneBtn)
	objs = append(objs, skipBtn)

	//uncomment for Hydrate Demo
	Active = true

	if Active {
		doneBtn.Show()
		skipBtn.Show()
		content.Refresh()
	}

	return objs
}