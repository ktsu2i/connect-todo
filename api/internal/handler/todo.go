package handler

import (
	"context"
	"errors"
	"time"

	connect "connectrpc.com/connect"
	v1 "github.com/ktsu2i/connect-todo/api/gen/todo/v1"
	"github.com/ktsu2i/connect-todo/api/gen/todo/v1/todov1connect"
	"github.com/ktsu2i/connect-todo/api/internal/model"
	"github.com/ktsu2i/connect-todo/api/internal/usecase"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TodoHandler struct {
	uc usecase.ITodoUsecase
}

var _ todov1connect.TodoServiceHandler = (*TodoHandler)(nil)

func NewTodoHandler(uc usecase.ITodoUsecase) *TodoHandler {
	return &TodoHandler{uc: uc}
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
	todo, err := h.uc.Create(req.Msg.GetTitle())
	if err != nil {
		if errors.Is(err, usecase.ErrEmptyTitle) {
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	resp := &v1.CreateTodoResponse{
		Todo: todoToProto(todo),
	}
	return connect.NewResponse(resp), nil
}

func (h *TodoHandler) UpdateTodo(
	ctx context.Context,
	req *connect.Request[v1.UpdateTodoRequest],
) (*connect.Response[v1.UpdateTodoResponse], error) {
	todo, err := h.uc.Update(
		req.Msg.GetId(),
		req.Msg.GetTitle(),
		req.Msg.GetDone(),
	)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	if todo == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("todo not found"))
	}
	resp := &v1.UpdateTodoResponse{
		Todo: todoToProto(todo),
	}
	return connect.NewResponse(resp), nil
}

func (h *TodoHandler) DeleteTodo(
	ctx context.Context,
	_ *connect.Request[v1.DeleteTodoRequest],
) (*connect.Response[v1.DeleteTodoResponse], error) {
	return connect.NewResponse(&v1.DeleteTodoResponse{}), nil
}

func todoToProto(todo *model.Todo) *v1.Todo {
	if todo == nil {
		return nil
	}
	return &v1.Todo{
		Id:        todo.ID,
		Title:     todo.Title,
		Done:      todo.Done,
		CreatedAt: timestamppb.New(todo.CreatedAt),
		UpdatedAt: timestamppb.New(todo.UpdatedAt),
	}
}
