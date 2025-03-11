package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	// Creating the main application
	app := tview.NewApplication()

	// Creating the list that will hold the .desktop files
	desktopFilesList := tview.NewList()
	desktopFilesList.Box = tview.NewBox().SetBorder(true).SetTitle("  ged  ").SetTitleAlign(tview.AlignCenter).SetBorderPadding(0,0,1,1)

	// Creating the layout, it will be a grid layout with 9x9
	layout := tview.NewGrid().SetRows(0,0,0,0,0,0,0,0,0).SetColumns(0,0,0,0,0,0,0,0,0)
	layout.Box = tview.NewBox().SetBackgroundColor(tcell.NewRGBColor(20, 20, 20))

	// Creating the page logic in order to add pop-ups
	pages := tview.NewPages()
	pages.AddPage("main", layout, true, true)

	// Getting the desktop file names
	names, err := GetDesktopFiles()
	if err != nil {
		panic(err)
	}

	// This function structure will be changed and FileOptions function will be used instead.
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
			// This InputCapture mechanism needs to be revised.
			askModal.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyESC {
					pages.RemovePage("ask")
					app.SetFocus(desktopFilesList)
				}
				return event
			})
		}).ShowSecondaryText(false)
	}

	// Adding the list to the grid layout
	layout.AddItem(desktopFilesList, 1, 1, 7, 7, 0, 0, false)

	// Setting up the app
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