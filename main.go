package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.NewWithID("com.bounteous.wellness-buddy")
	w := a.NewWindow("Wellness Buddy")
	loadApp(w)
	w.Resize(fyne.NewSize(680, 400))

	w.ShowAndRun()
}