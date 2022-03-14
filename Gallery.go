package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"io/ioutil"
	"log"
	"strings"
)

func showGallery(w fyne.Window) {
	root_src := "/mnt/data/Wallpapers"
	files, err := ioutil.ReadDir(root_src)
	if err != nil {
		log.Fatal(err)
	}
	tabs := container.NewAppTabs()
	for _, file := range files {
		if file.IsDir() == false {
			extension := strings.Split(file.Name(), ".")[1]
			if extension == "png" || extension == "jpeg" || extension == "jpg" {
				image := canvas.NewImageFromFile(root_src + "/" + file.Name())
				image.FillMode = canvas.ImageFillContain
				tabs.Append(container.NewTabItem(file.Name(), image))
			}
		}
	}
	tabs.SetTabLocation(container.TabLocationLeading)
	GalleryContainer := tabs

	w.SetContent(
		container.NewBorder(PanelContent, nil, nil, nil, GalleryContainer),
	)

	w.Resize(fyne.NewSize(600, 600))
	w.Show()
}
