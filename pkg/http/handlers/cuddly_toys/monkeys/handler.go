package monkeys

import (
	"encoding/json"
	"fpetkovski/monkeyisland/pkg/domain/cuddly_toys"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type Handler struct {
	repository cuddly_toys.MonkeysRepository
}

func NewMonkeysHandler(db *gorm.DB) Handler {
	return Handler{
		repository: cuddly_toys.NewMonkeysRepository(db),
	}
}

func (handler Handler) GetMonkeys(w http.ResponseWriter, r *http.Request) {
	items := handler.repository.GetAll()
	err := json.NewEncoder(w).Encode(items)
	if err != nil {
		panic(err)
	}
}

func (handler Handler) CreateMonkey(w http.ResponseWriter, r *http.Request) {
	payload := decodePayload(r)

	monkey := cuddly_toys.NewMonkey(payload.Name, payload.EnergyLevel)
	handler.repository.Create(monkey)

	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(monkey)
	if err != nil {
		panic(err)
	}

}

func (handler Handler) UpdateMonkey(w http.ResponseWriter, r *http.Request) {
	monkeyId := getId(r)

	monkey := handler.repository.GetById(monkeyId)
	if monkey == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	payload := decodePayload(r)
	monkey.Name = payload.Name
	monkey.EnergyLevel = payload.EnergyLevel
	handler.repository.Update(monkey)

	err := json.NewEncoder(w).Encode(monkey)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}

func (handler Handler) DeleteMonkey(w http.ResponseWriter, r *http.Request) {
	monkeyId := getId(r)
	monkey := handler.repository.GetById(monkeyId)
	if monkey == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler.repository.Delete(monkey)
	w.WriteHeader(http.StatusOK)
}

func getId(r *http.Request) uint64 {
	vars := mux.Vars(r)
	monkeyId, _ := strconv.ParseUint(vars["id"], 10, 64)

	return monkeyId
}

func decodePayload(r *http.Request) MonkeyPayload {
	payload := MonkeyPayload{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		panic(err)
	}
	return payload
}