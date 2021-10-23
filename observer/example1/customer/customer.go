package customer

import (
	"fmt"
)

/*
type Customer interface {
	Update(string)
}*/

type Customer struct {
	Name string
}

func (c Customer) Update(item string) {
	fmt.Printf("Hey %s! Your fav %s is back in stock. Go Order Now!", c.Name, item)
}
