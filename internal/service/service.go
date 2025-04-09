package service

import (
	"context"
	"todo-app/internal/entity"
	"todo-app/internal/repo"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Tasks interface {
	Create(ctx context.Context, userId int, task entity.InsertInput) (int, error)
	GetAll(ctx context.Context, userId int) ([]entity.Task, error)
	Update(ctx context.Context, userId int, taskId int, input entity.UpdateInput) (int64, error)
	Delete(ctx context.Context, userId int, taskId int) (int64, error)
}

type Auth interface {
	CreateUser(ctx context.Context, user entity.User) (int, error)
	GenerateToken(ctx context.Context, input entity.AuthInput) (string, error)
	ParseToken(ctx context.Context, token string) (int, error)
}

type Service struct {
	Tasks
	Auth
}

func NewService(repo *repo.Repository) *Service {
	return &Service{
		Tasks: NewTasksService(repo.Tasks),
		Auth:  NewAuthService(repo.Auth),
	}
}
