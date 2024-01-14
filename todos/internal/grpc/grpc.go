package grpc

import (
	"context"

	db "todos/internal/models"
	"todos/internal/pb"
	"todos/internal/services"
)

type Server struct {
	pb.UnimplementedTodoServiceServer
}

var TodoService services.TodoService

func (s *Server) GetTodo(ctx context.Context, in *pb.GetTodoRequest) (*pb.Todo, error) {
	todo, err := TodoService.GetTodo(in.Id)

	if err != nil {
		return nil, err
	}

	pbTodo := convertDBTodoToPBTodo(todo)

	return pbTodo, nil
}

func (s *Server) GetAllTodos(ctx context.Context, in *pb.Empty) (*pb.Todos, error) {
	todos, err := TodoService.GetAllTodos()

	if err != nil {
		return nil, err
	}

	pbTodos := convertDBTodoArrayToPBTodos(todos)

	return pbTodos, nil
}

func (s *Server) CreateTodo(ctx context.Context, in *pb.CreateTodoRequest) (*pb.Todo, error) {
	todo, err := TodoService.CreateTodo(in.Content)
	if err != nil {
		return nil, err
	}

	return convertDBTodoToPBTodo(todo), nil
}

func (s *Server) UpdateTodoContent(ctx context.Context, in *pb.UpdateTodoContentRequest) (*pb.Todo, error) {
	todo, err := TodoService.UpdateTodoContent(in.Id, in.Content)
	if err != nil {
		return nil, err
	}

	return convertDBTodoToPBTodo(todo), nil
}

func (s *Server) MarkTodoAsComplete(ctx context.Context, in *pb.UpdateTodoCompletenessRequest) (*pb.Todo, error) {
	todo, err := TodoService.MarkTodoAsCompleted(in.Id)
	if err != nil {
		return nil, err
	}

	return convertDBTodoToPBTodo(todo), nil
}

func (s *Server) MarkTodoAsInComplete(ctx context.Context, in *pb.UpdateTodoCompletenessRequest) (*pb.Todo, error) {
	todo, err := TodoService.MarkTodoAsInCompleted(in.Id)
	if err != nil {
		return nil, err
	}

	return convertDBTodoToPBTodo(todo), nil
}

func (s *Server) DeleteTodo(ctx context.Context, in *pb.DeleteTodoRequest) (*pb.Todo, error) {
	todo, err := TodoService.DeleteTodo(in.Id)
	if err != nil {
		return nil, err
	}

	return convertDBTodoToPBTodo(todo), nil
}

// private

func convertDBTodoToPBTodo(todo *db.Todo) *pb.Todo {
	return &pb.Todo{
		Id:         todo.ID,
		Content:    todo.Content,
		IsComplete: todo.IsComplete,
	}
}

func convertDBTodoArrayToPBTodos(todos []db.Todo) *pb.Todos {
	pbTodos := &pb.Todos{
		Todos: make([]*pb.Todo, 0, len(todos)),
	}
	for _, todo := range todos {
		pbTodo := convertDBTodoToPBTodo(&todo)
		pbTodos.Todos = append(pbTodos.Todos, pbTodo)
	}
	return pbTodos
}
