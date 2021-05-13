package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/stoyvo/wellness-buddy/panel/summary"
)

type Panel struct {
	Title	*widget.Label
	Obj		func(win fyne.Window) fyne.CanvasObject
}

var PanelList = []Panel{
	{
		widget.NewLabel("Summary"),
		func(win fyne.Window) fyne.CanvasObject {
			var obj summary.Summary
			return obj.Load(win)
		},
	},
	{
		Title: widget.NewLabel("Daily Challenge"),
	},
	{
		Title: widget.NewLabel("Breathing Exercises"),
	},
	{
		Title: widget.NewLabel("Walk"),
	},
	{
		Title: widget.NewLabel("Hydrate!"),
	},
	{
		Title: widget.NewLabel("Yoga"),
	},
	{
		Title: widget.NewLabel("Snacks"),
	},
}

func loadApp(win fyne.Window) {
	list := widget.NewList(func() int {
			return len(PanelList)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("A longish name")
		},
		func(id int, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(PanelList[id].Title.Text)
		})

	list.OnSelected = func(id int) {
		content := PanelList[id].Obj(win)
		win.SetContent(
			container.NewBorder(
				widget.NewLabel("Wellness Buddy"), nil, list, nil,
				container.NewBorder(
					nil, nil, nil, nil, content,
				),
			),
		)
	}
}
