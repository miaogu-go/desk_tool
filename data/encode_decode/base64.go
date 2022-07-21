package encode_decode

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

type base64EncodeDe struct {
	in  *widget.Entry
	out *widget.Entry
}

func (b *base64EncodeDe) input() fyne.CanvasObject {
	label := widget.NewLabel("请输入要进行 Base64 编码或解码的字符")
	b.in = widget.NewMultiLineEntry()
	b.in.SetMinRowsVisible(5)
	return container.NewVBox(label, b.in)
}

func (b *base64EncodeDe) output() fyne.CanvasObject {
	b.out = widget.NewMultiLineEntry()
	b.out.SetMinRowsVisible(5)
	return b.out
}

func (b *base64EncodeDe) encodeButton() fyne.CanvasObject {
	but := widget.NewButton("编码", func() {
		inStr := b.in.Text
		if inStr == "" {
			log.Printf("encodeButton in is empty")
			return
		}
	})
	return b.out
}

func (b *base64EncodeDe) layout() fyne.CanvasObject {

	return widget.NewCard("base64编解码", "", nil)
}
