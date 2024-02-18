package group

import (
	"architecture_go/services/contact/internal/domain/contact"
	"errors"
	"github.com/google/uuid"
	"unicode/utf8"
)

type Group struct {
	ID       uuid.UUID
	Name     string
	Contacts []*contact.Contact
}

func New(name string) (*Group, error) {
	if utf8.RuneCountInString(name) > 250 {
		return nil, errors.New("group name should be less than 250 characters")
	}
	return &Group{ID: uuid.New(), Name: name, Contacts: []*contact.Contact{}}, nil
}

func (g *Group) AddContact(contact *contact.Contact) {
	g.Contacts = append(g.Contacts, contact)
}

func (g *Group) DeleteContact(contactID uuid.UUID) {
	for i, c := range g.Contacts {
		if c.ID == contactID {
			g.Contacts = append(g.Contacts[:i], g.Contacts[i+1:]...)
			break
		}
	}
}

func (g *Group) UpdateContact(contactID uuid.UUID, newContact *contact.Contact) {
	for i, c := range g.Contacts {
		if c.ID == contactID {
			g.Contacts[i] = newContact
			break
		}
	}
}
