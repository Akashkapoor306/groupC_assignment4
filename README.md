# GroupC_Assignment4 README

Weather API Server

This Go application provides a simple API server to retrieve weather information for cities and additional details like wind speed, visibility, rain, and snow.
Getting Started

## Files in Our Project
- **main.go:** This file has all the code for our server and functions.

### How to Use
To use our project, you need to have Go installed on your computer. You can download Go from here.

1. Clone the Repository

First, clone this project to your local machine using Git:

```
https://go.dev/doc/install
```


```
git clone https://github.com/Akashkapoor306/groupC_assignment4.git
cd groupC_Assignment4
```

**Running the application**
- Open a terminal or command prompt.
- Change the directory to where you have our project files.
- Type “go mod init” to initiate the folder
- Type **'go run main.go'** to run the file
- 
```go run main.go```

- Now, the server should be running on port 8012.

### Available Actions
- **City Weather** 
To fetch weather data for a city, make a GET request to /city with the name parameter set to the desired city.

Example request:

```
GET /city?name=New York

```
- **Wind and Visibility**
To retrieve additional weather details like wind speed, visibility, rain, and snow, make a GET request to /WindAndVisibility with the city parameter set to the desired city

```
GET /WindAndVisibility?city=London

```

### Configuration

The application requires an API key from OpenWeatherMap. You need to set your API key in the apiKey variable inside the getWeather and getWindAndVisibility functions

```
apiKey := "your_openweathermap_api_key"

```

### Team Members
Our team worked together on this project. Here are our members:

- Abhinav Mahajan - 500230044
- Akash - 500218794
- Vinay Chhabra - 500228151
- Ashbir - 500228410
- Jevica - 500218849
- Bilal Nawaz -  500228652
- Mohamed Ayan Khatri - 500226334
- Simrandeep Singh - 500229180
- Nikhil Kaushik - 500223528
- Rajkarn kaur - 500226333