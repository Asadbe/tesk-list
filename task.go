package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

//Contact ...
type Contact struct {
	ID       int
	Name     string
	Assighec string
	Deadline time.Time
	Done     bool
}

//ContactList ...
type ContactList struct {
	contacts []Contact
}

// Errors
var (
	ErrorNotFound = errors.New("Contact Not Found")
)

//NewContactList ...
func NewContactList() ContactList {
	contactList := ContactList{}
	contactList.contacts = []Contact{}

	return contactList
}

//Add is function for adding contact to the list
func (cl *ContactList) Add(ct *Contact) (*Contact, error) {
	cl.contacts = append(cl.contacts, *ct)
	return ct, nil
}

//Delete ...
func (cl *ContactList) Delete(id int) error {
	cl.contacts = append(cl.contacts[:id], cl.contacts[id+1:]...)
	return nil

}

//Print ...
func (cl *ContactList) Print(ct *Contact) error {
	for i := 0; i < len(cl.contacts); i++ {

		fmt.Println(cl.contacts[i].Name)
		fmt.Println(cl.contacts[i].Deadline)
		fmt.Println(cl.contacts[i].Done)
		fmt.Println(cl.contacts[i].Assighec + "\n")
	}

	return nil
}

//Update ...
func (cl *ContactList) Update(id int, ct *Contact) error {
	var (
		isfound bool
		orderID int
	)
	for idd, contact := range cl.contacts {

		if contact.ID == id {

			orderID = idd
			isfound = true
		}
	}

	cl.contacts[orderID].Name = ct.Name
	cl.contacts[orderID].Assighec = ct.Assighec
	cl.contacts[orderID].Deadline = ct.Deadline
	cl.contacts[orderID].Done = ct.Done

	if isfound {
		return nil
	}

	return ErrorNotFound
}

// Search ...
func (cl *ContactList) Search(name string, done bool) ([]Contact, error) {
	var contacts []Contact

	for _, cl := range cl.contacts {
		if strings.Contains(strings.ToLower(cl.Name), strings.ToLower(name)) && cl.Done == done && strings.Contains(strings.ToLower(cl.Name), strings.ToLower(name)) {
			contacts = append(contacts, cl)
		} else if strings.Contains(strings.ToLower(cl.Name), strings.ToLower(name)) && cl.Done == done {
			contacts = append(contacts, cl)
		} else if strings.Contains(strings.ToLower(cl.Assighec), strings.ToLower(name)) && cl.Done == done {
			contacts = append(contacts, cl)
		}

	}

	return contacts, nil
}

//Date ...
func (cl *ContactList) Date(date time.Time) ([]Contact, error) {
	var contacts []Contact

	for _, cl := range cl.contacts {
		if (cl.Deadline).Unix() < (date).Unix() {

			contacts = append(contacts, cl)

		}
	}
	return contacts, nil
}

func main() {

	cl := NewContactList()

	cl.Add(&Contact{ID: 1, Name: "write code", Assighec: "rustam", Deadline: time.Now().UTC().AddDate(0, 0, 10), Done: true})
	cl.Add(&Contact{ID: 100, Name: "sign up", Assighec: "Oybek", Deadline: time.Now().UTC().AddDate(0, 0, -10), Done: true})
	cl.Add(&Contact{ID: 2, Name: "sign up", Assighec: "Oybek", Deadline: time.Now().UTC().AddDate(0, 0, -11), Done: true})
	cl.Add(&Contact{ID: 3, Name: "sign up", Assighec: "Oybek", Deadline: time.Now().UTC().AddDate(0, 0, -3), Done: true})
	cl.Add(&Contact{ID: 4, Name: "sign up", Assighec: "Oybek", Deadline: time.Now().UTC().AddDate(0, 0, 10), Done: false})
	cl.Add(&Contact{ID: 5, Name: "sign up", Assighec: "Oybek", Deadline: time.Now().UTC().AddDate(0, 0, 10), Done: true})
	cl.Add(&Contact{ID: 6, Name: "sign up", Assighec: "Oybek", Deadline: time.Now().UTC().AddDate(0, 0, 10), Done: false})
	cl.Add(&Contact{ID: 7, Name: "sign up", Assighec: "Oybek", Deadline: time.Now().UTC().AddDate(0, 0, -25), Done: true})
	cl.Add(&Contact{ID: 8, Name: "sign up", Assighec: "Ulug`bek", Deadline: time.Now().UTC().AddDate(0, 0, 4), Done: true})
	cl.Add(&Contact{ID: 9, Name: "sign up", Assighec: "Oybek", Deadline: time.Now().UTC().AddDate(0, 0, -8), Done: true})

	cl.Delete(1) //  delete function number is id

	err := cl.Update(100, &Contact{ID: 100, Name: "Rustam", Assighec: "Temr", Deadline: time.Now().UTC(), Done: false})

	if err == ErrorNotFound {
		fmt.Println("Contact not found. Id")
	}

	fmt.Println(cl.Search("ulug`bek", true))
	fmt.Println(cl.Date(time.Now()))
	// fmt.Println(cl.contacts[1])

}
