package dao

import "github.com/jinzhu/gorm"

type UserDao interface {
}

type UserDaoImpl struct {
	DB *gorm.DB
}
