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
	currentTime := binding.NewString()
	currentTimeLabel := widget.NewLabel("当前时间:")
	ticker := startTicker(currentTime)

	isStop := false
	stopButton := widget.NewButton("停止", func() {
		ticker.Stop()
		isStop = true
	})
	startButton := widget.NewButton("开始", func() {
		if isStop {
			ticker = startTicker(currentTime)
		}
	})
	refreshButton := widget.NewButton("刷新", func() {
		if isStop {
			err := currentTime.Set(fmt.Sprintf("%d", tool.GetCurrentUnixTime()))
			if err != nil {
				log.Printf("_unixTime unix ticker err:%s\n", err.Error())
			}
		}
	})
	copyButton := widget.NewButton("复制", func() {
		currentTimeStr, err := currentTime.Get()
		if err != nil {
			log.Printf("_unixTime get currentTime err:%s", err.Error())
			return
		}
		err = clipboard.WriteAll(currentTimeStr)
		if err != nil {
			log.Printf("_unixTime copyButton err:%s", err.Error())
		}
	})

	currentTimeCont := container.New(
		layout.NewHBoxLayout(),
		currentTimeLabel, widget.NewLabelWithData(currentTime),
		stopButton, startButton, refreshButton, copyButton,
	)

	return container.New(layout.NewVBoxLayout(), currentTimeCont)
}

func startTicker(currentTime binding.String) *time.Ticker {
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				err := currentTime.Set(fmt.Sprintf("%d", tool.GetCurrentUnixTime()))
				if err != nil {
					log.Printf("startTicker unix ticker err:%s\n", err.Error())
				}
			}
		}
	}()

	return ticker
}

type unixTime struct {
	ticker          *time.Ticker
	currentUnixTime binding.String
}

func New() *unixTime {
	return &unixTime{
		ticker:          time.NewTicker(time.Second),
		currentUnixTime: binding.NewString(),
	}
}

func (u *unixTime) startTicker() {
	go func() {
		for {
			select {
			case <-u.ticker.C:
				err := u.currentUnixTime.Set(fmt.Sprintf("%d", tool.GetCurrentUnixTime()))
				if err != nil {
					log.Printf("startTicker unix ticker err:%s\n", err.Error())
				}
			}
		}
	}()
}
