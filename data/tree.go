package data

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Content struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
	SupportWeb   bool
}

var (
	Contents = map[string]Content{
		"welcome": {
			Title: "welcome",
			Intro: "this is welcome",
			View: func(w fyne.Window) fyne.CanvasObject {
				widgetLabel := widget.NewLabel("test")
				return widgetLabel
			},
			SupportWeb: true,
		},
	}

	MenuTree = map[string][]string{
		"": {"welcome"},
	}
)
