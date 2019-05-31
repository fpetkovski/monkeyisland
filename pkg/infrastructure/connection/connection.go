package connection

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

// Instantiates a connection using the database parameters passed in the environment
func MakeDefaultConnection() *gorm.DB {
	return makeConnection(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)
}

// Instantiates a connection using the database parameters passed in the environment
func MakeTestConnection() *gorm.DB {
	return makeConnection(
		os.Getenv("TEST_DB_HOST"),
		os.Getenv("TEST_DB_PORT"),
		os.Getenv("TEST_DB_USER"),
		os.Getenv("TEST_DB_PASS"),
		os.Getenv("TEST_DB_NAME"),
	)
}

// Instantiates a connection using using the given parameters
func makeConnection(host string, port string, username string, password string, database string) *gorm.DB {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)

	connection, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	return connection
}
