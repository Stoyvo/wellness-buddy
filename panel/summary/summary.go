package summary

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func Load(app fyne.App, content *fyne.Container) []fyne.CanvasObject {
	var objs []fyne.CanvasObject

	objs = append(objs, widget.NewLabelWithStyle("Summary", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))

	objs = append(objs, container.NewGridWithRows(4,
		container.NewGridWithColumns(2,
			widget.NewCard(
				"Daily Challenges",
				"",
				widget.NewLabelWithStyle("123", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			),
			widget.NewCard(
				"Breathing Exercises",
				"",
				widget.NewLabelWithStyle("123", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			),
		),
		container.NewGridWithColumns(2,
			widget.NewCard(
				"Exercise",
				"",
				widget.NewLabelWithStyle(
					strconv.Itoa(app.Preferences().Int("exercise")),
					fyne.TextAlignCenter, fyne.TextStyle{Bold: true},
				),
			),
			widget.NewCard(
				"Hydrate",
				"",
				widget.NewLabelWithStyle("123", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			),
		),
		container.NewGridWithColumns(2,
			widget.NewCard(
				"Yoga",
				"",
				widget.NewLabelWithStyle("123", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			),
			widget.NewCard(
				"Snacks",
				"",
				widget.NewLabelWithStyle("123", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			),
		),
		container.NewGridWithColumns(1,
			widget.NewCard(
				"Connect",
				"",
				widget.NewLabelWithStyle(
					strconv.Itoa(app.Preferences().Int("connect")),
					fyne.TextAlignCenter, fyne.TextStyle{Bold: true},
				),
			),
		),
	))

	return objs
}