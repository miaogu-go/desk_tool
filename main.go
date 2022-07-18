package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	"image/color"
	"log"
	"os"
	"strings"
	"tool/data"
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
	//text := canvas.NewText("xxxx", color.Black)
	text2 := canvas.NewText("xxxx", color.Black)

	split := container.NewHSplit(makeMenuTree(), text2)

	mainWindow.SetContent(split)

	return mainWindow
}

func makeMenuTree() fyne.CanvasObject {
	tree := &widget.Tree{
		BaseWidget: widget.BaseWidget{},
		//Root:       "",
		ChildUIDs: func(uid string) (c []string) {
			c = data.MenuTree[uid]
			return c
		},
		CreateNode: func(branch bool) (o fyne.CanvasObject) {
			return widget.NewLabel("tree123456")
		},
		IsBranch: func(uid string) (ok bool) {
			children, ok := data.MenuTree[uid]
			return ok && len(children) > 0
		},
		OnBranchClosed: nil,
		OnBranchOpened: nil,
		OnSelected: func(uid widget.TreeNodeID) {
			log.Println(uid)
		},
		OnUnselected: nil,
		UpdateNode: func(uid widget.TreeNodeID, branch bool, node fyne.CanvasObject) {
			node.(*widget.Label).SetText(uid)
		},
	}
	/*data := map[string][]string{
		"":  {"A"},
		"A": {"B", "D", "H", "J", "L", "O", "P", "S", "V"},
		"B": {"C"},
		"C": {"abc"},
		"D": {"E"},
		"E": {"F", "G"},
		"F": {"adef"},
		"G": {"adeg"},
		"H": {"I"},
		"I": {"ahi"},
		"O": {"ao"},
		"P": {"Q"},
		"Q": {"R"},
		"R": {"apqr"},
		"S": {"T"},
		"T": {"U"},
		"U": {"astu"},
		"V": {"W"},
		"W": {"X"},
		"X": {"Y"},
		"Y": {"Z"},
		"Z": {"avwxyz"},
	}

	tree := widget.NewTreeWithStrings(data)
	tree.OnSelected = func(id string) {
		fmt.Println("Tree node selected:", id)
	}
	tree.OnUnselected = func(id string) {
		fmt.Println("Tree node unselected:", id)
	}
	tree.OpenBranch("A")
	tree.OpenBranch("D")
	tree.OpenBranch("E")
	tree.OpenBranch("L")
	tree.OpenBranch("M")*/

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
