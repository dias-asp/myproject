package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	db  *sql.DB
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	log.Println("Hello from Wails!")
	a.initDB() // Инициализация базы данных
	log.Println("Hello from Wails!")
	a.ctx = ctx

}

func (a *App) shutdown(ctx context.Context) {
	if a.db != nil {
		a.db.Close()
	}
}

func (a *App) initDB() {
	var err error
	a.db, err = sql.Open("sqlite3", "./todos.db")
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	_, err = a.db.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			content TEXT NOT NULL,
			done BOOLEAN NOT NULL DEFAULT 0
		)
	`)
	if err != nil {
		log.Fatal("Не удалось создать таблицу:", err)
	}
}
