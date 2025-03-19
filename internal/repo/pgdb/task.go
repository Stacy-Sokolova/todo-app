package pgdb

import (
	"context"
	"fmt"
	"strings"
	"time"
	"todo-app/internal/entity"
	"todo-app/pkg/postgres"
)

const tasksTable = "tasks"

type TasksRepo struct {
	pg *postgres.Postgres
}

func NewTasksRepo(pg *postgres.Postgres) *TasksRepo {
	return &TasksRepo{
		pg: pg,
	}
}

func (r *TasksRepo) Create(ctx context.Context, task entity.InsertInput) (int, error) {
	tx, err := r.pg.Pool.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("postgres.Create - r.pg.Pool.Begin: %v", err)
	}

	var id int
	sql := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", tasksTable)

	err = tx.QueryRow(ctx, sql, task.Title, task.Description).Scan(&id)
	if err != nil {
		tx.Rollback(ctx)
		return 0, fmt.Errorf("postgres.Create - tx.QueryRow: %v", err)
	}

	return id, tx.Commit(ctx)
}

func (r *TasksRepo) GetAll(ctx context.Context) ([]entity.Task, error) {
	sql := fmt.Sprintf("SELECT * FROM %s", tasksTable)

	rows, err := r.pg.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("postgres.GetAll - r.pg.Pool.Query: %v", err)
	}
	defer rows.Close()

	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
		err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.Created_at,
			&task.Updated_at,
		)
		if err != nil {
			return nil, fmt.Errorf("postgres.GetAll - rows.Scan: %v", err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TasksRepo) Update(ctx context.Context, taskId int, input entity.UpdateInput) (int64, error) {
	tx, err := r.pg.Pool.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("postgres.Update - r.pg.Pool.Begin: %v", err)
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Status != nil {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argId))
		args = append(args, *input.Status)
		argId++
	}

	updated_at := time.Now()
	setValues = append(setValues, fmt.Sprintf("updated_at=$%d", argId))
	args = append(args, updated_at)
	argId++

	setQuery := strings.Join(setValues, ", ")
	sql := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", tasksTable, setQuery, argId)
	args = append(args, taskId)
	n, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		tx.Rollback(ctx)
		return 0, fmt.Errorf("postgres.Update - tx.Exec: %v", err)
	}

	return n.RowsAffected(), tx.Commit(ctx)
}

func (r *TasksRepo) Delete(ctx context.Context, taskId int) (int64, error) {
	tx, err := r.pg.Pool.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("postgres.Delete - r.pg.Pool.Begin: %v", err)
	}

	sql := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tasksTable)
	n, err := tx.Exec(ctx, sql, taskId)
	if err != nil {
		tx.Rollback(ctx)
		return 0, fmt.Errorf("postgres.Delete - tx.Exec: %v", err)
	}

	return n.RowsAffected(), tx.Commit(ctx)
}
