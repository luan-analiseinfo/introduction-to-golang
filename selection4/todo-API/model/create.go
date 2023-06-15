package model

import (
	"../views"
	"fmt"
)

func CreateTodo(data *views.PostRequest) error {
	insertQ, err := conn.Query("INSERT INTO TODO VALUES(?, ?)", data.Name, data.Todo)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}
