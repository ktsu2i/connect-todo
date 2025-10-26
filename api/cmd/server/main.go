package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ktsu2i/connect-todo/api/gen/todo/v1/todov1connect"
	"github.com/ktsu2i/connect-todo/api/internal/handler"
)

func main() {
	addr := defaultAddr()

	mux := http.NewServeMux()
	todoHandler := handler.NewTodoHandler()
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
