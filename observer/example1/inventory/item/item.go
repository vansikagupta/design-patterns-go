package item

import (
	customer "github.com/vansikagupta/design-patterns-go/observer/example1/customer"
)

type Item struct {
	Name  string
	Price int
	Cust  customer.Customer
	count int
}

func (i *Item) Add() {
	if i.count == 0 {
		i.NotifyCustomer()
	}
	i.count++
}

func (i Item) NotifyCustomer() {
	i.Cust.Update(i.Name)
}
