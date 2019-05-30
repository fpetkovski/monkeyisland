package migrations

import (
	"fpetkovski/monkeyisland/pkg/domain/cuddly_toys"
	"fpetkovski/monkeyisland/pkg/domain/ghosts"
	"fpetkovski/monkeyisland/pkg/domain/weapons"
	"fpetkovski/monkeyisland/pkg/infrastructure/connection"
	"github.com/jinzhu/gorm"
)

func Migrate() {
	db := connection.MakeDefaultConnection()
	defer db.Close()

	migrate(db)
}

func MigrateTest() {
	db := connection.MakeTestConnection()
	defer db.Close()

	migrate(db)
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&weapons.Weapon{},
		&ghosts.Ghost{},
		&cuddly_toys.Monkey{},
		&cuddly_toys.Dog{},
	)
}
