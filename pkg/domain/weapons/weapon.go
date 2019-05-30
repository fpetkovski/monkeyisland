package weapons

import "fpetkovski/monkeyisland/domain"

type Weapon struct {
	domain.Toy
	PowerLevel int `json:"power_level"`
}

func NewWeapon(Name string, PowerLevel int) *Weapon {
	return &Weapon{
		PowerLevel: PowerLevel,
		Toy: domain.Toy{
			Name: Name,
		},
	}
}