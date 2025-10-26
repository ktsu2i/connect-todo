package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ktsu2i/connect-todo/api/gen/todo/v1/todov1connect"
	"github.com/ktsu2i/connect-todo/api/internal/handler"
	"github.com/ktsu2i/connect-todo/api/internal/repo"
	"github.com/ktsu2i/connect-todo/api/internal/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	addr := defaultAddr()
	dsn := databaseURL()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	mux := http.NewServeMux()

	todoRepo := repo.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepo)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	path, h := todov1connect.NewTodoServiceHandler(todoHandler)
	mux.Handle(path, h)

	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server exited: %v", err)
	}
}

func defaultAddr() string {
	if v := os.Getenv("TODO_SERVER_ADDR"); v != "" {
		return v
	}
	return ":8080"
}

func databaseURL() string {
	if v := os.Getenv("TODO_DATABASE_URL"); v != "" {
		return v
	}
	return "postgres://todo_user:todo_password@localhost:5432/todo_app?sslmode=disable"
}
