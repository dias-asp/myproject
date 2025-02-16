package main

import (
	_ "github.com/mattn/go-sqlite3"
)

func (a *App) AddTodo(content string) error {
	_, err := a.db.Exec("INSERT INTO todos (content, done) VALUES (?, 0)", content)
	return err
}

func (a *App) GetTodos() ([]Todo, error) {
	rows, err := a.db.Query("SELECT id, content, done FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Content, &todo.Done)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (a *App) UpdateTodoStatus(id int, done int) error {
	_, err := a.db.Exec("UPDATE todos SET done = ? WHERE id = ?", done, id)
	return err
}

func (a *App) DeleteTodo(id int) error {
	// Удаляем задачу по ID
	_, err := a.db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}

type Todo struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}
