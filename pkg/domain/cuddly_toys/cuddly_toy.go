package cuddly_toys

import "fpetkovski/monkeyisland/pkg/domain"

type CuddlyToy struct {
	domain.Toy
	EnergyLevel int `json:"energy_level"`
}
