package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("com.jhollowe.hdhr-signal-utils")
	w := a.NewWindow("HDHomeRun Signal Utils")

	// discoveredSelector := widget.NewSelect([]string{"Option1", "Option2"}, func(option string) { log.Println("selected: ", option) })
	discoverButton := widget.NewButton("Discover Tuners", func() { log.Println("discover button pressed") })

	w.SetContent(discoverButton)
	// w.SetContent(discoveredSelector)

	w.ShowAndRun()
}
