package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Address model to address ...
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Number  int    `json:"number"`
	Country string `json:"country"`
}

// Contact model to contact ...
type Contact struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Phone   string  `json:"phone"`
	Email   string  `json:"email"`
	Adrress Address `json:"address,omitempty"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, path %s!", r.URL.Path[1:])
}

// AllContacts return all contacts.
func AllContacts(w http.ResponseWriter, r *http.Request) {
	var contacts []Contact
	contacts = append(contacts, Contact{Name: "mario", Phone: "53531531", Email: "lol@lol.com"})
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(contacts); err != nil {
		fmt.Fprintf(w, "Invalid reponse payload")
		return
	}
}

// GetContact return a specific contact.
func GetContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contactID := vars["id"]

	// logic to search contact.
	var contact Contact
	contact.ID = contactID
	contact.Name = "mario"
	contact.Phone = "53153153"
	contact.Email = "lol@lol.com"

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(contact); err != nil {
		fmt.Fprintf(w, "Invalid reponse payload")
		return
	}
}

// SaveContact save a new contact.
func SaveContact(w http.ResponseWriter, r *http.Request) {
	var newContact Contact
	if err := json.NewDecoder(r.Body).Decode(&newContact); err != nil {
		fmt.Fprintf(w, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	// logic to save new contact
	newContact.ID = "3"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Println(newContact.ID, newContact.Name)
	if err := json.NewEncoder(w).Encode(newContact); err != nil {
		fmt.Fprintf(w, "Invalid response payload")
		return
	}
}

// UpdateContact update a specific contact.
func UpdateContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contactID := vars["id"]

	// logic to search contact
	var contact Contact
	contact.ID = contactID
	contact.Name = "idival"

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(contact); err != nil {
		fmt.Fprintf(w, "Invalid reponse payload")
		return
	}
}

// DeleteContact delete a specific contact.
func DeleteContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_ = vars["id"]

	// logic to remove contact

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/contacts/", AllContacts).Methods("GET")
	r.HandleFunc("/contacts/{id}/", GetContact).Methods("GET")
	r.HandleFunc("/contacts/", SaveContact).Methods("POST")
	r.HandleFunc("/contacts/{id}/", UpdateContact).Methods("PUT")
	r.HandleFunc("/contacts/{id}/", DeleteContact).Methods("DELETE")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err.Error())
	}
}
