package todos

import (
	"to-do-go/config"
)

func GetAll(user_id string) ([]Todo, error) {
	// todos := []Todo{
	// 	{1, 31, "Mastering Concurrency in Go", "asdasdasd", "Selesai"},
	// 	{2, 31, "Go Design Patterns", "asdasda", "Selesai"},
	// 	{3, 32, "Black Hat Go", "asdasdasds", "Selesai"},
	// }
	var todos []Todo
	db, _ := config.InitDatabase()
	err := db.Where("user_id = ?", user_id).Find(&todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func GetTaskById(task_id string) ([]Todo, error) {
	var todos []Todo
	db, _ := config.InitDatabase()
	err := db.Where("id = ?", task_id).Find(&todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func GetTaskByStatus(status string) ([]Todo, error) {
	var todos []Todo
	db, _ := config.InitDatabase()
	err := db.Where("status = ?", status).Find(&todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func Create(body Todo) error {
	db, _ := config.InitDatabase()
	_, err := db.Insert(&body)
	if err != nil {
		return err
	}
	return nil
}

func UpdateStatus(task_id int64, status string) error {
	var todo Todo
	db, _ := config.InitDatabase()
	db.ID(task_id).Get(&todo)
	todo.Status = status
	_, err := db.ID(task_id).Update(&todo)
	if err != nil {
		return err
	}
	return nil
}

func Delete(task_id int64) error {
	var todo Todo
	db, _ := config.InitDatabase()
	db.ID(task_id).Delete(&todo)
	db.ID(task_id).Get(&todo)
	db.ID(task_id).Delete(&todo)
	return nil
}
