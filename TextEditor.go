package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var count int = 1

func showTextEditor() {
	w := MyApp.NewWindow("Notepad")
	w.Resize(fyne.NewSize(600, 600))
	w.CenterOnScreen()
	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("My Text Editor"),
		),
	)

	content.Add(widget.NewButton("Add New File", func() {
		content.Add(widget.NewLabel("New File " + strconv.Itoa(count)))
		count++
	}))

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter Text...")

	saveBtn := widget.NewButton("Save", func() {
		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData := []byte(input.Text)
				uc.Write(textData)
			}, w)
		if count != 1 {
			saveFileDialog.SetFileName("New File " + strconv.Itoa(count-1) + ".txt")
		} else {
			saveFileDialog.SetFileName("New File " + strconv.Itoa(count) + ".txt")
		}
		saveFileDialog.Show()
	})

	openBtn := widget.NewButton("Open", func() {
		openFileDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				readData, _ := ioutil.ReadAll(r)

				output := fyne.NewStaticResource("New File", readData)
				ViewData := widget.NewMultiLineEntry()
				ViewData.SetText(string(output.StaticContent))

				w := fyne.CurrentApp().NewWindow(
					string(output.StaticName))
				w.SetContent(container.NewScroll(ViewData))
				w.Resize(fyne.NewSize(400, 400))
				w.Show()
			}, w)

		openFileDialog.SetFilter(
			storage.NewExtensionFileFilter([]string{".txt"}))
		openFileDialog.Show()
	})

	TextContainer := container.NewVBox(
		container.NewVBox(
			content,
			input,
			container.NewHBox(
				saveBtn,
				openBtn,
			),
		),
	)
	w.SetContent(
		container.NewBorder(Deskbtn, nil, nil, nil, TextContainer),
	)
	w.Show()
}
