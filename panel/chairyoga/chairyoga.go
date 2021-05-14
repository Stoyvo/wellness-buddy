package chairyoga

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/stoyvo/wellness-buddy/assets"
	"net/url"
	"time"
)

var Active bool

func fetchDefaultObjs() []fyne.CanvasObject{
	var objs []fyne.CanvasObject

	objs = append(objs, widget.NewLabel("Chair Yoga"))
	//add youtube video link
	u, _ := url.Parse("https://www.youtube.com/watch?v=m4t9nCW3630")
	objs = append(objs, widget.NewHyperlink("Watch the \"10 Minute Chair Yoga Practice\" video on YouTube", u))

	return objs
}

func Load(app fyne.App, content *fyne.Container) []fyne.CanvasObject {
	objs := fetchDefaultObjs()

	//uncomment for Chair Yoga Demo
	//Active = true

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

	imgContainer := container.NewCenter()
	var imgFrames = []*canvas.Image{
		canvas.NewImageFromResource(assets.ResourceChairyoga1Png),
		canvas.NewImageFromResource(assets.ResourceChairyoga2Png),
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

	return objs
}