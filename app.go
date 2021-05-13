package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/stoyvo/wellness-buddy/panel/dailychallenge"
	"github.com/stoyvo/wellness-buddy/panel/summary"
)

var content *fyne.Container
var navList *widget.List

type Panel struct {
	Title	*widget.Label
	Obj		func() []fyne.CanvasObject
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
		Title: widget.NewLabel("Breathing Exercises"),
	},
	{
		Title: widget.NewLabel("Exercise"),
	},
	{
		Title: widget.NewLabel("Hydrate"),
	},
	{
		Title: widget.NewLabel("Yoga"),
	},
	{
		Title: widget.NewLabel("Snacks"),
	},
}

func loadApp(win fyne.Window) {
	content = container.New(layout.NewVBoxLayout())

	navList = widget.NewList(func() int {
			return len(PanelList)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("A longish name")
		},
		func(id int, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(PanelList[id].Title.Text)
		},
	)

	navList.OnSelected = func(id int) {
		content.Objects = PanelList[id].Obj()
		content.Resize(content.Layout.MinSize(content.Objects))
		content.Layout.Layout(content.Objects, content.Size())
	}

	// Initialize first panel
	navList.Select(0)

	win.SetContent(
		container.NewBorder(
			widget.NewLabel("Wellness Buddy"), nil, navList, nil,
			content,
		),
	)
}
