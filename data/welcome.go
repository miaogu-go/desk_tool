package data

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	softDesc = `
> 小工具
`
	funcDesc = `
* 时间戳转换工具
`
)

func _welcome(w fyne.Window) fyne.CanvasObject {
	softDescCont := widget.NewRichTextFromMarkdown(softDesc)
	softDescCard := widget.NewCard("tool介绍", "", softDescCont)

	funcDescCont := widget.NewRichTextFromMarkdown(funcDesc)
	funcDescCard := widget.NewCard("功能介绍", "", funcDescCont)

	return container.NewGridWithColumns(2, softDescCard, funcDescCard)
}
