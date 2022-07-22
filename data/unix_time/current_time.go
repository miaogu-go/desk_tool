package unix_time

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"log"
	"time"
	"tool/tool"
)

func buildCurrentTime() fyne.CanvasObject {
	curTime := currentTime{
		currentTimeStr: binding.NewString(),
		isStop:         false,
		ticker:         time.NewTicker(time.Second),
	}

	return curTime.layout()
}

//currentTime 当前时间戳
type currentTime struct {
	currentTimeStr binding.String
	isStop         bool
	ticker         *time.Ticker
}

//currentTimeLabel 当前时间戳标签
func (u *currentTime) currentTimeLabel() fyne.CanvasObject {
	timeOut := widget.NewLabelWithData(u.currentTimeStr)
	u.startTicker()

	return timeOut
}

func (u *currentTime) stopButton() fyne.CanvasObject {
	stopButton := widget.NewButton("停止", func() {
		u.ticker.Stop()
		u.isStop = true
	})

	return stopButton
}

func (u *currentTime) startButton() fyne.CanvasObject {
	startButton := widget.NewButton("开始", func() {
		if u.isStop {
			u.ticker.Reset(time.Second)
		}
	})

	return startButton
}

func (u *currentTime) refreshButton() fyne.CanvasObject {
	refreshButton := widget.NewButton("刷新", func() {
		if u.isStop {
			err := u.currentTimeStr.Set(fmt.Sprintf("%d", tool.GetCurrentUnixTime()))
			if err != nil {
				log.Printf("_currentTime unix ticker err:%s\n", err.Error())
			}
		}
	})

	return refreshButton
}

func (u *currentTime) copyButton() fyne.CanvasObject {
	copyButton := widget.NewButton("复制", func() {
		currentTimeStr, err := u.currentTimeStr.Get()
		if err != nil {
			log.Printf("_currentTime get currentTime err:%s", err.Error())
			return
		}
		tool.Copy(currentTimeStr)
	})

	return copyButton
}

//currentTimeLayout 当前时间戳模块布局
func (u *currentTime) layout() fyne.CanvasObject {
	buts := container.NewHBox(u.stopButton(), u.startButton(), u.refreshButton(), u.copyButton())
	content := container.NewVBox(u.currentTimeLabel(), buts)
	return widget.NewCard("当前时间戳", "", content)
}

func (u *currentTime) startTicker() {
	go func() {
		for {
			select {
			case <-u.ticker.C:
				err := u.currentTimeStr.Set(fmt.Sprintf("%d", tool.GetCurrentUnixTime()))
				if err != nil {
					log.Printf("startTicker unix ticker err:%s\n", err.Error())
				}
			}
		}
	}()
}
