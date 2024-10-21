package models

import (
	"context"
	"time"

	"github.com/rcarvalho-pb/cli-todo/pkg/db"
)

type task struct {
	ID          int64
	Title       string
	IsCompleted bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

func (t *task) FromTBTask(task db.TbTask) {
	t.ID = task.ID
	t.Title = task.Title
	t.IsCompleted = task.IsCompleted
	t.CreatedAt = task.CreatedAt
	t.CompletedAt = task.CompletedAt
}

func GetAllTasks(t *task) ([]*task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result, err := queries.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	var tasks []*task
	for _, res := range result {
		var task task
		task.FromTBTask(res)
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (t *task) GetCompletedTasks() ([]*task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result, err := queries.GetAllUnfinishedTasks(ctx)
	if err != nil {
		return nil, err
	}

	var tasks []*task
	for _, res := range result {
		var task task
		task.FromTBTask(res)
		tasks = append(tasks, &task)
	}

	return tasks, nil

}

func (t *task) GetAllFinishedTasks() ([]*task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result, err := queries.GetAllFinishedTasks(ctx)
	if err != nil {
		return nil, err
	}

	var tasks []*task
	for _, res := range result {
		var task task
		task.FromTBTask(res)
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (t *task) UpdateTask(id int64, title string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	if err := queries.UpdateTaskTitle(ctx, db.UpdateTaskTitleParams{
		Title: title,
		ID:    id,
	}); err != nil {
		return err
	}

	return nil

}

func (t *task) AddTask(title string) (*task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result, err := queries.NewTask(ctx, title)
	if err != nil {
		return nil, err
	}

	var task *task
	task.FromTBTask(result)

	return task, nil
}

func (t *task) ToggleTask(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	res, err := queries.FindTaskById(ctx, id)
	if err != nil {
		return err
	}

	var toogleTaskParams db.ToogleTaskParams
	if res.IsCompleted {
		toogleTaskParams.IsCompleted = !res.IsCompleted
	} else {
		toogleTaskParams.IsCompleted = !res.IsCompleted
		toogleTaskParams.CompletedAt = time.Now()
	}

	if err = queries.ToogleTask(ctx, toogleTaskParams); err != nil {
		return err
	}

	return nil
}

func (t *task) DeleteTaks(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	if err := queries.DeleteTask(ctx, id); err != nil {
		return err
	}

	return nil
}
