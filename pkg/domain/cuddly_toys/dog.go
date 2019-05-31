package cuddly_toys

import "fpetkovski/monkeyisland/pkg/domain"

type Dog struct {
	CuddlyToy
}

func NewDog(name string, energyLevel int) *Dog {
	return &Dog{
		CuddlyToy: CuddlyToy{
			Toy: domain.Toy{
				Name:name,
			},
			EnergyLevel: energyLevel,
		},
	}
}