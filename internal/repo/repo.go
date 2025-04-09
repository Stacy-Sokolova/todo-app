package repo

import (
	"context"
	"todo-app/internal/entity"
	"todo-app/internal/repo/pgdb"
	"todo-app/pkg/postgres"
)

type Tasks interface {
	Create(ctx context.Context, userId int, task entity.InsertInput) (int, error)
	GetAll(ctx context.Context, userId int) ([]entity.Task, error)
	Update(ctx context.Context, userId int, taskId int, input entity.UpdateInput) (int64, error)
	Delete(ctx context.Context, userId int, taskId int) (int64, error)
}

type Auth interface {
	CreateUser(ctx context.Context, user entity.User) (int, error)
	GetUser(ctx context.Context, input entity.AuthInput) (entity.User, error)
}

type Repository struct {
	Tasks
	Auth
}

func NewRepository(pg *postgres.Postgres) *Repository {
	return &Repository{
		Tasks: pgdb.NewTasksRepo(pg),
		Auth:  pgdb.NewAuthRepo(pg),
	}
}
