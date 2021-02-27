package users

import (
	"errors"
	"to-do-go/config"
)

func GetAll() ([]User, error) {
	// users := []User{
	// 	{1, "Mastering Concurrency in Go", "asdasdasd"},
	// 	{2, "Go Design Patterns", "asdasda"},
	// 	{3, "Black Hat Go", "asdasdasds"},
	// }
	var users []User
	db, _ := config.InitDatabase()
	err := db.Find(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func FindByUsername(username string) ([]User, error) {
	var user []User
	db, _ := config.InitDatabase()
	err := db.Where("username = ?", username).Find(&user)
	if err != nil {
		return user, errors.New("error connetcting to database")
	}
	return user, nil
}

func Create(body User) error {
	db, _ := config.InitDatabase()
	_, err := db.Insert(&body)
	if err != nil {
		return err
	}
	return nil
}
