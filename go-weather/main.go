package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main(){

	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/", 
		func(w http.ResponseWriter, r *http.Request){
			city := strings.SplitN(r.URL.Path, "/", 3)[2]

			data, err := query(city)

			if err != nil{
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(data)
			// w.Write([]byte(city))
		},
	)


	http.ListenAndServe(":8089", nil)
}

func hello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from go!\n"))
}

func query(city string) (weatherData, error){
	// apiConfig, err := 
	apiConfig, err := loadApiConfig(".apiConfig")

	if err != nil{
		fmt.Printf("problem loading api key in query func")
		return weatherData{}, err
	}

	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)

	if err != nil{
		fmt.Printf("problem fetching data from api")
		return weatherData{}, err
	}

	defer resp.Body.Close()
	var d weatherData 
	if err := json.NewDecoder(resp.Body).Decode(&d); err!= nil{
		fmt.Printf("problem in decoding fetched data from api")
		return weatherData{}, nil
	}

	return d, nil

}

type weatherData struct{
	Name string `json:"name"`
	Main struct{
		Kelvin float64 `json:"temp"`
	}`json:main`
}

type apiConfigData struct{
	OpenWeatherMapApiKey string `json:"openWeatherMapApiKey"`
}

func loadApiConfig(filename string) (apiConfigData, error){
	bytes, err := os.ReadFile(filename)

	if err != nil{
		return apiConfigData{}, err
	}
	var c apiConfigData
	err = json.Unmarshal(bytes, &c)

	if err != nil{
		return apiConfigData{}, err
	}
	fmt.Println("your api key is: " + c.OpenWeatherMapApiKey)
	return c, nil
}