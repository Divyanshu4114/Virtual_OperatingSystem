package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func showWeatherApp(w fyne.Window) {

	// Api Part
	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=pantnagar&APPID=d1bd86ff64a06a8c6f59a20804f0f8d5")
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	myweather, err := UnmarshalWeather(body)
	if err != nil {
		fmt.Println(err)
	}

	img := canvas.NewImageFromFile("/home/divyanshu/Downloads/WeatherImage.jpg")
	img.FillMode = canvas.ImageFillOriginal

	label1 := canvas.NewText("WeatherDetails", color.White)
	label1.TextStyle = fyne.TextStyle{Italic: true, Bold: true}

	label2 := canvas.NewText(fmt.Sprintf("Country %s", myweather.Sys.Country), color.Black)
	label2.TextStyle = fyne.TextStyle{Italic: true, Bold: true}
	label3 := canvas.NewText(fmt.Sprintf("WindSpeed %.2f", myweather.Wind.Speed), color.Black)
	label3.TextStyle = fyne.TextStyle{Italic: true, Bold: true}
	label4 := canvas.NewText(fmt.Sprintf("MinTemp %.2f", myweather.Main.TempMin), color.Black)
	label4.TextStyle = fyne.TextStyle{Italic: true, Bold: true}
	label5 := canvas.NewText(fmt.Sprintf("MaxTemp %.2f", myweather.Main.TempMax), color.Black)
	label5.TextStyle = fyne.TextStyle{Italic: true, Bold: true}

	weatherContainer := container.NewVBox(
		label1,
		img,
		label2,
		label3,
		label4,
		label5,
		container.NewGridWithColumns(1),
	)

	w.SetContent(
		container.NewBorder(PanelContent, nil, nil, nil, weatherContainer))
	w.Show()
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    weather, err := UnmarshalWeather(bytes)
//    bytes, err = weather.Marshal()

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`
	Weather    []WeatherElement `json:"weather"`
	Base       string           `json:"base"`
	Main       Main             `json:"main"`
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`
	Clouds     Clouds           `json:"clouds"`
	Dt         int64            `json:"dt"`
	Sys        Sys              `json:"sys"`
	Timezone   int64            `json:"timezone"`
	ID         int64            `json:"id"`
	Name       string           `json:"name"`
	Cod        int64            `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
	SeaLevel  int64   `json:"sea_level"`
	GrndLevel int64   `json:"grnd_level"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type WeatherElement struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
	Gust  float64 `json:"gust"`
}
