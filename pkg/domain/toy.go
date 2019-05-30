package domain

type Toy struct {
	ID   uint64 `gorm:"primary_key",json:"id"`
	Name string `json:"name"`
}
