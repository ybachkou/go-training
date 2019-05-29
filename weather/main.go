package main

import (
	"errors"
	"fmt"
	"os"
	"weather/weather"
)

func main() {
	var city string
	fmt.Print("Укажите город: ")
	_, err := fmt.Scanln(&city)
	if err != nil {
		fmt.Println(err.Error())
	}
	m := weather.Meteorologist{}
	wf := weather.WeatherForecast{}
	w := weather.Meteorologist.GetWeather(m, city)
	if w.Cod != 200 {
		fmt.Println(errors.New("Error URL"))
		os.Exit(1)
	}
	result := weather.WeatherForecast.FormatWeather(wf, w)
	fmt.Println(result)
}
