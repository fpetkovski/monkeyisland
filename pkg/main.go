package main

import (
	"fpetkovski/monkeyisland/pkg/http/handlers/dogs"
	"fpetkovski/monkeyisland/pkg/http/handlers/ghosts"
	"fpetkovski/monkeyisland/pkg/http/handlers/monkeys"
	"fpetkovski/monkeyisland/pkg/http/handlers/weapons"
	"fpetkovski/monkeyisland/pkg/http/middleware"
	"fpetkovski/monkeyisland/pkg/infrastructure/connection"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httplib "net/http"
)

const (
	DOCS_DIR = "/docs/"
)

func main() {
	dbConnection := connection.MakeDefaultConnection()
	defer dbConnection.Close()

	router := mux.NewRouter()
	router.
		PathPrefix(DOCS_DIR).
		Handler(httplib.StripPrefix(DOCS_DIR, httplib.FileServer(httplib.Dir("."+DOCS_DIR))))

	apiRouter := router.PathPrefix("/api").Subrouter()
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

	dogsValidator := middleware.MakeValidator(dogs.ValidationRules())
	dogsHandler := dogs.NewDogsHandler(dbConnection)
	apiRouter.HandleFunc("/dogs", dogsHandler.GetDogs).Methods("GET")
	apiRouter.HandleFunc("/dogs", dogsValidator(dogsHandler.CreateDog)).Methods("POST")
	apiRouter.HandleFunc("/dogs/{id:[0-9]+}", dogsValidator(dogsHandler.UpdateDog)).Methods("PUT")
	apiRouter.HandleFunc("/dogs/{id:[0-9]+}", dogsHandler.DeleteDog).Methods("DELETE")

	monkeysValidator := middleware.MakeValidator(monkeys.ValidationRules())
	monkeysHandler := monkeys.NewMonkeysHandler(dbConnection)
	apiRouter.HandleFunc("/monkeys", monkeysHandler.GetMonkeys).Methods("GET")
	apiRouter.HandleFunc("/monkeys", monkeysValidator(monkeysHandler.CreateMonkey)).Methods("POST")
	apiRouter.HandleFunc("/monkeys/{id:[0-9]+}", monkeysValidator(monkeysHandler.UpdateMonkey)).Methods("PUT")
	apiRouter.HandleFunc("/monkeys/{id:[0-9]+}", monkeysHandler.DeleteMonkey).Methods("DELETE")

	ghostsHandler := ghosts.NewGhostsHandler()
	apiRouter.HandleFunc("/ghosts", ghostsHandler.GetGhosts).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8082"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	err := httplib.ListenAndServe(":80", handler)
	if err != nil {
		panic(err)
	}
}
