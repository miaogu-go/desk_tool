package encode_decode

import (
	"encoding/base64"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"tool/tool"
)

func buildEncode() fyne.CanvasObject {
	b := NewBase64EncodeDe()
	return b.layout()
}

type base64EncodeDe struct {
	in  *widget.Entry
	out *widget.Entry
}

func NewBase64EncodeDe() *base64EncodeDe {
	return &base64EncodeDe{}
}

func (b *base64EncodeDe) input() fyne.CanvasObject {
	label := widget.NewLabel("请输入要进行 Base64 编码或解码的字符")
	b.in = widget.NewMultiLineEntry()
	b.in.SetMinRowsVisible(8)
	return container.NewVBox(label, b.in)
}

func (b *base64EncodeDe) output() fyne.CanvasObject {
	b.out = widget.NewMultiLineEntry()
	b.out.SetMinRowsVisible(8)
	b.out.Disable()
	return b.out
}

func (b *base64EncodeDe) encodeButton() fyne.CanvasObject {
	but := widget.NewButton("编码", func() {
		inStr := b.in.Text
		if inStr == "" {
			log.Printf("encodeButton in is empty")
			return
		}
		encodeStr := base64.StdEncoding.EncodeToString([]byte(inStr))
		b.out.SetText(encodeStr)
	})

	return but
}

func (b *base64EncodeDe) copyButton() fyne.CanvasObject {
	but := widget.NewButton("复制", func() {
		tool.Copy(b.out.Text)
	})

	return but
}

func (b *base64EncodeDe) decodeButton() fyne.CanvasObject {
	but := widget.NewButton("解码", func() {
		inStr := b.in.Text
		if inStr == "" {
			log.Printf("decodeButton in is empty")
			return
		}
		decodeBytes, err := base64.StdEncoding.DecodeString(inStr)
		if err != nil {
			log.Printf("DecodeString err::%s", err.Error())
			return
		}
		b.out.SetText(string(decodeBytes))
	})

	return but
}

func (b *base64EncodeDe) changeButton() fyne.CanvasObject {
	but := widget.NewButton("交换", func() {
		inStr := b.in.Text
		b.in.SetText(b.out.Text)
		b.out.SetText(inStr)
	})

	return but
}

func (b *base64EncodeDe) layout() fyne.CanvasObject {
	buts := container.NewHBox(b.encodeButton(), b.decodeButton(), b.copyButton(), b.changeButton())
	content := container.NewVBox(b.input(), b.output(), buts)
	return widget.NewCard("base64编解码", "", content)
}
