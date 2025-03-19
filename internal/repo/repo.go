package repo

import (
	"context"
	"todo-app/internal/entity"
	"todo-app/internal/repo/pgdb"
	"todo-app/pkg/postgres"
)

type Tasks interface {
	Create(ctx context.Context, task entity.InsertInput) (int, error)
	GetAll(ctx context.Context) ([]entity.Task, error)
	Update(ctx context.Context, taskId int, input entity.UpdateInput) (int64, error)
	Delete(ctx context.Context, taskId int) (int64, error)
}

type Repository struct {
	Tasks
}

func NewRepository(pg *postgres.Postgres) *Repository {
	return &Repository{
		Tasks: pgdb.NewTasksRepo(pg),
	}
}
