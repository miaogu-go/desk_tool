package data

import (
	"fyne.io/fyne/v2"
	"tool/data/unix_time"
)

type ScreenRenderHandle func(w fyne.Window) fyne.CanvasObject

type Content struct {
	Title, Intro string
	View         ScreenRenderHandle
	SupportWeb   bool
}

var (
	Contents = map[string]Content{
		"introduce": {
			Title:      "Introduce",
			Intro:      "this is welcome",
			View:       _welcome,
			SupportWeb: true,
		},
		"unixTime": {
			Title: "UnixTime",
			View:  unix_time.UnixTime,
		},
	}

	Menus = map[string][]string{
		"": {"introduce", "unixTime"},
	}
)
