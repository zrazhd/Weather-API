package main

import (
	"fmt"
	"log"
	"net/http"
)

type WeatherResponse struct {
	City     string  `json:"city"`
	Humidity float32 `json:"humidity"`
	Celcius  float32 `json:"celcius"`
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	city := r.PathValue("city")
	_, err := w.Write([]byte(city))
	if err != nil {
		log.Println("cant write in header")
	}

}
func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET {city}", getWeather)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println("can not start server")
	}
}
