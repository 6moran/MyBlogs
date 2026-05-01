package entity

type Like struct {
	ID        int `gorm:"primaryKey"`
	UserID    int `gorm:"uniqueIndex:idx_like;not null"`
	ArticleID int `gorm:"uniqueIndex:idx_like;not null"`
}
