package unix_time

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func UnixTime(window fyne.Window) fyne.CanvasObject {
	unixTimeObj := New()
	return unixTimeObj.layout()
}

type unixTime struct {
	containers []fyne.CanvasObject
}

func New() *unixTime {
	ut := &unixTime{
		containers: make([]fyne.CanvasObject, 0),
	}
	return ut
}

//layout unix time 界面整体布局
func (u *unixTime) layout() fyne.CanvasObject {
	u.containers = append(u.containers, buildCurrentTime(), buildUnixToDateStr(), buildDateToUnixStr())
	return container.NewVBox(u.containers...)
}
