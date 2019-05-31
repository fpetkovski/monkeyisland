package cuddly_toys

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_GetById_MonkeyDoesNotExist(t *testing.T) {
	setupTestcase()
	defer teardownTestcase()

	repository := NewMonkeysRepository(dbConnection)

	dbMonkey := repository.GetById(123)
	assert.Nil(t, dbMonkey, "retrieved monkey should be nil")
}

func TestRepository_CreateMonkey(t *testing.T) {
	setupTestcase()
	defer teardownTestcase()

	repository := NewMonkeysRepository(dbConnection)

	monkey := NewMonkey("Test monkey", 42)
	repository.Create(monkey)

	dbMonkey := repository.GetById(monkey.ID)
	assert.Equal(t, dbMonkey.ID, monkey.ID, "Retrieved ID does not match stored ID")
	assert.Equal(t, dbMonkey.Name, "Test monkey", "Retrieved name does not match stored name")
	assert.Equal(t, dbMonkey.EnergyLevel, 42, "Retrieved energy level does not match stored energy level")
}

func TestRepository_GetAllMonkeys_NoMonkeysExist(t *testing.T) {
	setupTestcase()
	defer teardownTestcase()

	repository := NewMonkeysRepository(dbConnection)
	retrievedMonkeys := repository.GetAll()
	assert.Empty(t, retrievedMonkeys, "retrievedMonkeys should be empty")
}

func TestRepository_GetAllMonkeys_MonkeysExist(t *testing.T) {
	setupTestcase()
	defer teardownTestcase()

	repository := NewMonkeysRepository(dbConnection)

	monkey1 := NewMonkey("Test monkey 1", 42)
	repository.Create(monkey1)

	monkey2 := NewMonkey("Test monkey 2", 58)
	repository.Create(monkey2)

	retrievedMonkeys := repository.GetAll()
	assert.Equal(t, retrievedMonkeys[0].Name, "Test monkey 1", "Monkey 1 name is incorrect")
	assert.Equal(t, retrievedMonkeys[0].EnergyLevel, 42, "Monkey 1 energy level incorrect")

	assert.Equal(t, retrievedMonkeys[1].Name, "Test monkey 2", "Monkey 2 name is incorrect")
	assert.Equal(t, retrievedMonkeys[1].EnergyLevel, 58, "Monkey 2 energy level incorrect")
}

func TestRepository_UpdateMonkey(t *testing.T) {
	setupTestcase()
	defer teardownTestcase()

	repository := NewMonkeysRepository(dbConnection)

	monkey := NewMonkey("Test monkey", 42)
	repository.Create(monkey)

	monkey.Name = "Updated monkey name"
	monkey.EnergyLevel = 55
	repository.Update(monkey)

	dbMonkey := repository.GetById(monkey.ID)
	assert.Equal(t, dbMonkey.ID, monkey.ID, "Retrieved ID does not match stored ID")
	assert.Equal(t, dbMonkey.Name, "Updated monkey name", "Retrieved name does not match stored name")
	assert.Equal(t, dbMonkey.EnergyLevel, 55, "Retrieved energy level does not match stored energy level")
}


func TestRepository_DeleteMonkey(t *testing.T) {
	setupTestcase()
	defer teardownTestcase()

	repository := NewMonkeysRepository(dbConnection)

	monkey := NewMonkey("Test monkey", 42)
	repository.Create(monkey)
	monkeyId := monkey.ID

	dbMonkey := repository.GetById(monkey.ID)
	assert.NotNil(t, dbMonkey)
	repository.Delete(monkey)

	retrievedMonkey := repository.GetById(monkeyId)
	assert.Nil(t, retrievedMonkey, "Retrieved monkey should be nil")
}