package todos

import (
	"fmt"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func getTasks(c iris.Context) {
	todos, _ := GetAll(c.GetCookie("id"))
	c.StatusCode(iris.StatusOK)
	c.JSON(todos, context.
		JSON{Indent: "  "})
	return
}

func getTaskByTaskId(c iris.Context) {
	todos, err := GetTaskById(c.URLParam("task_id"))
	if err != nil {
		c.StatusCode(iris.StatusInternalServerError)
		c.JSON(iris.Map{
			"message": err.Error(),
		})
		return
	}
	//Check if todos nill
	if len(todos) == 0 {
		c.StatusCode(iris.StatusNotFound)
		c.JSON(iris.Map{
			"message": "Doesn't Exist",
		})
		return
	}
	//Check user_id == user_id cookies
	if !checkUserId(todos[0].UserId, c.GetCookie("id")) {
		c.StatusCode(iris.StatusUnauthorized)
		c.JSON(iris.Map{
			"message": "Unauthorized",
		})
		return
	}
	c.StatusCode(iris.StatusOK)
	c.JSON(todos, context.
		JSON{Indent: "  "})
	return
}

func getTaskByTaskStatus(c iris.Context) {
	todos, err := GetTaskByStatus(c.URLParam("task_status"), c.GetCookie("id"))
	if err != nil {
		c.StatusCode(iris.StatusInternalServerError)
		c.JSON(iris.Map{
			"message": err.Error(),
		})
		return
	}
	//Check if todos nill
	if len(todos) == 0 {
		c.StatusCode(iris.StatusNotFound)
		c.JSON(iris.Map{
			"message": "Doesn't Exist",
		})
		return
	}
	//Check user_id == user_id cookies
	if !checkUserId(todos[0].UserId, c.GetCookie("id")) {
		c.StatusCode(iris.StatusUnauthorized)
		c.JSON(iris.Map{
			"message": "Unauthorized",
		})
		return
	}
	c.StatusCode(iris.StatusOK)
	c.JSON(todos, context.
		JSON{Indent: "  "})
	return
}

func createTask(c iris.Context) {
	var todo Todo

	err := c.ReadJSON(&todo)

	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.JSON(iris.Map{
			"message": err.Error(),
		})
		return
	}

	//Check If task_id empty
	if todo.Id != 0 {
		c.StatusCode(iris.StatusBadRequest)
		c.JSON(iris.Map{
			"message": "Fill Title and Task Only",
		})
		return
	}

	user_id, _ := strconv.ParseInt(c.GetCookie("id"), 10, 64)
	todo.UserId = user_id
	todo.Status = "Belum"
	//Create
	if err := Create(todo); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.JSON(iris.Map{
			"message": err.Error(),
		})
		return
	}
	c.StatusCode(iris.StatusCreated)
	c.JSON(iris.Map{
		"message": "Task Created",
	})
	return
}

func deleteTask(c iris.Context) {
	todos, err := GetTaskById(c.URLParam("task_id"))
	if err != nil {
		c.StatusCode(iris.StatusInternalServerError)
		c.JSON(iris.Map{
			"message": err.Error(),
		})
		return
	}
	//Check if todos nill
	if len(todos) == 0 {
		c.StatusCode(iris.StatusNotFound)
		c.JSON(iris.Map{
			"message": "Doesn't Exist",
		})
		return
	}
	//Check user_id == user_id cookies
	if !checkUserId(todos[0].UserId, c.GetCookie("id")) {
		c.StatusCode(iris.StatusUnauthorized)
		c.JSON(iris.Map{
			"message": "Unauthorized",
		})
		return
	}
	if err := Delete(todos[0].Id); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.JSON(iris.Map{
			"message": err.Error(),
		})
		return
	}
	c.StatusCode(iris.StatusOK)
	c.JSON(iris.Map{
		"message": "Task Deleted",
	})
	return
}

func taskFinish(c iris.Context) {
	todos, err := GetTaskById(c.URLParam("task_id"))
	if err != nil {
		c.StatusCode(iris.StatusInternalServerError)
		c.JSON(iris.Map{
			"message": err.Error(),
		})
		return
	}
	//Check if todos nill
	if len(todos) == 0 {
		c.StatusCode(iris.StatusNotFound)
		c.JSON(iris.Map{
			"message": "Doesn't Exist",
		})
		return
	}
	//Check user_id == user_id cookies
	if !checkUserId(todos[0].UserId, c.GetCookie("id")) {
		c.StatusCode(iris.StatusUnauthorized)
		c.JSON(iris.Map{
			"message": "Unauthorized",
		})
		return
	}
	if err := UpdateStatus(todos[0].Id, "Selesai"); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.JSON(iris.Map{
			"message": err.Error(),
		})
		return
	}
	c.StatusCode(iris.StatusOK)
	c.JSON(iris.Map{
		"message": "Task Update",
	})
	return
}

func taskOngoing(c iris.Context) {
	todos, err := GetTaskById(c.URLParam("task_id"))
	if err != nil {
		c.StatusCode(iris.StatusInternalServerError)
		c.JSON(iris.Map{
			"message": err.Error(),
		})
		return
	}
	//Check if todos nill
	if len(todos) == 0 {
		c.StatusCode(iris.StatusNotFound)
		c.JSON(iris.Map{
			"message": "Doesn't Exist",
		})
		return
	}
	//Check user_id == user_id cookies
	if !checkUserId(todos[0].UserId, c.GetCookie("id")) {
		c.StatusCode(iris.StatusUnauthorized)
		c.JSON(iris.Map{
			"message": "Unauthorized",
		})
		return
	}
	if err := UpdateStatus(todos[0].Id, "Belum"); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.JSON(iris.Map{
			"message": err.Error(),
		})
		return
	}
	c.StatusCode(iris.StatusOK)
	c.JSON(iris.Map{
		"message": "Task Update",
	})
	return
}

func checkUserId(user_id int64, cookies_id string) bool {
	//Check If userID == loginID
	if fmt.Sprint(user_id) != cookies_id {
		return false
	} else {
		return true
	}
}
