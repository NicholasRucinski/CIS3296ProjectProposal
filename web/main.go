package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type WeatherData struct {
	Temperature string
	Humidity    string
}

func main() {
	fmt.Println("Hello World")

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-weather-data/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func h1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	data := map[string][]WeatherData{
		"WeatherData": {
			{Temperature: "65", Humidity: "50"},
			{Temperature: "60", Humidity: "51"},
		},
	}
	tmpl.Execute(w, data)
}

func h2(w http.ResponseWriter, r *http.Request) {
	temp := r.PostFormValue("Temperature")
	hum := r.PostFormValue("Humidity")
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "WeatherDataElement", WeatherData{Temperature: temp, Humidity: hum})
}
