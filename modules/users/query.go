package users

import "to-do-go/config"

func GetAll() []User {
	// users := []User{
	// 	{1, "Mastering Concurrency in Go", "asdasdasd"},
	// 	{2, "Go Design Patterns", "asdasda"},
	// 	{3, "Black Hat Go", "asdasdasds"},
	// }
	var users []User
	db, _ := config.InitDatabase()
	db.Find(&users)
	return users
}
