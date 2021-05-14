package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/stoyvo/wellness-buddy/panel/breathingexercises"
	"github.com/stoyvo/wellness-buddy/panel/chairyoga"
	"github.com/stoyvo/wellness-buddy/panel/connect"
	"github.com/stoyvo/wellness-buddy/panel/dailychallenge"
	"github.com/stoyvo/wellness-buddy/panel/exercise"
	"github.com/stoyvo/wellness-buddy/panel/hydrate"
	"github.com/stoyvo/wellness-buddy/panel/snacks"
	"github.com/stoyvo/wellness-buddy/panel/summary"
)

var content *fyne.Container
var navList *widget.List

type Panel struct {
	Title	*widget.Label
	Obj		func(app fyne.App, content *fyne.Container) []fyne.CanvasObject
}

var PanelList = []Panel{
	{
		widget.NewLabel("Summary"),
		summary.Load,
	},
	{
		widget.NewLabel("Daily Challenge"),
		dailychallenge.Load,
	},
	{
		widget.NewLabel("Breathing Exercises"),
		breathingexercises.Load,
	},
	{
		widget.NewLabel("Exercise"),
		exercise.Load,
	},
	{
		widget.NewLabel("Hydrate"),
		hydrate.Load,
	},
	{
		widget.NewLabel("Chair Yoga"),
		chairyoga.Load,
	},
	{
		widget.NewLabel("Snacks"),
		snacks.Load,
	},
	{
		widget.NewLabel("Connect"),
		connect.Load,
	},
}

func loadApp(win fyne.Window) {
	content = container.New(layout.NewVBoxLayout())

	navList = widget.NewList(func() int {
			return len(PanelList)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Breathing Exercises") //this has to be the longest panel label
		},
		func(id int, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(PanelList[id].Title.Text)
		},
	)

	navList.OnSelected = func(id int) {
		if id == 6 && !snacks.Active {
			//randomize healthy snack and show without buttons
			snacks.Snack = randomizeSnack()
		}

		content.Objects = PanelList[id].Obj(App, content)
		content.Resize(content.Layout.MinSize(content.Objects))
		content.Layout.Layout(content.Objects, content.Size())
	}

	// Initialize first panel
	navList.Select(0)

	win.SetContent(
		container.NewBorder(
			widget.NewLabel("Wellness Buddy"),
			widget.NewLabelWithStyle("Copyright Bounteous 2021", fyne.TextAlignCenter, fyne.TextStyle{}),
			navList,
			nil,
			content,
		),
	)
}
