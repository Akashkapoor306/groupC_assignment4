package main

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
