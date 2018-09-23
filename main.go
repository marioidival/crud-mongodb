package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	. "github.com/marioidival/crud-mongodb/dao"
	. "github.com/marioidival/crud-mongodb/model"
)

var dao = Dao{Database: "crudmongodb", Collection: "contacts"}

func init() {
	dao.Connect()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, path %s!", r.URL.Path[1:])
}

// AllContacts return all contacts.
func AllContacts(w http.ResponseWriter, r *http.Request) {
	var contacts []Contact
	contacts, err := dao.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

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

	contact, err := dao.FindByID(contactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

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
	newContact.ID = bson.NewObjectId()
	err := dao.Insert(newContact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
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
	contact.ID = bson.ObjectIdHex(contactID)
	contact.Name = "idival"

	err := dao.Update(contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(contact); err != nil {
		fmt.Fprintf(w, "Invalid reponse payload")
		return
	}
}

// DeleteContact delete a specific contact.
func DeleteContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contactID := vars["id"]

	err := dao.Delete(contactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	port := os.Getenv("PORT")
	r := mux.NewRouter()

	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/contacts/", AllContacts).Methods("GET")
	r.HandleFunc("/contacts/{id}/", GetContact).Methods("GET")
	r.HandleFunc("/contacts/", SaveContact).Methods("POST")
	r.HandleFunc("/contacts/{id}/", UpdateContact).Methods("PUT")
	r.HandleFunc("/contacts/{id}/", DeleteContact).Methods("DELETE")

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err.Error())
	}
}
