package exercise

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"net/url"
	"time"
)

var Active = true
var TakeWalk = true
var ticker *time.Ticker

func Load(content *fyne.Container) []fyne.CanvasObject {
	var objs []fyne.CanvasObject
	var StopTicker = false

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
			tickerStart := 1
			StopTicker = true
			ticker = time.NewTicker(time.Second * 1)

			for{
				select{
				case <-ticker.C:
					fmt.Println(tickerStart)
					message := fmt.Sprint("Timer: \r\n", tickerStart)
					content.Objects = append(objs, widget.NewLabel(message))
					content.Refresh()
					tickerStart++
				}
			}
		}))
	}

	if StopTicker {
		objs = append(objs, widget.NewButton("Stop", func() {
			//start walk timer and show stop button
			StopTicker = false
			ticker.Stop()
		}))
	}

	return objs
}