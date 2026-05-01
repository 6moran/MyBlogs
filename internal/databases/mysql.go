package databases

import (
	"MyBlogs/internal/config"
	"MyBlogs/internal/models/entity"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MySQL.User,
		config.MySQL.Password,
		config.MySQL.Host,
		config.MySQL.Port,
		config.MySQL.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, fmt.Errorf("gorm.Open() failed,err:%w", err)
	}

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Tag{},
		&entity.Category{},
		&entity.Article{},
		&entity.Like{},
		&entity.Collect{},
		&entity.Comment{},
		&entity.Message{})
	if err != nil {
		return nil, fmt.Errorf("db.AutoMigrate failed,err:%w", err)
	}
	return db, nil
}
