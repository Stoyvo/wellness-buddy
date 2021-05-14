package snacks

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var Active bool
var Snack string
var snackLabel *widget.Label
var doneBtn *widget.Button
var skipBtn *widget.Button

func init() {
	snackLabel = widget.NewLabel("")
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

	objs = append(objs, widget.NewLabel("Time for a healthy snack!"))

	//uncomment for Snack Demo
	//Active = true
	//Snack = "a Baby Carrot"
	if Snack != "" {
		snackLabel.SetText("How about " + Snack + "?")
	}
	objs = append(objs, snackLabel)

	// Buttons
	objs = append(objs, doneBtn)
	objs = append(objs, skipBtn)

	if Active {
		doneBtn.Show()
		skipBtn.Show()
		content.Refresh()
	}

	return objs
}