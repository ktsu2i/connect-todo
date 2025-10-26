package handler

import (
	"context"
	"time"

	connect "connectrpc.com/connect"
	v1 "github.com/ktsu2i/connect-todo/api/gen/todo/v1"
	"github.com/ktsu2i/connect-todo/api/gen/todo/v1/todov1connect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TodoHandler struct{}

var _ todov1connect.TodoServiceHandler = (*TodoHandler)(nil)

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{}
}

func (h *TodoHandler) ListTodos(
	ctx context.Context,
	_ *connect.Request[v1.ListTodosRequest],
) (*connect.Response[v1.ListTodosResponse], error) {
	now := timestamppb.New(time.Now())
	resp := &v1.ListTodosResponse{
		Todos: []*v1.Todo{
			{
				Id:        1,
				Title:     "Sample todo",
				Done:      false,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
	}
	return connect.NewResponse(resp), nil
}

func (h *TodoHandler) GetTodo(
	ctx context.Context,
	req *connect.Request[v1.GetTodoRequest],
) (*connect.Response[v1.GetTodoResponse], error) {
	now := timestamppb.New(time.Now())
	resp := &v1.GetTodoResponse{
		Todo: &v1.Todo{
			Id:        req.Msg.GetId(),
			Title:     "Dummy todo",
			Done:      false,
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	return connect.NewResponse(resp), nil
}

func (h *TodoHandler) CreateTodo(
	ctx context.Context,
	req *connect.Request[v1.CreateTodoRequest],
) (*connect.Response[v1.CreateTodoResponse], error) {
	now := timestamppb.New(time.Now())
	resp := &v1.CreateTodoResponse{
		Todo: &v1.Todo{
			Id:        42,
			Title:     req.Msg.GetTitle(),
			Done:      false,
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	return connect.NewResponse(resp), nil
}

func (h *TodoHandler) UpdateTodo(
	ctx context.Context,
	req *connect.Request[v1.UpdateTodoRequest],
) (*connect.Response[v1.UpdateTodoResponse], error) {
	now := timestamppb.New(time.Now())
	resp := &v1.UpdateTodoResponse{
		Todo: &v1.Todo{
			Id:        req.Msg.GetId(),
			Title:     req.Msg.GetTitle(),
			Done:      req.Msg.GetDone(),
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	return connect.NewResponse(resp), nil
}

func (h *TodoHandler) DeleteTodo(
	ctx context.Context,
	_ *connect.Request[v1.DeleteTodoRequest],
) (*connect.Response[v1.DeleteTodoResponse], error) {
	return connect.NewResponse(&v1.DeleteTodoResponse{}), nil
}
