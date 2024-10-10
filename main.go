package main

import (
	"net/http"
)

// Kartik

// sanjay
func postUser(w http.ResponseWriter, r *http.Request) {
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
