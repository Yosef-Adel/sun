package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	q := "Cairo"
	if len(os.Args) > 2 {
		q = os.Args[1]
	}
	weatherAPI := os.Getenv("WEATHER_API")
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + weatherAPI + "&q=" + q + "&days=1&aqi=no&alerts=no")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic("Weather API Not Available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var w Weather
	if err := json.Unmarshal(body, &w); err != nil {
		panic(err)
	}
	location, current, hours := w.Location, w.Current, w.Forecast.Forecastday[0].Hour
	fmt.Printf("%s, %s: %.0fC, %s\n",
		location.Country,
		location.Name,
		current.TempC,
		current.Condition.Text,
	)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)
		if date.Before(time.Now()) {
			continue
		}
		msg := fmt.Sprintf(
			"%s - %.0fC, %.0f%%, %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)
		if hour.ChanceOfRain < 40 {
			fmt.Print(msg)
		} else {
			color.Red(msg)
		}
	}
}
