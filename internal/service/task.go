package service

import (
	"context"
	"todo-app/internal/entity"
	"todo-app/internal/repo"
)

type TasksService struct {
	repo repo.Tasks
}

func NewTasksService(repo repo.Tasks) *TasksService {
	return &TasksService{
		repo: repo,
	}
}

func (t *TasksService) Create(ctx context.Context, userId int, task entity.InsertInput) (int, error) {
	return t.repo.Create(ctx, userId, task)
}

func (t *TasksService) GetAll(ctx context.Context, userId int) ([]entity.Task, error) {
	return t.repo.GetAll(ctx, userId)
}

func (t *TasksService) Update(ctx context.Context, userId int, taskId int, input entity.UpdateInput) (int64, error) {
	return t.repo.Update(ctx, userId, taskId, input)
}

func (t *TasksService) Delete(ctx context.Context, userId int, taskId int) (int64, error) {
	return t.repo.Delete(ctx, userId, taskId)
}
