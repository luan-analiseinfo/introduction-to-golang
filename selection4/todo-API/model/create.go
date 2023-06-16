package model

import (
	"../views"
	"fmt"
	"log"
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

func DeleteTodo(name string) error {
	deleteQ, err := conn.Query("DELETE FROM TODO WHERE name = ?", name)
	defer deleteQ.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
