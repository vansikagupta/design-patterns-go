package main

import (
	customer "github.com/vansikagupta/design-patterns-go/observer/example1/customer"
	invetory "github.com/vansikagupta/design-patterns-go/observer/example1/inventory"
	item "github.com/vansikagupta/design-patterns-go/observer/example1/inventory/item"
)

func main() {
	myInventory := invetory.Inventory{
		Items: make(map[string]item.Item),
	}

	cust1 := customer.Customer{Name: "Alex"}
	oreo := item.Item{Name: "Oreo",
		Price: 100,
		Cust:  cust1,
	}

	myInventory.AddItem(oreo)
}
