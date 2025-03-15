package service

import (
	"context"
	"todo-app/internal/entity"
	"todo-app/internal/repo"
)

type Tasks interface {
	Create(ctx context.Context, task entity.InsertInput) (int, error)
	GetAll(ctx context.Context) ([]entity.Task, error)
	Update(ctx context.Context, taskId int, input entity.UpdateInput) error
	Delete(ctx context.Context, taskId int) error
}

type Service struct {
	Tasks
}

func NewService(repo *repo.Repository) *Service {
	return &Service{
		Tasks: NewTasksService(repo),
	}
}
