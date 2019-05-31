package cuddly_toys

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_GetById_DogDoesNotExist(t *testing.T) {
	setupTestcase()
	defer teardownTestcase()

	repository := NewDogsRepository(dbConnection)

	dbDog := repository.GetById(123)
	assert.Nil(t, dbDog, "retrieved dog should be nil")
}

func TestRepository_CreateDog(t *testing.T) {
	setupTestcase()
	defer teardownTestcase()

	repository := NewDogsRepository(dbConnection)

	dog := NewDog("Test dog", 42)
	repository.Create(dog)

	dbDog := repository.GetById(dog.ID)
	assert.Equal(t, dbDog.ID, dog.ID, "Retrieved ID does not match stored ID")
	assert.Equal(t, dbDog.Name, "Test dog", "Retrieved name does not match stored name")
	assert.Equal(t, dbDog.EnergyLevel, 42, "Retrieved energy level does not match stored energy level")
}

func TestRepository_GetAllDogs_NoDogsExist(t *testing.T) {
	setupTestcase()
	defer teardownTestcase()

	repository := NewDogsRepository(dbConnection)
	retrievedDogs := repository.GetAll()
	assert.Empty(t, retrievedDogs, "retrievedDogs should be empty")
}

func TestRepository_GetAllDogs_DogsExist(t *testing.T) {
	setupTestcase()
	defer teardownTestcase()

	repository := NewDogsRepository(dbConnection)

	dog1 := NewDog("Test dog 1", 42)
	repository.Create(dog1)

	dog2 := NewDog("Test dog 2", 58)
	repository.Create(dog2)

	retrievedDogs := repository.GetAll()
	assert.Equal(t, retrievedDogs[0].Name, "Test dog 1", "Dog 1 name is incorrect")
	assert.Equal(t, retrievedDogs[0].EnergyLevel, 42, "Dog 1 energy level incorrect")

	assert.Equal(t, retrievedDogs[1].Name, "Test dog 2", "Dog 2 name is incorrect")
	assert.Equal(t, retrievedDogs[1].EnergyLevel, 58, "Dog 2 energy level incorrect")
}

func TestRepository_UpdateDog(t *testing.T) {
	setupTestcase()
	defer teardownTestcase()

	repository := NewDogsRepository(dbConnection)

	dog := NewDog("Test dog", 42)
	repository.Create(dog)

	dog.Name = "Updated dog name"
	dog.EnergyLevel = 55
	repository.Update(dog)

	dbDog := repository.GetById(dog.ID)
	assert.Equal(t, dbDog.ID, dog.ID, "Retrieved ID does not match stored ID")
	assert.Equal(t, dbDog.Name, "Updated dog name", "Retrieved name does not match stored name")
	assert.Equal(t, dbDog.EnergyLevel, 55, "Retrieved energy level does not match stored energy level")
}

func TestRepository_DeleteDog(t *testing.T) {
	setupTestcase()
	defer teardownTestcase()

	repository := NewDogsRepository(dbConnection)

	dog := NewDog("Test dog", 42)
	repository.Create(dog)
	dogId := dog.ID

	dbDog := repository.GetById(dog.ID)
	assert.NotNil(t, dbDog)
	repository.Delete(dog)

	retrievedDog := repository.GetById(dogId)
	assert.Nil(t, retrievedDog, "Retrieved dog should be nil")
}
