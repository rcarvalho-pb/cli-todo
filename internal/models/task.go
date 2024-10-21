package models

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
	"github.com/rcarvalho-pb/cli-todo/pkg/db"
)

type Task struct {
	ID          int64
	Title       string
	IsCompleted bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

func (t *Task) FromTBTask(task db.TbTask) {
	t.ID = task.ID
	t.Title = task.Title
	t.IsCompleted = task.IsCompleted
	t.CreatedAt = task.CreatedAt
	if task.CompletedAt.Valid {
		t.CompletedAt = task.CompletedAt.Time
	}
}

func (t Task) GetAllTasks() ([]*Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result, err := queries.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	var tasks []*Task
	for _, res := range result {
		var task Task
		task.FromTBTask(res)
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (t Task) GetUnfinishedTasks() ([]*Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result, err := queries.GetAllUnfinishedTasks(ctx)
	if err != nil {
		return nil, err
	}

	var tasks []*Task
	for _, res := range result {
		var task Task
		task.FromTBTask(res)
		tasks = append(tasks, &task)
	}

	return tasks, nil

}

func (t Task) GetAllFinishedTasks() ([]*Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result, err := queries.GetAllFinishedTasks(ctx)
	if err != nil {
		return nil, err
	}

	var tasks []*Task
	for _, res := range result {
		var task Task
		task.FromTBTask(res)
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (t Task) UpdateTask(id int64, title string) error {
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

func (t Task) AddTask(title string) (*Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result, err := queries.NewTask(ctx, title)
	if err != nil {
		return nil, err
	}

	var task Task
	task.FromTBTask(result)

	return &task, nil
}

func (t Task) ToggleTask(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	res, err := queries.FindTaskById(ctx, id)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	var toggleTaskParams db.ToogleTaskParams
	toggleTaskParams.ID = id
	if res.IsCompleted {
		toggleTaskParams.IsCompleted = !res.IsCompleted
	} else {
		toggleTaskParams.IsCompleted = !res.IsCompleted
		toggleTaskParams.CompletedAt = time.Now()
	}

	fmt.Printf("%+v\n", toggleTaskParams)

	if err = queries.ToogleTask(ctx, toggleTaskParams); err != nil {
		return err
	}

	return nil
}

func (t Task) DeleteTaks(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	if err := queries.DeleteTask(ctx, id); err != nil {
		return err
	}

	return nil
}

func (t Task) Print(tasks []*Task) {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Título", "Completa?", "Criada Em", "Finalizada Em")

	for _, task := range tasks {
		fmt.Println(task.Title)
		completed := "❌"
		completedAt := "-"
		if task.IsCompleted {
			completedAt = task.CompletedAt.Format("02/01/2006 15:04:05")
			completed = "✅"
		}
		fmt.Println(completedAt)

		table.AddRow(strconv.Itoa(int(task.ID)), task.Title, completed, task.CreatedAt.Format("02/01/2006 15:04:05"), completedAt)
	}

	table.Render()
}
