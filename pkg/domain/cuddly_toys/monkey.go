package cuddly_toys

import "fpetkovski/monkeyisland/pkg/domain"

type Monkey struct {
	CuddlyToy
}

func NewMonkey(name string, powerLevel int) *Monkey {
	return &Monkey{
		CuddlyToy: CuddlyToy{
			Toy: domain.Toy{
				Name: name,
			},
			EnergyLevel: powerLevel,
		},
	}
}
