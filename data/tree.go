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
			Title: "时间戳转换",
			View:  unix_time.UnixTime,
		},
		"encode": {
			Title: "编解码",
			View:  nil,
		},
	}

	Menus = map[string][]string{
		"": {"introduce", "unixTime", "encode"},
	}
)
