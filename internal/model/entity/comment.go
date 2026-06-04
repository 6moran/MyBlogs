package entity

import "time"

type Comment struct {
	ID        int    `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	LikeCount int    `gorm:"default:0"`
	CreatedAt time.Time

	//属于哪一篇文章
	ArticleID int `gorm:"not null"`

	//回复的哪条评论
	ParentID int `gorm:"default:0"`

	//根评论是哪个
	RootID int `gorm:"default:0"`

	//属于哪个用户
	UserID int
}
