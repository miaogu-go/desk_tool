package unix_time

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	"log"
	"strconv"
	"time"
)

var (
	timeDateFormat = "2006-01-02 15:04:05"
)

func buildUnixToDateStr() fyne.CanvasObject {
	ut := unixToDateStr{
		unixTimeStr: binding.NewString(),
		dateStr:     binding.NewString(),
	}

	return ut.layout()
}

type unixToDateStr struct {
	unixTimeStr binding.String
	dateStr     binding.String
	unit        string
}

func (u *unixToDateStr) unixIn() fyne.CanvasObject {
	in := widget.NewEntryWithData(u.unixTimeStr)
	in.SetPlaceHolder("输入时间戳")
	return in
}

func (u *unixToDateStr) dateOut() fyne.CanvasObject {
	label := widget.NewLabel("输出：")
	out := widget.NewLabelWithData(u.dateStr)
	return container.NewHBox(label, out)
}

func (u *unixToDateStr) unitSelect() fyne.CanvasObject {
	options := []string{"s", "ms"}
	unit := widget.NewSelect(options, func(s string) {
		u.unit = s
	})
	unit.SetSelected("s")

	return unit
}

func (u *unixToDateStr) confirmButton() fyne.CanvasObject {
	but := widget.NewButton("转换", func() {
		t, err := u.unixTimeStr.Get()
		if err != nil {
			log.Printf("confirmButton unixTimeStr.Get err:%s\n", err.Error())
			return
		}
		if t == "" {
			log.Printf("confirmButton unixTimeStr is empty")
			return
		}

		tInt, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			log.Printf("confirmButton strconv.ParseInt err,err:%s,val:%s", err.Error(), t)
			return
		}

		switch u.unit {
		case "s":
			inTime := time.Unix(tInt, 0)
			err = u.dateStr.Set(u.unixToStr(inTime))
			if err != nil {
				log.Printf("confirmButton dateStr.Set err:%s", err.Error())
			}
		case "ms":
			inTime := time.UnixMilli(tInt)
			err = u.dateStr.Set(u.unixToStr(inTime))
			if err != nil {
				log.Printf("confirmButton dateStr.Set err:%s", err.Error())
			}
		}
	})

	return but
}

func (u *unixToDateStr) unixToStr(t time.Time) string {
	return t.Format(timeDateFormat)
}

func (u *unixToDateStr) copyButton() fyne.CanvasObject {
	but := widget.NewButton("复制", func() {
		dateStr, err := u.dateStr.Get()
		if err != nil {
			log.Printf("copyButton get currentTime err:%s", err.Error())
			return
		}
		err = clipboard.WriteAll(dateStr)
		if err != nil {
			log.Printf("copyButton copyButton err:%s", err.Error())
		}
	})

	return but
}

func (u *unixToDateStr) layout() fyne.CanvasObject {
	buts := container.NewHBox(u.unitSelect(), u.confirmButton(), u.copyButton())
	in := container.NewVBox(u.unixIn(), u.dateOut(), buts)
	return widget.NewCard("时间戳转字符串", "", in)
}
