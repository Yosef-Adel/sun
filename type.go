package main

type Location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Condition struct {
	Text string `json:"text"`
}

type Current struct {
	TempC     float64   `json:"temp_c"`
	Condition Condition `json:"condition"`
}

type HourForecast struct {
	TimeEpoch    int64     `json:"time_epoch"`
	TempC        float64   `json:"temp_c"`
	Condition    Condition `json:"condition"`
	ChanceOfRain float64   `json:"chance_of_rain"`
}

type DayForecast struct {
	Hour []HourForecast `json:"hour"`
}

type Forecast struct {
	Forecastday []DayForecast `json:"forecastday"`
}

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
	Forecast Forecast `json:"forecast"`
}
