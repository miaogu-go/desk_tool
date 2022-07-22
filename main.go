package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
	"tool/data"
)

type RightViewHandle func(cont data.Content)

func main() {
	mainApp := app.New()
	mw := makeMainWindow(mainApp)
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	mw.ShowAndRun()

	os.Unsetenv("FYNE_FONT")
}

func makeMainWindow(mainApp fyne.App) fyne.Window {
	mainWindow := mainApp.NewWindow("tools")
	//设置窗口大小
	mainWindow.Resize(fyne.NewSize(700, 500))
	mainWindow.CenterOnScreen()
	//设置主窗口
	mainWindow.SetMaster()
	//设置主菜单
	mainWindow.SetMainMenu(makeMainMenu(mainApp, mainWindow))

	rightCont := container.NewMax()
	viewHandle := func(cont data.Content) {
		view := cont.View(mainWindow)
		rightCont.Objects = []fyne.CanvasObject{view}
		rightCont.Refresh()
	}

	rightContent := container.NewBorder(nil, nil, nil, nil, rightCont)
	split := container.NewHSplit(makeMenuTree(viewHandle), rightContent)

	mainWindow.SetContent(split)

	return mainWindow
}

func makeMenuTree(viewHandle RightViewHandle) fyne.CanvasObject {
	tree := &widget.Tree{
		BaseWidget: widget.BaseWidget{},
		Root:       "",
		ChildUIDs: func(uid string) (c []string) {
			c = data.Menus[uid]
			return c
		},
		CreateNode: func(branch bool) (o fyne.CanvasObject) {
			return widget.NewLabel("tree")
		},
		IsBranch: func(uid string) (ok bool) {
			children, ok := data.Menus[uid]
			return ok && len(children) > 0
		},
		OnBranchClosed: nil,
		OnBranchOpened: nil,
		OnSelected: func(uid widget.TreeNodeID) {
			content, ok := data.Contents[uid]
			if !ok {
				log.Println("content is not exist", uid)
				return
			}
			viewHandle(content)
		},
		OnUnselected: nil,
		UpdateNode: func(uid widget.TreeNodeID, branch bool, node fyne.CanvasObject) {
			if cont, ok := data.Contents[uid]; ok {
				node.(*widget.Label).SetText(cont.Title)
				return
			}
			node.(*widget.Label).SetText(uid)
		},
	}

	return container.NewBorder(nil, nil, nil, nil, tree)
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
		//fmt.Println(path)
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		//YaHei.ttf
		if strings.Contains(path, "simhei.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
	//fmt.Println("=============")
}
