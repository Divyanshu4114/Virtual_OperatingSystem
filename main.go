package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

	var MyApp fyne.App = app.New()
	var MyWindow fyne.Window = MyApp.NewWindow("My OS")

	var btn1 fyne.Widget
	var btn2 fyne.Widget
	var btn3 fyne.Widget
	var btn4 fyne.Widget
	var Deskbtn fyne.Widget

	var image fyne.CanvasObject

    var PanelContent *fyne.Container

    func main(){
		MyApp.Settings().SetTheme(theme.LightTheme())
		img:= canvas.NewImageFromFile("/mnt/data/My Os Wallpaper.jpg")

		btn1 = widget.NewButtonWithIcon("Weather App", theme.InfoIcon(),func(){
			showWeatherApp(MyWindow)
		})

		btn2 = widget.NewButtonWithIcon("Gallery", theme.StorageIcon(),func(){
			showGallery(MyWindow)
		})

		btn3 = widget.NewButtonWithIcon("Calculator", theme.ContentAddIcon(),func(){
			showCalculator()
		})

		btn4 = widget.NewButtonWithIcon("Text Editor", theme.DocumentIcon(),func(){
			showTextEditor()
		})

		Deskbtn = widget.NewButtonWithIcon("Desktop", theme.HomeIcon(),func(){
			MyWindow.SetContent(container.NewBorder(PanelContent,nil,nil,nil,img))
		})

		PanelContent = container.NewVBox(container.NewGridWithColumns(5,Deskbtn,btn1,btn2,btn3,btn4))

		MyWindow.Resize(fyne.NewSize(1280,720))
        MyWindow.CenterOnScreen()
        
		MyWindow.SetContent(
			container.NewBorder(PanelContent,nil,nil,nil,img),
		)
	
		MyWindow.ShowAndRun()
	}