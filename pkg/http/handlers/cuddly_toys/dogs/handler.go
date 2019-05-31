package dogs

import (
	"encoding/json"
	"fpetkovski/monkeyisland/pkg/domain/cuddly_toys"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type Handler struct {
	repository cuddly_toys.DogsRepository
}

func NewDogsHandler(db *gorm.DB) Handler {
	return Handler{
		repository: cuddly_toys.NewDogsRepository(db),
	}
}

func (handler Handler) GetDogs(w http.ResponseWriter, r *http.Request) {
	items := handler.repository.GetAll()
	err := json.NewEncoder(w).Encode(items)
	if err != nil {
		panic(err)
	}
}

func (handler Handler) CreateDog(w http.ResponseWriter, r *http.Request) {
	payload := decodePayload(r)

	dog := cuddly_toys.NewDog(payload.Name, payload.EnergyLevel)
	handler.repository.Create(dog)

	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(dog)
	if err != nil {
		panic(err)
	}

}

func (handler Handler) UpdateDog(w http.ResponseWriter, r *http.Request) {
	dogId := getId(r)

	dog := handler.repository.GetById(dogId)
	if dog == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	payload := decodePayload(r)
	dog.Name = payload.Name
	dog.EnergyLevel = payload.EnergyLevel
	handler.repository.Update(dog)

	err := json.NewEncoder(w).Encode(dog)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}

func (handler Handler) DeleteDog(w http.ResponseWriter, r *http.Request) {
	dogId := getId(r)
	dog := handler.repository.GetById(dogId)
	if dog == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler.repository.Delete(dog)
	w.WriteHeader(http.StatusOK)
}

func getId(r *http.Request) uint64 {
	vars := mux.Vars(r)
	dogId, _ := strconv.ParseUint(vars["id"], 10, 64)

	return dogId
}

func decodePayload(r *http.Request) DogPayload {
	payload := DogPayload{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		panic(err)
	}
	return payload
}
