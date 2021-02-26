package users

import "to-do-go/config"

func GetAll() []User {
	db, _ := config.InitDatabase()
	var users []User
	db.Find(&users)
	return users
}
