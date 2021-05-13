package dailychallenge

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func Load() []fyne.CanvasObject {
	var objs []fyne.CanvasObject

	objs = append(objs, widget.NewLabel("SummaryAHHHHH"))

	objs = append(objs, widget.NewLabel("9999999"))

	objs = append(objs, widget.NewButton("DO NOT CLICK ME", func() {}))

	return objs
}