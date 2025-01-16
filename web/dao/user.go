package dao

import (
	"log"
	db "hf/database"
)

type User struct {
	ID			int		`json:"id" description:"user id"`
	UserName	string	`json:"username" description:"user name"`
	EMail		string	`json:"email" description:"email"`
}

func Login(username string, password string) (u User, err error) {
	log.Printf("sqlSelectAllCarouselItemsByCarouselId param->", carouselId)
	var user User
	var id string
	var name string
	var email string
    rows, err := db.DBConnectPool.QueryRowContext(`
select u.id, u.username, u.email 
from users u 
where u.username = $1 and u."password" = md5(concat(u.slat, $2));`, username, password).Scan(&id, &name, &email)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %d\n", id)
		return nil, nil
	case err != nil:
		log.Fatalf("query error: %v\n", err)
		return nil, err
	default:
		log.Printf("username is %q, account email is %s\n", username, email)
	}

    var user User
	user.ID, _ = strconv.Atoi(id)
	user.UserName = name
	user.EMail = email

    return user, err
}