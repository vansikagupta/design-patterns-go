package weapon

import (
	"fmt"
)

type Weapon interface {
	UseWeapon()
}

type SwordWeapon struct {
	Length int
}

func (s SwordWeapon) UseWeapon() {
	//define slaying action
	fmt.Println("Slay with a sword")
}

type GunWeapon struct {
	RangeMM int
}

func (gn GunWeapon) UseWeapon() {
	//define fire action
	fmt.Println("Fire with a gun")
}

type GrenadeWeapon struct {
	ImpactRadius int
}

func (gr GrenadeWeapon) UseWeapon() {
	//define throw action
	fmt.Println("Throw a grenade")
}
