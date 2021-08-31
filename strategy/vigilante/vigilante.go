package vigilante

import (
	"fmt"

	"github.com/vansikagupta/design-patterns-go/strategy/weapon"
)

type Vigilante struct {
	Name   string
	Weapon weapon.Weapon
}

func (v Vigilante) Introduction() {
	//Introduces the Vigilante character
	fmt.Printf("Hey there!! %s is here to save you from Zombies", v.Name)
}
