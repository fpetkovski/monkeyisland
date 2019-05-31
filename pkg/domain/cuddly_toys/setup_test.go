package cuddly_toys

import (
	"fpetkovski/monkeyisland/pkg/infrastructure/connection"
	"github.com/jinzhu/gorm"
)

var dbConnection *gorm.DB
func setupTestcase()  {
	dbConnection = connection.MakeTestConnection()
	dbConnection = dbConnection.Begin()
}

func teardownTestcase()  {
	dbConnection = dbConnection.Rollback()
	dbConnection.Close()
}

