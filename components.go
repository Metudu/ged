package main

import (
	"github.com/rivo/tview"
)

// This function builds the pop-up when user selects a desktop file
func FileOptions(filename string) {
	askModal := tview.NewModal()
	askModal.Box = tview.NewBox().SetTitle(" " + filename + " ").SetBorder(true)
	askModal.AddButtons([]string{"YES", "CANCEL"})
	pages.AddPage("ask", askModal, true, true)
	visible := GetVisibility(filename)
	if visible {
		askModal.SetText("This app icon is set to visible. Do you want to make it invisible?")
	} else {
		askModal.SetText("This app icon is set to invisible. Do you want to make it visible?")
	}
	askModal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "YES" {
			if err := SetVisibility(filename, !visible); err != nil {
				askModal.SetText(err.Error())
			} else {
				if !visible {
					askModal.SetText("Visibility is set to visible!")
				} else {
					askModal.SetText("Visibility is set to invisible!")
				}
			}

			askModal.ClearButtons()
		} else if buttonLabel == "CANCEL" {
			pages.RemovePage("ask")
			app.SetFocus(desktopFilesList)
		}
	})
}