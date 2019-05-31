package weapons

import (
	"github.com/jinzhu/gorm"
)

type Repository struct {
	*gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		db,
	}
}

func (repo Repository) GetById(id uint64) *Weapon {
	weapon := new(Weapon)
	err := repo.DB.Find(weapon, id).Error
	if err != nil {
		return nil
	}

	return weapon
}

func (repo Repository) GetAll() []Weapon {
	var weapons []Weapon
	repo.DB.Find(&weapons)

	return weapons
}

func (repo Repository) Create(weapon *Weapon) {
	repo.DB.Create(weapon)
}

func (repo Repository) Update(weapon *Weapon) {
	repo.DB.Save(weapon)
}

func (repo Repository) Delete(weapon *Weapon) {
	repo.DB.Delete(weapon)
}
