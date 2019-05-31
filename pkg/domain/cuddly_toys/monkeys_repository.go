package cuddly_toys

import (
	"github.com/jinzhu/gorm"
)

type MonkeysRepository struct {
	*gorm.DB
}

func NewMonkeysRepository(db *gorm.DB) MonkeysRepository {
	return MonkeysRepository{
		db,
	}
}

func (repo MonkeysRepository) GetById(id uint64) *Monkey {
	monkey := new(Monkey)
	err := repo.DB.Find(monkey, id).Error
	if err != nil {
		return nil
	}

	return monkey
}

func (repo MonkeysRepository) GetAll() []Monkey {
	var monkeys []Monkey
	repo.DB.Find(&monkeys)

	return monkeys
}

func (repo MonkeysRepository) Create(monkey *Monkey) {
	repo.DB.Create(monkey)
}

func (repo MonkeysRepository) Update(monkey *Monkey) {
	repo.DB.Save(monkey)
}

func (repo MonkeysRepository) Delete(monkey *Monkey) {
	repo.DB.Delete(monkey)
}
