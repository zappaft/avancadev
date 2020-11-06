package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}

// Result struct contains status
type Result struct {
	Status string
}

func home(w http.ResponseWriter, r *http.Request) {
	valid := r.PostFormValue("valid")

	result := Result{Status: getMessage(valid)}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))
}

func getMessage(valid string) string {
	msg := "Sorry, not sorry"
	if valid == "valid" {
		msg = "Congrats on your buy."
	}
	return msg
}
