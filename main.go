package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func CityHandler(res http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal("{'cities':'San Francisco, Amsterdam, Berlin, New York, Tokyo, Chicago, Wheaton, blah, Glen Ellyn, Carol Stream'}")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func IndexHandler (res http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal("hello, world")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func main() {
	port := os.Getenv("PORT") 
	if  port == "" {
		port = ":5000"
	}
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/cities.json", CityHandler)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
