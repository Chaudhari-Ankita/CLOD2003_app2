package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Kartik
type Message struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var data []Message

// sanjay
func postUser(w http.ResponseWriter, r *http.Request) {
	var user Message
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &user)

	if err != nil {
		http.Error(w, "Invalid read request body", http.StatusBadRequest)
		return
	}
	data = append(data, user)
	fmt.Fprint(w, "Your data is successfully added")
}

// Kartik
func getParticularUser(w http.ResponseWriter, id string) {
}

// Roshni
func getAllUsers(w http.ResponseWriter) {

}

// Roshni
func getUser(w http.ResponseWriter, r *http.Request) {
}

// Ankita
func updateUser(w http.ResponseWriter, r *http.Request) {
}

// Mili
func deleteUser(w http.ResponseWriter, r *http.Request) {
}

// avneet
func handleMethod(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodPost:
		postUser(w, r)
	case http.MethodGet:
		getUser(w, r)
	case http.MethodPut:
		updateUser(w, r)
	case http.MethodDelete:
		deleteUser(w, r)
	}
}

// avneet
func main() {
	pNumber := ":8091"
	http.HandleFunc("/tasks/", handleMethod)
	http.HandleFunc("/tasks", handleMethod)
	http.ListenAndServe(pNumber, nil)
}
