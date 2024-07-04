package main

import (
	"fmt"
	"net/http"
	"time"
)

func TimeHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		date := time.Now()
		week := date.Weekday()
		day := date.Day()
		month := date.Month()
		year := date.Year()

		w.Header().Set("Content-Type", "text/plain")
		res := fmt.Sprintf("%s, %d %s %d", week, day, month, year)
		w.Write([]byte(res))
	}) // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		sayHallo := ""
		if name == "" {
			sayHallo = fmt.Sprintf("Hello there")
			w.Write([]byte(sayHallo))
			return
		}

		sayHallo = fmt.Sprintf("Hello, %s!", name)
		w.Write([]byte(sayHallo))
	}) // TODO: replace this
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/time", TimeHandler())
	mux.HandleFunc("/hello", SayHelloHandler())
	http.ListenAndServe("localhost:8080", nil)
}
