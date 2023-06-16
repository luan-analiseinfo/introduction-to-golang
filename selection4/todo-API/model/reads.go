package model

import (
	"../views"
)

func ReadAll() ([]views.PostRequest, error) {
	rows, err := conn.Query("SELECT * FROM TODO")
	if err != nil {
		return nil, err
	}
	var todos []views.PostRequest
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}
