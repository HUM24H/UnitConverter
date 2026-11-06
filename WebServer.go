package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)
	http.Handle("/length", http.HandlerFunc(lengthHandler))
	http.Handle("/weight", http.HandlerFunc(weightHandler))
	http.Handle("/temperature", http.HandlerFunc(temperatureHandler))

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func lengthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	convertFrom := r.FormValue("from")
	convertTo := r.FormValue("to")
	convertValue := r.FormValue("value")

	value, err := strconv.ParseFloat(convertValue, 64)
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}

	result, err := convertLength(value, convertFrom, convertTo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.FormatFloat(result, 'f', -1, 64)))
}

func weightHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	convertFrom := r.FormValue("from")
	convertTo := r.FormValue("to")
	convertValue := r.FormValue("value")

	value, err := strconv.ParseFloat(convertValue, 64)
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}

	result, err := convertWeight(value, convertFrom, convertTo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.FormatFloat(result, 'f', -1, 64)))
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	convertFrom := r.FormValue("from")
	convertTo := r.FormValue("to")
	convertValue := r.FormValue("value")

	value, err := strconv.ParseFloat(convertValue, 64)
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}

	result, err := convertTemperature(value, convertFrom, convertTo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.FormatFloat(result, 'f', -1, 64)))
}
