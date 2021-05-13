package summary

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func Load() []fyne.CanvasObject {
	var objs []fyne.CanvasObject

	objs = append(objs, widget.NewLabel("Summary"))

	objs = append(objs, widget.NewLabel("500"))

	objs = append(objs, widget.NewButton("CLICK ME", func() {}))

	return objs
}