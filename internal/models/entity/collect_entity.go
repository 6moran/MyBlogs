package entity

type Collect struct {
	ID        int `gorm:"primaryKey"`
	UserID    int `gorm:"uniqueIndex:idx_collect;not null"`
	ArticleID int `gorm:"uniqueIndex:idx_collect;not null"`
}
