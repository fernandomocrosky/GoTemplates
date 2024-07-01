package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8000"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 3)
	fmt.Fprintf(w, "This is the about page %d", sum)
}

func Divide(w http.ResponseWriter, r *http.Request) {
	result, err := divideValues(1.0, 2.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero")
	}

	fmt.Fprintf(w, "%.2f divided by %.2f is %.2f", 1.0, 2.0, result)
}

func addValues(n1, n2 int) (int) {
	return n1 + n2
}

func divideValues(n1, n2 float32) (float32, error) {

	if n2 == 0 {
        return 0, errors.New("Cannot divide by zero")
    }

	result := n1 / n2

	return result, nil
}


func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	fmt.Printf("Starting on port %s...\n", portNumber[1:])

	_ = http.ListenAndServe(portNumber, nil)
}