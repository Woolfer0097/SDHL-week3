package models

type Foo struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}
