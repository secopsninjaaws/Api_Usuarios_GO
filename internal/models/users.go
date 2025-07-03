package models

type User struct {
	ID        uint `gorm:"primaryKey"`
	Nome      string
	Sobrenome string
}
