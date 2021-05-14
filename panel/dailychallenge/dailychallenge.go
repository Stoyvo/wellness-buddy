package dailychallenge

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/stoyvo/wellness-buddy/assets"
	"time"
)

var ChallengeAction string
var challengeLabel *widget.Label
var doneBtn *widget.Button
var skipBtn *widget.Button

func initBtns(app fyne.App) {
	challengeLabel = widget.NewLabel(ChallengeAction)
	doneBtn = widget.NewButton("Done!", func() {
		app.Preferences().SetInt(
			"challenge",
			app.Preferences().Int("challenge") + 1,
		)
		hideBtns()
	})
	doneBtn.Hidden = true
	skipBtn = widget.NewButton("Skip!", hideBtns)
	skipBtn.Hidden = true
}

func hideBtns() {
	ChallengeAction = ""
	challengeLabel.Hide()
	doneBtn.Hide()
	skipBtn.Hide()
}
func Load(app fyne.App, content *fyne.Container) []fyne.CanvasObject {
	initBtns(app)
	var objs []fyne.CanvasObject

	objs = append(objs, widget.NewLabel("Daily Challenge"))
	objs = append(objs, widget.NewLabel("Come back here when you have a cool challenge!"))

	// Images
	var day = int(time.Now().Weekday())
	//day = 4 // TO test days, set a number between 1-5
	imgContainer := container.NewCenter()
	var imgFrames []*canvas.Image
	switch challenge := day; challenge {
		case 1:
			imgFrames = append(imgFrames, canvas.NewImageFromResource(assets.ResourceTreePng))
		case 2:
			imgFrames = append(imgFrames, canvas.NewImageFromResource(assets.ResourceSquat1Png))
			imgFrames = append(imgFrames, canvas.NewImageFromResource(assets.ResourceSquat2Png))
		case 3:
			imgFrames = append(imgFrames, canvas.NewImageFromResource(assets.ResourceJumpingjacks1Png))
			imgFrames = append(imgFrames, canvas.NewImageFromResource(assets.ResourceJumpingjacks2Png))
		case 4:

		case 5:
			imgFrames = append(imgFrames, canvas.NewImageFromResource(assets.ResourceHydrantPng))
	}
	totalFrames := len(imgFrames)
	go func() {
		var i int
		for {
			if i >= totalFrames {
				i=0
			}
			imgFrames[i].FillMode = canvas.ImageFillOriginal
			imgContainer.Objects = []fyne.CanvasObject{imgFrames[i]}
			imgContainer.Layout.Layout(imgContainer.Objects, imgContainer.Size())
			content.Refresh()
			time.Sleep(1 * time.Second)
			i++
		}
	}()
	objs = append(objs, imgContainer)

	// Buttons
	objs = append(objs, challengeLabel)
	objs = append(objs, doneBtn)
	objs = append(objs, skipBtn)

	//uncomment for Daily Challenge Demo
	//ChallengeAction = "Have a burnout exercise (1 minute of squats or jumping jacks). Feel free to share your experience on #b_active"

	if ChallengeAction != "" {
		challengeLabel.SetText(ChallengeAction)
		challengeLabel.Show()
		doneBtn.Show()
		skipBtn.Show()
		content.Refresh()
	}



	return objs
}