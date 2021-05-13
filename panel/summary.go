package panel

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var SummaryPanel *Summary

type Summary struct {
	name, totalPoints *widget.Label
	btnExample *widget.Button
}

func (s *Summary) Load(win fyne.Window) fyne.CanvasObject {
	s.name = widget.NewLabel("")
	s.totalPoints = widget.NewLabel("")

	s.btnExample = widget.NewButton("CLICK ME", func() {

	})
	buttons := container.NewHBox(
		layout.NewSpacer(),
		s.btnExample,
	)

	container.NewBorder(nil, buttons, nil, nil, nil)
}