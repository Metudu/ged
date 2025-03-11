package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	// Initial 
	app := tview.NewApplication()
	desktopFilesList := tview.NewList()
	layout := getLayout()
	pages := tview.NewPages()
	pages.AddPage("main", layout, true, true)

	layout.Box = tview.NewBox().SetBackgroundColor(tcell.ColorGray)
	names, err := GetDesktopFiles()
	if err != nil {
		panic(err)
	}

	for _, name := range names {
		desktopFilesList.AddItem(name, "", 0, func() {
			askModal := tview.NewModal()
			askModal.SetText("Visible?").AddButtons([]string{"YES", "NO"})
			pages.AddPage("ask",askModal, true, true)
			askModal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "YES" {
					if err := SetVisibility(name, true); err != nil {
						askModal.SetText(err.Error())
					} else {
						askModal.SetText("Visibility set to true!")
					}
				} else if buttonLabel == "NO" {
					if err := SetVisibility(name, false); err != nil {
						askModal.SetText(err.Error())
					} else {
						askModal.SetText("Visibility set to false!")
					}
				}
			})
			askModal.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyESC {
					pages.RemovePage("ask")
					app.SetFocus(desktopFilesList)
				}
				return event
			})
		}).ShowSecondaryText(false)
	}

	layout.AddItem(desktopFilesList, 1, 1, 1, 1, 0, 0, false)


	app.SetRoot(pages, true).SetFocus(desktopFilesList)
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyESC {
			app.Stop()	
		}
		return event
	})

    if err := app.Run(); err != nil {
        panic(err)
    }
}

func getLayout() *tview.Grid {
	return tview.NewGrid().SetRows(0,0,0).SetColumns(0,0,0).SetBorders(true)
}
