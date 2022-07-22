package encode_decode

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"net/url"
	"tool/tool"
)

func buildUrlEncode() fyne.CanvasObject {
	u := NewUrlEncode()
	return u.layout()
}

type urlEncode struct {
	in  *widget.Entry
	out *widget.Entry
}

func NewUrlEncode() *urlEncode {
	return &urlEncode{}
}

func (u *urlEncode) input() fyne.CanvasObject {
	label := widget.NewLabel("请输入要进行 url 编解码字符串")
	u.in = widget.NewMultiLineEntry()
	u.in.SetMinRowsVisible(8)
	return container.NewVBox(label, u.in)
}

func (u *urlEncode) output() fyne.CanvasObject {
	u.out = widget.NewMultiLineEntry()
	u.out.SetMinRowsVisible(8)
	u.out.Disable()
	return u.out
}

func (u *urlEncode) encodeButton() fyne.CanvasObject {
	but := widget.NewButton("编码", func() {
		inStr := u.in.Text
		if inStr == "" {
			log.Printf("in is empty")
			return
		}
		//encStr := base64.URLEncoding.EncodeToString([]byte(inStr))
		encStr := url.QueryEscape(inStr)
		u.out.SetText(encStr)
	})

	return but
}

func (u *urlEncode) decodeButton() fyne.CanvasObject {
	but := widget.NewButton("解码", func() {
		inStr := u.in.Text
		if inStr == "" {
			log.Printf("in is empty")
			return
		}
		deStr, err := url.QueryUnescape(inStr)
		if err != nil {
			log.Printf("URLEncoding.DecodeString err:%s", err.Error())
			return
		}
		u.out.SetText(deStr)
	})

	return but
}

func (u *urlEncode) copyButton() fyne.CanvasObject {
	but := widget.NewButton("复制", func() {
		tool.Copy(u.out.Text)
	})

	return but
}

func (u *urlEncode) changeButton() fyne.CanvasObject {
	but := widget.NewButton("交换", func() {
		inStr := u.in.Text
		u.in.SetText(u.out.Text)
		u.out.SetText(inStr)
	})

	return but
}

func (u *urlEncode) layout() fyne.CanvasObject {
	buts := container.NewHBox(u.encodeButton(), u.decodeButton(), u.copyButton(), u.changeButton())
	content := container.NewVBox(u.input(), u.output(), buts)
	return widget.NewCard("Url编解码", "", content)
}
