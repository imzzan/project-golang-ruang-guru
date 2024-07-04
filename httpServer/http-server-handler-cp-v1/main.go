package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		date := time.Now()
		week := date.Weekday()
		day := date.Day()
		month := date.Month()
		year := date.Year()

		writer.Header().Set("Content-Type", "text/plain")
		res := fmt.Sprintf("%s, %d %s %d", week, day, month, year)
		writer.Write([]byte(res))
	} // TODO: replace this
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())
}
