package main

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {

	cl := NewContactList()
	cl.Add(&Contact{ID: 5, Name: "write code", Assighec: "Doston", Deadline: time.Now().UTC().AddDate(0, 0, 15), Done: false})

	for i := 0; i < len(cl.contacts); i++ {
if cl.contacts[i].ID==5{
	done := cl.contacts[i]

if cl.contacts[i]!=done{
 
	t.Errorf("the function is incorrect")
}
	}}

	
}
