package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Weather struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
}

func main() {
	//To use OpenWeather you need key 
	key := ""
	// In varable 'place' remember about use '+' instead of spaces
	place := "New+York+City,US"
	url := "https://api.openweathermap.org/data/2.5/weather?q="+place+"&appid=" + key + "&units=metric"


	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	var weather Weather
	if err := json.NewDecoder(response.Body).Decode(&weather); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Temperature:", weather.Main.Temp)
	fmt.Println("Now in your city is:", weather.Weather[0].Main)
}
