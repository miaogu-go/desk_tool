package unix_time

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	"log"
	"time"
)

var (
	//cstZone 东八区时区
	cstZone = time.FixedZone("CST", 3600*8)
)

func buildDateToUnixStr() fyne.CanvasObject {
	dt := dateToUnix{
		dateStr: binding.NewString(),
		unix:    binding.NewString(),
	}

	return dt.layout()
}

type dateToUnix struct {
	dateStr binding.String
	unix    binding.String
}

func (d *dateToUnix) dateIn() fyne.CanvasObject {
	in := widget.NewEntryWithData(d.dateStr)
	in.SetPlaceHolder("eg:2022-07-08 11:29:25")
	return in
}

func (d *dateToUnix) unixOut() fyne.CanvasObject {
	label := widget.NewLabel("输出：")
	out := widget.NewLabelWithData(d.unix)
	return container.NewHBox(label, out)
}

func (d *dateToUnix) confirmButton() fyne.CanvasObject {
	but := widget.NewButton("转换", func() {
		date, err := d.dateStr.Get()
		if err != nil {
			log.Printf("confirmButton dateStr.Get err:%s", err.Error())
			return
		}
		if date == "" {
			log.Printf("date is empty")
			return
		}
		inTime, err := time.ParseInLocation(timeDateFormat, date, cstZone)
		if err != nil {
			log.Printf("confirmButton time.Parse err:%s,val:%s", err.Error(), date)
			return
		}
		log.Println(inTime.Unix())
		err = d.unix.Set(fmt.Sprintf("%d", inTime.Unix()))
		if err != nil {
			log.Printf("confirmButton unix.Set err:%s", err.Error())
		}
	})

	return but
}

func (d *dateToUnix) copyButton() fyne.CanvasObject {
	but := widget.NewButton("复制", func() {
		dateStr, err := d.unix.Get()
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

func (d *dateToUnix) layout() fyne.CanvasObject {
	buts := container.NewHBox(d.confirmButton(), d.copyButton())
	content := container.NewVBox(d.dateIn(), d.unixOut(), buts)
	return widget.NewCard("时间日期转时间戳", "", content)
}
