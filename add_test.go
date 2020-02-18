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
	cl.Add(&Contact{ID: 8, Name: "write code", Assighec: "Doston", Deadline: time.Now().UTC().AddDate(0, 0, 15), Done: false})
	for i := 0; i < len(cl.contacts); i++ {
	}
	if cl.contacts[i].ID != number {
		t.Errorf("the function is incorrect")
	}

}
