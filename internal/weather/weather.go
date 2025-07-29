package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/hectormalot/omgo"

	"weather-telegram-bot/internal/utils"
)

func NewClient() *omgo.Client {
	c, err := omgo.NewClient()
	if err != nil {
		log.Fatalf("error creating OpenMeteo client: %v", err)
	}
	return &c
}

// GeocodeCity queries the Nominatim API to find places matching the given city name
func GeocodeCity(city string, limit int) ([]NominatimPlace, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	uri := "https://nominatim.openstreetmap.org/search?" + url.Values{"q": {city}, "format": {"json"}, "limit": {strconv.Itoa(limit)}}.Encode()
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("User-Agent", "github.com/itskoshkin/weather-bot")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer utils.Closer(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status code: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	var result []NominatimPlace
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %v", err)
	}
	return result, nil
}

// ReverseGeocode queries the Nominatim API to find the address of a location given its latitude and longitude
func ReverseGeocode(lat, lon float64) (string, error) {
	endpoint := "https://nominatim.openstreetmap.org/reverse"
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?lat=%f&lon=%f&format=json", endpoint, lat, lon), nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("User-Agent", "github.com/itskoshkin/weather-bot")
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer utils.Closer(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected response status code: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}
	var result NominatimPlace
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response body: %v", err)
	}
	return result.Address.City + ", " + result.Address.Road, nil
}

// GetCurrentWeatherByCode returns a human-readable description of the current weather based on the provided WMO weather interpretation codes
//
// See https://gist.github.com/stellasphere/9490c195ed2b53c707087c8c2db4ec0c
func GetCurrentWeatherByCode(code float64) string {
	switch code {
	case 0:
		return "Ясно"
	case 1:
		return "В основном ясно"
	case 2:
		return "Переменная облачность"
	case 3:
		return "Облачно"
	case 45, 48:
		return "Туман"
	case 51:
		return "Слабая морось"
	case 53:
		return "Морось"
	case 55:
		return "Сильная морось"
	case 56:
		return "Слабая ледяная морось"
	case 57:
		return "Ледяная морось"
	case 61:
		return "Небольшой дождь"
	case 63:
		return "Дождь"
	case 65:
		return "Сильный дождь"
	case 66:
		return "Слабый ледяной дождь"
	case 67:
		return "Ледяной дождь"
	case 71:
		return "Небольшой снег"
	case 73:
		return "Снег"
	case 75:
		return "Сильный снег"
	case 77:
		return "Снежные зерна"
	case 80:
		return "Небольшие ливни"
	case 81:
		return "Ливни"
	case 82:
		return "Сильные ливни"
	case 85:
		return "Небольшой снежный ливень"
	case 86:
		return "Снежный ливень"
	case 95:
		return "Гроза"
	case 96:
		return "Гроза с небольшим градом"
	case 99:
		return "Гроза с градом"
	default:
		return "Неизвестно"
	}
}

// GetWindDirection returns the wind direction name in Russian based on the input angle in degrees.
// It divides the compass into eight 45° sectors and normalizes any angle (including negative values) into the [0, 360) range.
// The sectors are:
//
//	North:          [337.5°, 360) and [0°, 22.5°)
//	North-East:     [22.5°, 67.5°)
//	East:           [67.5°, 112.5°)
//	South-East:     [112.5°, 157.5°)
//	South:          [157.5°, 202.5°)
//	South-West:     [202.5°, 247.5°)
//	West:           [247.5°, 292.5°)
//	North-West:     [292.5°, 337.5°)
//
// If the angle falls outside these ranges (which is unlikely after normalization), it returns "неизвестно".
func GetWindDirection(degrees float64) string {
	for degrees < 0 {
		degrees += 360
	}
	degrees = math.Mod(degrees, 360)

	switch {
	case degrees >= 337.5 || degrees < 22.5:
		return "северный"
	case degrees >= 22.5 && degrees < 67.5:
		return "северо-восточный"
	case degrees >= 67.5 && degrees < 112.5:
		return "восточный"
	case degrees >= 112.5 && degrees < 157.5:
		return "юго-восточный"
	case degrees >= 157.5 && degrees < 202.5:
		return "южный"
	case degrees >= 202.5 && degrees < 247.5:
		return "юго-западный"
	case degrees >= 247.5 && degrees < 292.5:
		return "западный"
	case degrees >= 292.5 && degrees < 337.5:
		return "северо-западный"
	default:
		return "неизвестно"
	}
}
