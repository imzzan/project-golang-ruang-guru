package main

import (
	"net/http"
)

var students = []string{
	"Aditira",
	"Dito",
	"Afis",
	"Eddy",
}

func IsNameExists(name string) bool {
	for _, n := range students {
		if n == name {
			return true
		}
	}

	return false
}

func CheckStudentName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Data not found"))
			return
		}

		if r.Method != "GET" {
			w.WriteHeader(405)
			w.Write([]byte("Method is not allowed"))
			return
		}

		if IsNameExists(name) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Name is exists"))
			return
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Data not found"))
	}
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/students", CheckStudentName())
	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
