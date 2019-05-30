package weapons

import (
	"fpetkovski/monkeyisland/infrastructure/connection"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

var dbConnection *gorm.DB
func setupTestcase()  {
	dbConnection = connection.MakeDefaultConnection()
	dbConnection = dbConnection.Begin()
}

func teardownTestcase()  {
	dbConnection = dbConnection.Rollback()
	dbConnection.Close()
}

func TestRepository_GetById_WeaponDoesNotExist(t *testing.T) {
	setupTestcase()
	defer teardownTestcase()

	repository := NewRepository(dbConnection)

	dbWeapon := repository.GetById(123)
	assert.Nil(t, dbWeapon, "retrieved weapon should be nil")
}

func TestRepository_CreateWeapon(t *testing.T) {
	setupTestcase()
	defer teardownTestcase()

	repository := NewRepository(dbConnection)

	weapon := NewWeapon("Test weapon", 42)
	repository.Create(weapon)

	dbWeapon := repository.GetById(weapon.ID)
	assert.Equal(t, dbWeapon.ID, weapon.ID, "Retrieved ID does not match stored ID")
	assert.Equal(t, dbWeapon.Name, "Test weapon", "Retrieved name does not match stored name")
	assert.Equal(t, dbWeapon.PowerLevel, 42, "Retrieved power level does not match stored power level")
}

func TestRepository_GetAllWeapons_NoWeaponsExist(t *testing.T) {
	setupTestcase()
	defer teardownTestcase()

	repository := NewRepository(dbConnection)
	retrievedWeapons := repository.GetAll()
	assert.Empty(t, retrievedWeapons, "retrievedWeapons should be empty")
}

func TestRepository_GetAllWeapons_WeaponsExist(t *testing.T) {
	setupTestcase()
	defer teardownTestcase()

	repository := NewRepository(dbConnection)

	weapon1 := NewWeapon("Test weapon 1", 42)
	repository.Create(weapon1)

	weapon2 := NewWeapon("Test weapon 2", 58)
	repository.Create(weapon2)

	retrievedWeapons := repository.GetAll()
	assert.Equal(t, retrievedWeapons[0].Name, "Test weapon 1", "Weapon 1 name is incorrect")
	assert.Equal(t, retrievedWeapons[0].PowerLevel, 42, "Weapon 1 power level incorrect")

	assert.Equal(t, retrievedWeapons[1].Name, "Test weapon 2", "Weapon 2 name is incorrect")
	assert.Equal(t, retrievedWeapons[1].PowerLevel, 58, "Weapon 2 power level incorrect")
}

func TestRepository_UpdateWeapon(t *testing.T) {
	setupTestcase()
	defer teardownTestcase()

	repository := NewRepository(dbConnection)

	weapon := NewWeapon("Test weapon", 42)
	repository.Create(weapon)

	weapon.Name = "Updated weapon name"
	weapon.PowerLevel = 55
	repository.Update(weapon)

	dbWeapon := repository.GetById(weapon.ID)
	assert.Equal(t, dbWeapon.ID, weapon.ID, "Retrieved ID does not match stored ID")
	assert.Equal(t, dbWeapon.Name, "Updated weapon name", "Retrieved name does not match stored name")
	assert.Equal(t, dbWeapon.PowerLevel, 55, "Retrieved power level does not match stored power level")
}


func TestRepository_DeleteWeapon(t *testing.T) {
	setupTestcase()
	defer teardownTestcase()

	repository := NewRepository(dbConnection)

	weapon := NewWeapon("Test weapon", 42)
	repository.Create(weapon)
	weaponId := weapon.ID

	dbWeapon := repository.GetById(weapon.ID)
	assert.NotNil(t, dbWeapon)
	repository.Delete(weapon)

	retrievedWeapon := repository.GetById(weaponId)
	assert.Nil(t, retrievedWeapon, "Retrieved weapon should be nil")
}