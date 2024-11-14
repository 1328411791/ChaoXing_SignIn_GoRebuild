package service

import (
	"log"
)

type IChaoxing interface {
	UserLogin(phone string, password string) (err error)
}

type Chaoxing struct {
}

func UserLogin(phone string, password string) (err error) {
	log.Println("User Login")

	// 查询用户是否存在
	//g.DB().Model("chaoxing").Where("phone =?", phone).Scan(&chaoxing.Chaoxing{})

	//login, err := chaoxing.UserLogin(phone, password)

	return nil
}
