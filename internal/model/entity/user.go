package entity

import "time"

type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"size:50;not null;unique"`
	Password  string `gorm:"size:255"`
	Nickname  string `gorm:"size:50;unique"`
	Avatar    string `gorm:"size:500"`
	Email     string `gorm:"size:100"`
	GitHubID  int    `gorm:"uniqueIndex"`
	Status    int    `gorm:"default:0"`
	Role      int    `gorm:"default:0"` // 0=visitor, 1=admin
	CreatedAt time.Time

	//拥有的文章
	Articles []Article
	//拥有的评论
	Comments []Comment
}
