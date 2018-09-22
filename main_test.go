package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestAllContacts(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/contacts/", AllContacts).Methods("GET")
	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/contacts/")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error("error to transform body")
	}

	var contacts []Contact

	err = json.Unmarshal(body, &contacts)
	if err != nil {
		t.Error("error to transform body into contact")
	}

	if res.StatusCode != 200 {
		t.Error("expected 200")
	}

	if res.Header.Get("Content-Type") != "application/json" {
		t.Error("expected a json application")
	}

	if len(contacts) != 1 {
		t.Error("unexpected len of contacts")
	}
}

func TestGetContact(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/contacts/{id}/", GetContact).Methods("GET")
	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/contacts/1/")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error("error to transform body")
	}

	var contact Contact
	err = json.Unmarshal(body, &contact)
	if err != nil {
		t.Error("error to transform body into contact")
	}

	if res.StatusCode != 200 {
		t.Error("expected 200")
	}

	if res.Header.Get("Content-Type") != "application/json" {
		t.Error("expected a json application")
	}

	if contact.ID != "1" {
		t.Error("unexpected contact ID")
	}
}

func TestSaveContact(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/contacts/", SaveContact).Methods("POST")
	ts := httptest.NewServer(r)
	defer ts.Close()

	var contact Contact
	contact.Name = "mario"
	contact.Phone = "531531"
	contact.Email = "lol@lel.com"

	bContact, err := json.Marshal(contact)
	if err != nil {
		t.Error(err.Error())
	}
	updatedContact := bytes.NewReader(bContact)
	res, err := http.Post(ts.URL+"/contacts/", "application/json", updatedContact)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error("error to transform body")
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("unexpected status code, got %d", res.StatusCode)
	}
	if res.Header.Get("Content-Type") != "application/json" {
		t.Error("expected a json application")
	}

	var newContact Contact
	err = json.Unmarshal(body, &newContact)
	if err != nil {
		t.Error(err.Error())
	}

	if newContact.Name != contact.Name {
		t.Error("unexpeced name")
	}

	if newContact.ID == "" {
		t.Error("Unexpected contact ID")
	}
}

func TestUpdateContact(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/contacts/{id}/", UpdateContact).Methods("PUT")
	ts := httptest.NewServer(r)
	defer ts.Close()

	var contact Contact
	contact.Name = "idival"

	bContact, err := json.Marshal(contact)
	if err != nil {
		t.Error(err.Error())
	}
	updatedContactPayLoad := bytes.NewReader(bContact)
	req, err := http.NewRequest("PUT", ts.URL+"/contacts/3/", updatedContactPayLoad)
	if err != nil {
		log.Fatal(err.Error())
	}
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Error(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error("error to transform body")
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code, got %d", res.StatusCode)
	}
	if res.Header.Get("Content-Type") != "application/json" {
		t.Error("expected a json application")
	}

	var updatedContact Contact
	err = json.Unmarshal(body, &updatedContact)
	if err != nil {
		t.Error(err.Error())
	}

	if updatedContact.Name == "mario" {
		t.Errorf("unexpeced name, got: %s", updatedContact.Name)
	}

	if updatedContact.ID == "" {
		t.Error("Unexpected contact ID")
	}
}

func TestDeleteContact(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/contacts/{id}/", DeleteContact).Methods("DELETE")
	ts := httptest.NewServer(r)
	defer ts.Close()
	req, err := http.NewRequest("DELETE", ts.URL+"/contacts/3/", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Error(err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("unexpected status code, got %d", res.StatusCode)
	}
	if res.Header.Get("Content-Type") != "application/json" {
		t.Error("expected a json application")
	}

	// retry get contact removed.
}
