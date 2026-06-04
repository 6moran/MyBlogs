package entity

import "time"

type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"size:50;not null;unique"`
	Password  string `gorm:"size:255;not null"`
	Nickname  string `gorm:"size:50;unique"`
	Avatar    string `gorm:"size:500"`
	Status    int    `gorm:"default:0"`
	Role      int    `gorm:"default:0"`
	CreatedAt time.Time

	//拥有的文章
	Articles []Article
}
