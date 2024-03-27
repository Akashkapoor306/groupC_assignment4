package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ExternalAPIResponse defines the structure for the external weather API response
type ExternalAPIResponse struct {
	Temperature string `json:"temperature"`
	Wind        string `json:"wind"`
	Description string `json:"description"`
	Forecast    struct {
		Day         string `json:"day"`
		Temperature string `json:"temperature"`
		Wind        string `json:"wind"`
	} `json:"forecast"`
}

// WeatherResponse defines the structure of our API response
type WeatherResponse struct {
	City        string `json:"city"`
	Temperature string `json:"temperature"`
	Weather     string `json:"weather"` // We'll use the "Description" from the external API here
}

type WeatherResponse2 struct {
	Wind       string `json:"wind"`
	Visibility string `json:"visibility"`
	Rain       string `json:"rain"`
	Snow       string `json:"snow"`
}

// getWeather fetches weather data for a city using the external API
func getWeather(cityName string) (WeatherResponse, error) {
	url := fmt.Sprintf("https://goweather.herokuapp.com/weather/%s", cityName)
	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()

	var apiResponse ExternalAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return WeatherResponse{}, err
	}

	return WeatherResponse{
		City:        strings.Title(cityName),
		Temperature: apiResponse.Temperature,
		Weather:     apiResponse.Description,
	}, nil
}

//Function created by Nikhil

func cityHandler(w http.ResponseWriter, r *http.Request) {
	var weather WeatherResponse
	var cityName string
	var err error

	if r.Method == "GET" {
		cityName = r.URL.Query().Get("name")
	} else if r.Method == "POST" {
		var requestData struct{ Name string }
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		cityName = requestData.Name
	}

	weather, err = getWeather(cityName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}

//function by Jevica

func getWindAndVisibility(cityName string) (WeatherResponse2, error) {
	apiKey := "d51319b8aafa1e0618c55136562d617b"
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", cityName, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse2{}, err
	}
	defer resp.Body.Close()
	var weatherData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return WeatherResponse2{}, err
	}

	wind := ""
	if windSpeed, ok := weatherData["wind"].(map[string]interface{})["speed"].(float64); ok {
		wind = fmt.Sprintf("%.2f m/s", windSpeed)
	}

	visibility := ""
	if visibilityVal, ok := weatherData["visibility"].(float64); ok {
		visibility = fmt.Sprintf("%.2f meters", visibilityVal)
	}

	rain := ""
	if rainData, ok := weatherData["rain"].(map[string]interface{}); ok {
		if rainVal, ok := rainData["1h"].(float64); ok {
			rain = fmt.Sprintf("%.2f mm", rainVal)

		}
	}

	snow := ""
	if snowData, ok := weatherData["snow"].(map[string]interface{}); ok {
		if snowVal, ok := snowData["1h"].(float64); ok {
			snow = fmt.Sprintf("%.2f mm", snowVal)
		}
	}
	return WeatherResponse2{
		Wind:       wind,
		Visibility: visibility,
		Rain:       rain,
		Snow:       snow,
	}, nil
}
//function by Ashbir

func WindAndVisibilityHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		city := r.URL.Query().Get("city")
		if city == "" {
			http.Error(w, "City parameter is required", http.StatusBadRequest)
			return
		}

		weatherData, err := getWindAndVisibility(city)
		if err != nil {
			http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
			return
		}

		jsonData, err := json.Marshal(weatherData)
		if err != nil {
			http.Error(w, "Failed to encode weather data", http.StatusInternalServerError)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/WindAndVisibility", WindAndVisibilityHandler)
	//To consume the API ,Please refer as below url
	// http://localhost:8012/WindAndVisibility?city=CityName

	Dport := "8012"
	fmt.Printf("Server is starting on port: %v\n", Dport)
	http.ListenAndServe(":"+Dport, nil)
}
