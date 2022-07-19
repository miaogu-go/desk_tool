package data

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	"log"
	"time"
	"tool/tool"
)

func _unixTime(window fyne.Window) fyne.CanvasObject {
	unixTimeObj := New()
	unixTimeObj.currentTimeLayout()
	return unixTimeObj.layout()
}

type unixTime struct {
	ticker      *time.Ticker
	currentTime *currentTime
	containers  []fyne.CanvasObject
}

//currentTime 当前时间戳
type currentTime struct {
	currentTimeStr binding.String
	isStop         bool
	widgets        []fyne.CanvasObject
}

func New() *unixTime {
	ut := &unixTime{
		ticker:     time.NewTicker(time.Second),
		containers: make([]fyne.CanvasObject, 0),
	}

	currTime := &currentTime{
		currentTimeStr: binding.NewString(),
		isStop:         false,
		widgets:        make([]fyne.CanvasObject, 0),
	}
	ut.currentTime = currTime

	return ut
}

//currentTimeLabel 当前时间戳标签
func (u *unixTime) currentTimeLabel() *unixTime {
	currentTimeLabel := widget.NewLabel("当前时间戳:")
	timeCont := widget.NewLabelWithData(u.currentTime.currentTimeStr)
	u.startTicker()
	u.currentTime.widgets = append(u.currentTime.widgets, currentTimeLabel, timeCont)

	return u
}

func (u *unixTime) stopButton() *unixTime {
	stopButton := widget.NewButton("停止", func() {
		u.ticker.Stop()
		u.currentTime.isStop = true
	})
	u.currentTime.widgets = append(u.currentTime.widgets, stopButton)

	return u
}

func (u *unixTime) startButton() *unixTime {
	startButton := widget.NewButton("开始", func() {
		if u.currentTime.isStop {
			u.ticker.Reset(time.Second)
		}
	})
	u.currentTime.widgets = append(u.currentTime.widgets, startButton)

	return u
}

func (u *unixTime) refreshButton() *unixTime {
	refreshButton := widget.NewButton("刷新", func() {
		if u.currentTime.isStop {
			err := u.currentTime.currentTimeStr.Set(fmt.Sprintf("%d", tool.GetCurrentUnixTime()))
			if err != nil {
				log.Printf("_unixTime unix ticker err:%s\n", err.Error())
			}
		}
	})
	u.currentTime.widgets = append(u.currentTime.widgets, refreshButton)

	return u
}

func (u *unixTime) copyButton() *unixTime {
	copyButton := widget.NewButton("复制", func() {
		currentTimeStr, err := u.currentTime.currentTimeStr.Get()
		if err != nil {
			log.Printf("_unixTime get currentTime err:%s", err.Error())
			return
		}
		err = clipboard.WriteAll(currentTimeStr)
		if err != nil {
			log.Printf("_unixTime copyButton err:%s", err.Error())
		}
	})
	u.currentTime.widgets = append(u.currentTime.widgets, copyButton)

	return u
}

//currentTimeLayout 当前时间戳模块布局
func (u *unixTime) currentTimeLayout() {
	u.currentTimeLabel().stopButton().startButton().refreshButton().copyButton()
	currentTimeCont := container.New(layout.NewHBoxLayout(), u.currentTime.widgets...)
	u.containers = append(u.containers, currentTimeCont)
}

type UnixToDateStr struct {
	unixTimeStr binding.String
	dateStr     binding.String
}

//layout unix time 界面整体布局
func (u *unixTime) layout() *fyne.Container {
	return container.New(layout.NewVBoxLayout(), u.containers...)
}

func (u *unixTime) startTicker() {
	go func() {
		for {
			select {
			case <-u.ticker.C:
				err := u.currentTime.currentTimeStr.Set(fmt.Sprintf("%d", tool.GetCurrentUnixTime()))
				if err != nil {
					log.Printf("startTicker unix ticker err:%s\n", err.Error())
				}
			}
		}
	}()
}
