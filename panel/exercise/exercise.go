package exercise

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"net/url"
	"time"
)

var Active bool
var TakeWalk bool
var ticker *time.Ticker

func fetchDefaultObjs() []fyne.CanvasObject{
	var objs []fyne.CanvasObject
	objs = append(objs, widget.NewLabel("Time to exercise"))
	//add youtube video link
	u, _ := url.Parse("https://www.youtube.com/watch?v=ferw4VhbN54")
	objs = append(objs, widget.NewHyperlink("Watch the \"5 Minute Full Body Stretching Routine!\" video on YouTube", u))

	return objs
}

func Load(app fyne.App, content *fyne.Container) []fyne.CanvasObject {
	tickerStart := 1
	objs := fetchDefaultObjs()

	if Active {
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
	}

	if TakeWalk {
		objs = append(objs, widget.NewButton("Start", func() {
			//start walk timer and show stop button
			ticker = time.NewTicker(time.Second * 1)

			objs = append(objs, widget.NewButton("Stop", func() {
				//start walk timer and show stop button
				ticker.Stop()
				objs = fetchDefaultObjs()
				message := fmt.Sprint("Timer Stopped at: ", tickerStart)
				objs = append(objs, widget.NewLabel(message))
				content.Objects = objs
				content.Layout.Layout(content.Objects, content.Size())
			}))

			objs = fetchDefaultObjs()
			timerLabel := widget.NewLabel("Timer: 0")
			objs = append(objs, timerLabel)
			content.Objects = objs
			content.Layout.Layout(content.Objects, content.Size())

			go func(timerLabel *widget.Label) {
				for _ = range ticker.C {
					message := fmt.Sprint("Timer: ", tickerStart)
					timerLabel.SetText(message)
					tickerStart++
				}
			}(timerLabel)
		}))
	}

	content.Objects = objs
	content.Layout.Layout(content.Objects, content.Size())

	return objs
}