package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	"image/color"
	"os"
	"strings"
)

func main() {
	mainApp := app.New()
	mw := makeMainWindow(mainApp)
	mw.ShowAndRun()

	os.Unsetenv("FYNE_FONT")
}

func makeMainWindow(mainApp fyne.App) fyne.Window {
	mainWindow := mainApp.NewWindow("tools")
	//设置窗口大小
	mainWindow.Resize(fyne.NewSize(500, 500))
	mainWindow.CenterOnScreen()
	//设置主窗口
	mainWindow.SetMaster()
	//设置主菜单
	mainWindow.SetMainMenu(makeMainMenu(mainApp, mainWindow))

	//todo left tree
	text := canvas.NewText("xxxx", color.Black)
	text2 := canvas.NewText("xxxx", color.Black)

	split := container.NewHSplit(text, text2)

	mainWindow.SetContent(split)

	return mainWindow
}

func makeMainMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {

	mainMenu := fyne.NewMainMenu(makeAboutMenu())

	return mainMenu
}

func makeAboutMenu() *fyne.Menu {
	aboutItem := fyne.NewMenuItem("关于", func() {
		w := fyne.CurrentApp().NewWindow("关于")
		richText := widget.NewRichTextFromMarkdown(`author: pu.qiang@qq.com`)
		content := container.New(layout.NewCenterLayout(), richText)
		w.SetContent(content)
		w.CenterOnScreen()
		w.Resize(fyne.NewSize(300, 200))
		w.SetFixedSize(true)
		w.Show()
	})
	about := fyne.NewMenu("关于", aboutItem)

	return about
}

func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		fmt.Println(path)
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		//YaHei.ttf
		if strings.Contains(path, "simhei.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
	fmt.Println("=============")
}
