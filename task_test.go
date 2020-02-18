package main
import (
	"testing"
	"time"
)
func TestAdd(t *testing.T) {
	var i int
	number := 5
	cl := NewContactList()
	cl.Add(&Contact{ID: number, Name: "write code", Assighec: "Doston", Deadline: time.Now().UTC().AddDate(0, 0, 15), Done: false})
	cl.Add(&Contact{ID: 8, Name: "write code", Assighec: "Doston", Deadline: time.Now().UTC().AddDate(0, 0, 15), Done: true})
	for i := 0; i < len(cl.contacts); i++ {
	}
	if cl.contacts[i].ID != number {
		t.Errorf("the function is incorrect")
	}

}
func TestUpdate(t *testing.T) {
	var i int
	 id:=8
	 name:="code"     
	 assighec:="doni"
	 deadline:=time.Now().UTC().AddDate(0, 0, -11)
	 done :=true
	cl := NewContactList()
	cl.Add(&Contact{ID: 100, Name: "write code", Assighec: "Doston", Deadline: time.Now().UTC().AddDate(0, 0, 15), Done: true})
	cl.Update(100, &Contact{ID: id, Name: name, Assighec: assighec, Deadline: deadline, Done: done})
	for i := 0; i < len(cl.contacts); i++ {
	}
	if cl.contacts[i].ID != id &&cl.contacts[i].Name != name&&cl.contacts[i].Assighec != assighec&&cl.contacts[i].Deadline != deadline&&cl.contacts[i].Done != done{
		t.Errorf("the function is incorrect")
	}

}