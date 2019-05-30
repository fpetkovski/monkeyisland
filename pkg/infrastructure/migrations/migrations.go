package migrations

import (
	"fpetkovski/monkeyisland/domain/cuddly_toys"
	"fpetkovski/monkeyisland/domain/ghosts"
	"fpetkovski/monkeyisland/domain/weapons"
	"fpetkovski/monkeyisland/infrastructure/connection"
)

func Migrate() {
	db := connection.MakeDefaultConnection()
	defer db.Close()

	db.AutoMigrate(
		&weapons.Weapon{},
		&ghosts.Ghost{},
		&cuddly_toys.Monkey{},
		&cuddly_toys.Dog{},
	)
}
