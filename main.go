package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func init() {
	startJobs()
}

func main() {
	a := app.NewWithID("com.bounteous.wellness-buddy")
	w := a.NewWindow("Wellness Buddy")
	loadApp(w)
	w.Resize(fyne.NewSize(680, 480))

	w.ShowAndRun()
}