package weapons

import (
	"encoding/json"
	"fpetkovski/monkeyisland/pkg/domain/weapons"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type Handler struct {
	repository weapons.Repository
}

func NewWeaponsHandler(db *gorm.DB) Handler {
	return Handler{
		repository: weapons.NewRepository(db),
	}
}

func (handler Handler) GetWeapons(w http.ResponseWriter, r *http.Request) {
	items := handler.repository.GetAll()
	err := json.NewEncoder(w).Encode(items)
	if err != nil {
		panic(err)
	}
}

func (handler Handler) CreateWeapon(w http.ResponseWriter, r *http.Request) {
	payload := decodePayload(r)

	weapon := weapons.NewWeapon(payload.Name, payload.PowerLevel)
	handler.repository.Create(weapon)

	err := json.NewEncoder(w).Encode(weapon)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler Handler) UpdateWeapon(w http.ResponseWriter, r *http.Request) {
	weaponId := getId(r)

	weapon := handler.repository.GetById(weaponId)
	if weapon == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	payload := decodePayload(r)
	weapon.Name = payload.Name
	weapon.PowerLevel = payload.PowerLevel
	handler.repository.Update(weapon)

	err := json.NewEncoder(w).Encode(weapon)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}

func (handler Handler) DeleteWeapon(w http.ResponseWriter, r *http.Request) {
	weaponId := getId(r)
	weapon := handler.repository.GetById(weaponId)
	if weapon == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler.repository.Delete(weapon)
	w.WriteHeader(http.StatusOK)
}

func getId(r *http.Request) uint64 {
	vars := mux.Vars(r)
	weaponId, _ := strconv.ParseUint(vars["id"], 10, 64)

	return weaponId
}

func decodePayload(r *http.Request) WeaponPayload {
	payload := WeaponPayload{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		panic(err)
	}
	return payload
}