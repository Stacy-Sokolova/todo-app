package service

import (
	"context"
	"todo-app/internal/entity"
	"todo-app/internal/repo"
)

type TasksService struct {
	repo *repo.Repository
}

func NewTasksService(repo *repo.Repository) *TasksService {
	return &TasksService{
		repo: repo,
	}
}

func (t *TasksService) Create(ctx context.Context, task entity.InsertInput) (int, error) {
	return t.repo.Tasks.Create(ctx, task)
}

func (t *TasksService) GetAll(ctx context.Context) ([]entity.Task, error) {
	return t.repo.Tasks.GetAll(ctx)
}

func (t *TasksService) Update(ctx context.Context, taskId int, input entity.UpdateInput) error {
	return t.repo.Tasks.Update(ctx, taskId, input)
}

func (t *TasksService) Delete(ctx context.Context, taskId int) error {
	return t.repo.Tasks.Delete(ctx, taskId)
}
