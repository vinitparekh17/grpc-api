package main

import (
	"context"

	pb "github.com/vinitparekh17/grpc-api/golang-service/protobufs"
)

var todoList []*pb.Todo = []*pb.Todo{
	{
		Id:          1,
		Title:       "Walk",
		Description: "Go to Walk",
	},
	{
		Id:          2,
		Title:       "Dance",
		Description: "Dance for fun",
	},
	{
		Id:          3,
		Title:       "Clean",
		Description: "Clean your stuff",
	},
	{
		Id:          4,
		Title:       "Eat",
		Description: "Have a good meal",
	},
	{
		Id:          5,
		Title:       "Study",
		Description: "Go to Study",
	},
}

func (s *Server) GetTodos(ctx context.Context, req *pb.GetTodosRequest) (*pb.GetTodosResponse, error) {

	return &pb.GetTodosResponse{
		Todos: todoList,
	}, nil
}

func (s *Server) AddTodo(ctx context.Context, req *pb.AddTodoRequest) (*pb.AddTodoResponse, error) {

	newId := todoList[len(todoList)-1].Id + 1

	todoList = append(todoList, &pb.Todo{
		Id:          newId,
		Title:       req.Title,
		Description: req.Description,
	})

	return &pb.AddTodoResponse{
		Id: newId,
	}, nil
}

func (s *Server) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {

	index := -1
	for i, item := range todoList {
		if item.Id == req.Id {
			index = i
			break
		}
	}

	// If the item was found, remove it from the slice
	if index != -1 {
		todoList = append(todoList[:index], todoList[index+1:]...)
	}

	return &pb.DeleteTodoResponse{
		Success: true,
	}, nil
}
