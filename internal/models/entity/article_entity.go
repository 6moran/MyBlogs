package entity

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title           string `gorm:"size:255;not null"`
	Summary         string `gorm:"size:255"`
	CoverImage      string `gorm:"size:500"`
	Content         string `gorm:"type:text;not null"`
	Status          int    `gorm:"default:0"`
	LikeCount       int    `gorm:"default:0"`
	ViewCount       int    `gorm:"default:0"`
	CollectionCount int    `gorm:"default:0"`
	CommentCount    int    `gorm:"default:0"`

	//属于哪一个用户
	UserID int `gorm:"not null"`
	User   User

	//属于哪一个分类
	CategoryID int `gorm:"not null"`
	Category   Category

	//拥有的标签
	Tags []Tag `gorm:"many2many:article_tags;"`

	//拥有的评论
	Comments []Comment
}
