package ghosts

import "fpetkovski/monkeyisland/pkg/domain"

type Ghost struct {
	domain.Toy
}

func NewGhost(name string) *Ghost {
	return &Ghost{
		Toy: domain.Toy{
			ID:   generateGhostId(),
			Name: name,
		},
	}
}
