package main

import (
	"encoding/json"
	"net/http"
)

// define the Model for a To Do item
type ToDoItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// store ToDoItem using a slice named ToDoItems
var ToDoItems []ToDoItem

func main() {
	// Your code here
	http.HandleFunc("/", ToDoListHandler)
	http.ListenAndServe(":8080", nil)

}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// define what function to call if we get a GET request
	if r.Method == http.MethodGet {
		// define what function to call if we get a GET request
		GetList(w, r)
	} else if r.Method == http.MethodPost {
		// define what function to call if we get a POST request
		AddToList(w, r)
	} else {
		// error handling for if we recieve any other kind of requesy
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)

	}

}

func GetList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// return empty slice to avoid null map error on frontend
	if len(ToDoItems) == 0 {
		_, _ = w.Write([]byte("[]"))
		return
	}

	// return the to do items or an error if there is one
	if err := json.NewEncoder(w).Encode(ToDoItems); err != nil {
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
	}
}

func AddToList(w http.ResponseWriter, r *http.Request) {
	var newToDo ToDoItem // initialise new to do item

	// check if the givn to do can be decoded into a ToDoItem
	if err := json.NewDecoder(r.Body).Decode(&newToDo); err != nil {
		// if it cannot we return a invalid input error
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	// check if title or description are empty
	if newToDo.Title == "" || newToDo.Description == "" {
		// if they are we return a bad request
		http.Error(w, "Title and Description are required fields", http.StatusBadRequest)
		return
	}

	// append the new item into memory/ storage if all is well
	ToDoItems = append(ToDoItems, newToDo)

	// return an OK status
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ToDoItems)
}
