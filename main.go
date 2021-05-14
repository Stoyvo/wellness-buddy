package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var App fyne.App

func init() {
	startJobs()
}

func main() {
	App = app.NewWithID("com.bounteous.wellness-buddy")
	w := App.NewWindow("Wellness Buddy")
	loadApp(w)
	w.Resize(fyne.NewSize(680, 480))

	w.ShowAndRun()
}