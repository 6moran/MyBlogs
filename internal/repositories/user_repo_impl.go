package repositories

import "gorm.io/gorm"

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &UserRepoImpl{db: db}
}

func (u *UserRepoImpl) CreateUser() {

}

func (u *UserRepoImpl) DeleteUser() {

}

func (u *UserRepoImpl) EditUser() {

}

func (u *UserRepoImpl) QueryUserLimit() {

}
