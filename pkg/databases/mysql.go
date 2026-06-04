package databases

import (
	"MyBlogs/internal/model/entity"
	"MyBlogs/pkg/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.MySQLConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database)
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
		&entity.Comment{})
	if err != nil {
		return nil, fmt.Errorf("db.AutoMigrate failed,err:%w", err)
	}
	return db, nil
}
