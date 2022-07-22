package encode_decode

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func EncodeDecode(window fyne.Window) fyne.CanvasObject {
	e := New()
	return e.layout()
}

type encodeDecode struct {
	containers []fyne.CanvasObject
}

func New() *encodeDecode {
	return &encodeDecode{
		containers: make([]fyne.CanvasObject, 0),
	}
}

func (e *encodeDecode) layout() fyne.CanvasObject {
	e.containers = append(e.containers, buildEncode(), buildUrlEncode())
	return container.NewVBox(e.containers...)
}
