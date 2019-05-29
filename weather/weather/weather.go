package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const firstPartUrl = "http://api.openweathermap.org/data/2.5//weather?q="
const lastPartUrl = "&lang=ru&units=metric&appid=2c19a8c670afc70f2ae7a81f229fce3d"

type Meteorologist struct {
	Weather
}

type Weather struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`

	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`

	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`

	Sys struct {
		Sunrise int `json:"sunrise"`
		Sunset  int `json:"sunset"`
	} `json:"sys"`

	Name string `json:"name"`
	Cod  int    `json:"cod"`
}
type WeatherForecast struct {
}

func (wf WeatherForecast) FormatWeather(w Weather) string {
	temp, tempMin, tempMax := w.GetTemperature()
	speed, direction, gust := w.GetWind()
	sunrise := time.Unix(int64(w.Sys.Sunrise), 0).Format("15:04")
	sunset := time.Unix(int64(w.Sys.Sunset), 0).Format("15:04")
	firstSentence := "Сегодня в городе " + w.Name + " " + w.GetCloudiness() + ", средняя температура воздуха " + strconv.FormatFloat(temp, 'f', 1, 64) + "°С, минимальная температура " + strconv.FormatFloat(tempMin, 'f', 1, 64) + "°С, максимальная температура " + strconv.FormatFloat(tempMax, 'f', 1, 64) + "°С, ветер " + direction + " " + strconv.FormatFloat(speed, 'f', 1, 64) + "м/c с порывами до " + strconv.FormatFloat(gust, 'f', 1, 64) + "м/c. "
	secondSentence := "Влажность воздуха " + strconv.Itoa(w.GetHumidity()) + "%. "
	thirdSentence := "Восход солнца " + sunrise + ", заход солнца " + sunset + ". "
	str := firstSentence + secondSentence + thirdSentence
	return str
}

func (m Meteorologist) GetWeather(city string) Weather {
	url := firstPartUrl + city + lastPartUrl
	req, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	data, err1 := ioutil.ReadAll(req.Body)
	if err1 != nil {
		fmt.Println("Error at the data reading stage", err1)
	}
	defer req.Body.Close()
	var w Weather
	err2 := json.Unmarshal(data, &w)
	if err2 != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err2.Error())
	}
	return w
}

func (w Weather) GetTemperature() (temp, tempMin, tempMax float64) {
	return float64(w.Main.Temp), float64(w.Main.TempMin), float64(w.Main.TempMax)
}

func (w Weather) GetCloudiness() (description string) {
	return w.Weather[0].Description
}

func (w Weather) GetHumidity() (humidity int) {
	return w.Main.Humidity
}

func (w Weather) GetWind() (speed float64, direction string, gust float64) {
	result := getDirection(w.Wind.Deg)
	return w.Wind.Speed, result, w.Wind.Gust
}

func getDirection(deg float64) string {
	switch {
	case deg >= 338 && deg <= 360:
		return "северный"
	case deg >= 0 && deg <= 22:
		return "северный"
	case deg >= 23 && deg <= 67:
		return "северо-восточный"
	case deg >= 68 && deg <= 112:
		return "восточный"
	case deg >= 113 && deg <= 157:
		return "юго-восточный"
	case deg >= 158 && deg <= 202:
		return "южный"
	case deg >= 203 && deg <= 247:
		return "юго-западный"
	case deg >= 248 && deg <= 292:
		return "западный"
	case deg >= 293 && deg <= 337:
		return "северо-западный"
	default:
		return "Direction error"
	}
}
