package cuddly_toys

import (
	"github.com/jinzhu/gorm"
)

type DogsRepository struct {
	*gorm.DB
}

func NewDogsRepository(db *gorm.DB) DogsRepository {
	return DogsRepository{
		db,
	}
}

func (repo DogsRepository) GetById (id uint64) *Dog {
	dog := new (Dog)
	err := repo.DB.Find(dog, id).Error
	if err != nil {
		return nil
	}

	return dog
}

func (repo DogsRepository) GetAll() []Dog {
	var dogs []Dog
	repo.DB.Find(&dogs)

	return dogs
}

func (repo DogsRepository) Create(dog *Dog) {
	repo.DB.Create(dog)
}

func (repo DogsRepository) Update(dog *Dog) {
	repo.DB.Save(dog)
}

func (repo DogsRepository) Delete(dog *Dog) {
	repo.DB.Delete(dog)
}
