package main

import (
	"fpetkovski/monkeyisland/pkg/http/handlers/cuddly_toys"
	"fpetkovski/monkeyisland/pkg/http/handlers/cuddly_toys/dogs"
	"fpetkovski/monkeyisland/pkg/http/handlers/cuddly_toys/monkeys"
	"fpetkovski/monkeyisland/pkg/http/handlers/ghosts"
	"fpetkovski/monkeyisland/pkg/http/handlers/weapons"
	"fpetkovski/monkeyisland/pkg/http/middleware"
	"fpetkovski/monkeyisland/pkg/infrastructure/connection"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"
	httplib "net/http"
)

func main() {
	dbConnection := connection.MakeDefaultConnection()
	defer dbConnection.Close()

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(
		middleware.JsonResponse,
		middleware.PanicRecovery,
	)

	attachCuddlyToysRoutes(dbConnection, apiRouter)
	attachDogRoutes(dbConnection, apiRouter)
	attachMonkeyRoutes(dbConnection, apiRouter)
	attachWeaponRoutes(dbConnection, apiRouter)
	attachGhostRoutes(apiRouter)

	c := makeCORSHandler()
	err := httplib.ListenAndServe(":80", c.Handler(router))
	if err != nil {
		panic(err)
	}
}

func makeCORSHandler() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8082"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	})
}

func attachGhostRoutes(apiRouter *mux.Router) {
	ghostsHandler := ghosts.NewGhostsHandler()
	apiRouter.HandleFunc("/ghosts", ghostsHandler.GetGhosts).Methods("GET")
}

func attachCuddlyToysRoutes(dbConnection *gorm.DB, apiRouter *mux.Router) {
	cuddlyToysHandler := cuddly_toys.NewCuddlyToysHandler(dbConnection)
	apiRouter.HandleFunc("/cuddly_toys", cuddlyToysHandler.GetCuddlyToys).Methods("GET")
}

func attachMonkeyRoutes(dbConnection *gorm.DB, apiRouter *mux.Router) {
	monkeysValidator := middleware.MakeValidator(monkeys.ValidationRules())
	monkeysHandler := monkeys.NewMonkeysHandler(dbConnection)
	apiRouter.HandleFunc("/cuddly_toys/monkeys", monkeysHandler.GetMonkeys).Methods("GET")
	apiRouter.HandleFunc("/cuddly_toys/monkeys", monkeysValidator(monkeysHandler.CreateMonkey)).Methods("POST")
	apiRouter.HandleFunc("/cuddly_toys/monkeys/{id:[0-9]+}", monkeysValidator(monkeysHandler.UpdateMonkey)).Methods("PUT")
	apiRouter.HandleFunc("/cuddly_toys/monkeys/{id:[0-9]+}", monkeysHandler.DeleteMonkey).Methods("DELETE")
}

func attachDogRoutes(dbConnection *gorm.DB, apiRouter *mux.Router) {
	dogsValidator := middleware.MakeValidator(dogs.ValidationRules())
	dogsHandler := dogs.NewDogsHandler(dbConnection)
	apiRouter.HandleFunc("/cuddly_toys/dogs", dogsHandler.GetDogs).Methods("GET")
	apiRouter.HandleFunc("/cuddly_toys/dogs", dogsValidator(dogsHandler.CreateDog)).Methods("POST")
	apiRouter.HandleFunc("/cuddly_toys/dogs/{id:[0-9]+}", dogsValidator(dogsHandler.UpdateDog)).Methods("PUT")
	apiRouter.HandleFunc("/cuddly_toys/dogs/{id:[0-9]+}", dogsHandler.DeleteDog).Methods("DELETE")
}

func attachWeaponRoutes(dbConnection *gorm.DB, apiRouter *mux.Router) {
	weaponValidator := middleware.MakeValidator(weapons.ValidationRules())
	weaponsHandler := weapons.NewWeaponsHandler(dbConnection)
	apiRouter.HandleFunc("/weapons", weaponsHandler.GetWeapons).Methods("GET")
	apiRouter.HandleFunc("/weapons", weaponValidator(weaponsHandler.CreateWeapon)).Methods("POST")
	apiRouter.HandleFunc("/weapons/{id:[0-9]+}", weaponValidator(weaponsHandler.UpdateWeapon)).Methods("PUT")
	apiRouter.HandleFunc("/weapons/{id:[0-9]+}", weaponsHandler.DeleteWeapon).Methods("DELETE")
}
