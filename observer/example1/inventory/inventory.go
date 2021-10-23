package inventory

import (
	item "github.com/vansikagupta/design-patterns-go/observer/example1/inventory/item"
)

type Inventory struct {
	Items map[string]item.Item
}

func (i *Inventory) AddItem(product item.Item) {
	if _, present := i.Items[product.Name]; !present {
		i.Items[product.Name] = product
	}
	product.Add()
}
