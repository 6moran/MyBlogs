package entity

type Tag struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}
