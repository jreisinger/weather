package weather_test

import (
	"os"
	"testing"
	"weather"

	"github.com/google/go-cmp/cmp"
)

func TestParseResponse(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/weather.json")
	if err != nil {
		t.Fatal(err)
	}
	want := weather.Conditions{
		Summary:     "Fog",
		Temperature: 277.38,
	}
	got, err := weather.ParseResponse(data)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestFomatURL(t *testing.T) {
	t.Parallel()
	want := "https://api.openweathermap.org/data/2.5/weather?q=Bratislava&appid=dummyAPIKey&units=metric"
	got := weather.FormatURL("Bratislava", "dummyAPIKey")
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}
