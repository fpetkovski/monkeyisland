package main

import (
	"fpetkovski/monkeyisland/http/handlers/weapons"
	"fpetkovski/monkeyisland/http/middleware"
	"fpetkovski/monkeyisland/infrastructure/connection"
	"github.com/gorilla/mux"
	httplib "net/http"
)

func main() {
	dbConnection := connection.MakeDefaultConnection()
	defer dbConnection.Close()

	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Use(
		middleware.JsonResponse,
		middleware.PanicRecovery,
	)

	weaponValidator := middleware.MakeValidator(weapons.ValidationRules())
	weaponsHandler := weapons.NewWeaponsHandler(dbConnection)
	apiRouter.HandleFunc("/weapons", weaponsHandler.GetWeapons).Methods("GET")
	apiRouter.HandleFunc("/weapons", weaponValidator(weaponsHandler.CreateWeapon)).Methods("POST")
	apiRouter.HandleFunc("/weapons/{id:[0-9]+}", weaponValidator(weaponsHandler.UpdateWeapon)).Methods("PUT")
	apiRouter.HandleFunc("/weapons/{id:[0-9]+}", weaponsHandler.DeleteWeapon).Methods("DELETE")

	err := httplib.ListenAndServe(":80", r)
	if err != nil {
		panic(err)
	}
}
