package main

import (
	"testing"
	"time"
)

var (
	cl ContactList
)

func TestAdd(t *testing.T) {
	cl = NewContactList()
	cl.Add(&Contact{ID: 0, Name: "write code", Assighec: "Doston", Deadline: time.Now().UTC().AddDate(0, 0, 15), Done: false})
	cl.Add(&Contact{ID: 1, Name: "write code", Assighec: "Doston", Deadline: time.Now().UTC().AddDate(0, 0, 15), Done: true})
	if len(cl.contacts) <= 0 {
		t.Error("Error while adding contacts")
	}

	if cl.contacts[0].ID != 0 {
		t.Error("Error while adding")
	}
}
func TestUpdate(t *testing.T) {

	i := 0
	id := 0
	name := "code"
	assighec := "doni"
	deadline := time.Now().UTC().AddDate(0, 0, -11)
	done := true
	cl.Update(0, &Contact{ID: id, Name: name, Assighec: assighec, Deadline: deadline, Done: done})

	if cl.contacts[i].ID != id && cl.contacts[i].Name != name && cl.contacts[i].Assighec != assighec && cl.contacts[i].Deadline != deadline && cl.contacts[i].Done != done {
		t.Errorf("the function is incorrect")
	}

}
func TestDelete(t *testing.T) {
	cl.Delete(0)
	if len(cl.contacts) != 1 {
		t.Error("Error while adding contacts")

	}

}
