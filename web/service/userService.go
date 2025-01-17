package service

import (
	"log"
	dao "hf/web/dao"
)

func Login(username string, password string) (*dao.User, error){
	log.Printf("用户登录， username:%v password:%v", username, password)
	return dao.Login(username, password)
}