package main

import (
	"fmt"
	"time"
)

type Observer interface {
	updateValue(string)
	getID() int
}

type Topic interface {
	register(Observer)
	broadcast()
}

type Item struct {
	observers   []Observer
	name        string
	isAvailable bool
}

func (p *Item) UpdateAvailable() {
	p.isAvailable = !p.isAvailable
	p.broadcast()
}

func (p *Item) register(o Observer) {
	p.observers = append(p.observers, o)
	fmt.Printf("observer %d registered for item %s\n", o.getID(), p.name)
}

func (p *Item) broadcast() {
	for _, o := range p.observers {
		fmt.Printf("product %s broadcasted to observer %d\n", p.name, o.getID())
		o.updateValue(p.name)
	}
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

type EmailClient struct {
	id int
}

func (ec EmailClient) updateValue(name string) {
	fmt.Printf("observer %d received broadcast for product %s\n", ec.id, name)
	ec.sendEmail()
}

func (ec EmailClient) getID() int {
	return ec.id
}

func (ec EmailClient) sendEmail() {
	fmt.Printf("Email Client %d sent email\n", ec.id)
}

func main() {
	i := NewItem("RTX 3080")
	o1 := &EmailClient{
		id: 1,
	}
	o2 := &EmailClient{
		id: 2,
	}

	i.register(o1)
	i.register(o2)

	time.Sleep(3 * time.Second)
	i.UpdateAvailable()
}
