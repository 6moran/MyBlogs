package entity

type Category struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}
