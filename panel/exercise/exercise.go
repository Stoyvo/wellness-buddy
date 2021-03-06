package exercise

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/stoyvo/wellness-buddy/assets"
	"math"
	"net/url"
	"strconv"
	"time"
)

var Active bool
var TakeWalk bool
var ticker *time.Ticker
var tickerStart = 1

func fetchDefaultObjs() []fyne.CanvasObject{
	var objs []fyne.CanvasObject
	objs = append(objs, widget.NewLabel("Time to exercise"))
	//add youtube video link
	u, _ := url.Parse("https://www.youtube.com/watch?v=ferw4VhbN54")
	objs = append(objs, widget.NewHyperlink("Watch the \"5 Minute Full Body Stretching Routine!\" video on YouTube", u))

	return objs
}

func Load(app fyne.App, content *fyne.Container) []fyne.CanvasObject {
	objs := fetchDefaultObjs()

	imgContainer := container.NewCenter()
	var imgFrames = []*canvas.Image{
		canvas.NewImageFromResource(assets.ResourceStretch1Png),
		canvas.NewImageFromResource(assets.ResourceStretch2Png),
		canvas.NewImageFromResource(assets.ResourceStretch3Png),
		canvas.NewImageFromResource(assets.ResourceStretch4Png),
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

	//uncomment for Stretch Demo
	//Active = true
	//uncomment for Take a Walk Timer Demo
	TakeWalk = true

	if Active {
		objs = append(objs, widget.NewButton("Done!", func() {
			//add a point to the user and reset the panel
			app.Preferences().SetInt(
				"exercise",
				app.Preferences().Int("exercise") + 1,
			)
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

	if TakeWalk {
		objs = append(objs, widget.NewButton("Start", func() {
			//start walk timer and show stop button
			//timerStarted = true
			objs = fetchDefaultObjs()
			ticker = time.NewTicker(time.Second * 1)

			objs = append(objs, widget.NewButton("Stop", func() {
				//start walk timer and show stop button
				//timerStopped = true
				TakeWalk = false
				ticker.Stop()
				objs = fetchDefaultObjs()
				message := fmt.Sprint("Timer Stopped at: ", secondsToHuman(tickerStart))
				objs = append(objs, widget.NewLabel(message))
				content.Objects = objs
				content.Layout.Layout(content.Objects, content.Size())
			}))

			timerLabel := widget.NewLabel("Timer:")
			objs = append(objs, timerLabel)
			content.Objects = objs
			content.Layout.Layout(content.Objects, content.Size())

			go func(timerLabel *widget.Label) {
				for _ = range ticker.C {
					message := fmt.Sprint("Timer: ", secondsToHuman(tickerStart))
					timerLabel.SetText(message)
					tickerStart++
				}
			}(timerLabel)
		}))
	}

	return objs
}

func plural(count int, singular string) (result string) {
	if (count == 1) || (count == 0) {
		result = strconv.Itoa(count) + " " + singular + " "
	} else {
		result = strconv.Itoa(count) + " " + singular + "s "
	}
	return
}

func secondsToHuman(input int) (result string) {
	years := math.Floor(float64(input) / 60 / 60 / 24 / 7 / 30 / 12)
	seconds := input % (60 * 60 * 24 * 7 * 30 * 12)
	months := math.Floor(float64(seconds) / 60 / 60 / 24 / 7 / 30)
	seconds = input % (60 * 60 * 24 * 7 * 30)
	weeks := math.Floor(float64(seconds) / 60 / 60 / 24 / 7)
	seconds = input % (60 * 60 * 24 * 7)
	days := math.Floor(float64(seconds) / 60 / 60 / 24)
	seconds = input % (60 * 60 * 24)
	hours := math.Floor(float64(seconds) / 60 / 60)
	seconds = input % (60 * 60)
	minutes := math.Floor(float64(seconds) / 60)
	seconds = input % 60

	if years > 0 {
		result = plural(int(years), "year") + plural(int(months), "month") + plural(int(weeks), "week") + plural(int(days), "day") + plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
	} else if months > 0 {
		result = plural(int(months), "month") + plural(int(weeks), "week") + plural(int(days), "day") + plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
	} else if weeks > 0 {
		result = plural(int(weeks), "week") + plural(int(days), "day") + plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
	} else if days > 0 {
		result = plural(int(days), "day") + plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
	} else if hours > 0 {
		result = plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
	} else if minutes > 0 {
		result = plural(int(minutes), "minute") + plural(int(seconds), "second")
	} else {
		result = plural(int(seconds), "second")
	}

	return
}