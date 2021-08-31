package main

import (
	"github.com/vansikagupta/design-patterns-go/strategy/vigilante"
	"github.com/vansikagupta/design-patterns-go/strategy/weapon"
)

func main() {
	gunWeapon := weapon.GunWeapon{10}
	alexZ := vigilante.Vigilante{Name: "AlexZ", Weapon: gunWeapon}
	alexZ.Weapon.UseWeapon()

	swordWeapon := weapon.SwordWeapon{10}
	alexZ.Weapon = swordWeapon
	alexZ.Weapon.UseWeapon()

	grenadeWeapon := weapon.GrenadeWeapon{10}
	alexZ.Weapon = grenadeWeapon
	alexZ.Weapon.UseWeapon()

}
