package cuddly_toys

import (
	"encoding/json"
	"fpetkovski/monkeyisland/pkg/domain/cuddly_toys"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Handler struct {
	dogsRepository    cuddly_toys.DogsRepository
	monkeysRepository cuddly_toys.MonkeysRepository
}

func NewCuddlyToysHandler(db *gorm.DB) Handler {
	return Handler{
		dogsRepository:    cuddly_toys.NewDogsRepository(db),
		monkeysRepository: cuddly_toys.NewMonkeysRepository(db),
	}
}

func (handler Handler) GetCuddlyToys(w http.ResponseWriter, r *http.Request) {
	dogs := handler.dogsRepository.GetAll()
	monkeys := handler.monkeysRepository.GetAll()

	toys := handler.mergeToys(dogs, monkeys)
	err := json.NewEncoder(w).Encode(toys)
	if err != nil {
		panic(err)
	}
}

func (handler Handler) mergeToys(dogs []cuddly_toys.Dog, monkeys []cuddly_toys.Monkey) []cuddly_toys.CuddlyToy {
	var toys []cuddly_toys.CuddlyToy
	for _, dog := range dogs {
		toys = append(toys, dog.CuddlyToy)
	}
	for _, monkey := range monkeys {
		toys = append(toys, monkey.CuddlyToy)
	}
	return toys
}
