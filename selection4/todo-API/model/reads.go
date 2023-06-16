package model

import (
	"../views"
	"net/url"
	"strings"
)

func ReadTodo(params url.Values) ([]views.PostRequest, error) {
	query := "SELECT * FROM TODO WHERE 1 = 1 "
	for k, v := range params {
		query += "AND " + k + " = '" + strings.Join(v, "") + "'"
	}

	rows, err := conn.Query(query)
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
