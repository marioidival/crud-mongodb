package dao

import (
	"log"
	"os"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	. "github.com/marioidival/crud-mongodb/model"
)

var db *mgo.Database

// Dao struct ...
type Dao struct {
	Database   string
	Collection string
}

// Connect with mongo server
func (m *Dao) Connect() {
	var mongoURI string
	mongoURI, ok := os.LookupEnv("MONGODB_URI")
	if !ok {
		mongoURI = "localhost:27017"
	}

	session, err := mgo.Dial(mongoURI)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Insert just save a new contact
func (m *Dao) Insert(contact Contact) (err error) {
	err = db.C(m.Collection).Insert(&contact)
	return
}

// Update just update a contact
func (m *Dao) Update(contact Contact) (err error) {
	err = db.C(m.Collection).UpdateId(contact.ID, &contact)
	return
}

// FindAll return all contacts
func (m *Dao) FindAll() (contacts []Contact, err error) {
	err = db.C(m.Collection).Find(bson.M{}).All(&contacts)
	return
}

// FindByID return specific contact
func (m *Dao) FindByID(id string) (contact Contact, err error) {
	err = db.C(m.Collection).FindId(bson.ObjectIdHex(id)).One(&contact)
	return
}

// Delete remove specific contact
func (m *Dao) Delete(id string) (err error) {
	err = db.C(m.Collection).RemoveId(bson.ObjectIdHex(id))
	return err
}

// FakeInsert ...
func (m *Dao) FakeInsert() (id bson.ObjectId, err error) {
	var contact Contact
	id = bson.NewObjectId()
	dao := Dao{Database: "crudmongodb", Collection: "contacts"}
	contact.ID = id
	contact.Name = "teste name"
	contact.Email = "teste email"
	contact.Phone = "57315315"
	err = dao.Insert(contact)
	return
}
