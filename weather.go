package weather

import (
	"encoding/json"
	"fmt"
)

const BaseURL = "https://api.openweathermap.org"

type OWMResponse struct {
	Weather []struct {
		Main string
	}
	Main struct {
		Temp float64
	}
}

type Conditions struct {
	Summary     string
	Temperature float64
}

func ParseResponse(data []byte) (Conditions, error) {
	var resp OWMResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return Conditions{}, fmt.Errorf(
			"invalid API response: %s: %w", data, err)
	}
	if len(resp.Weather) < 1 {
		return Conditions{}, fmt.Errorf(
			"invalid API response %s: want at least one weather element", data)
	}
	conditions := Conditions{
		Summary:     resp.Weather[0].Main,
		Temperature: resp.Main.Temp,
	}
	return conditions, nil
}

func FormatURL(city, apikey string) string {
	return fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s&units=metric",
		BaseURL, city, apikey)
}
