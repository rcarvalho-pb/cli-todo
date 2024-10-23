-- name: GetAllTasks :many
SELECT * FROM tb_tasks ORDER BY is_completed, case when completed_at IS NOT NULL THEN completed_at else is_completed end DESC;

-- name: GetAllUnfinishedTasks :many
SELECT * FROM tb_tasks WHERE is_completed = 0;

-- name: GetAllFinishedTasks :many
SELECT * FROM tb_tasks WHERE is_completed = 1 ORDER BY completed_at DESC;

-- name: NewTask :exec
INSERT INTO tb_tasks (title) VALUES (?);

-- name: DeleteTask :exec
DELETE FROM tb_tasks WHERE id = ?;

-- name: FindTaskByTitle :many
SELECT * FROM tb_tasks WHERE title LIKE CONCAT('%', ?, '%') ORDER BY created_at DESC;

-- name: FindTaskById :one
SELECT * FROM tb_tasks WHERE id = ?;

-- name: UpdateTaskTitle :exec
UPDATE tb_tasks SET title = ? WHERE id = ?;

-- name: ToogleTask :exec
UPDATE tb_tasks SET is_completed = ?, completed_at = ? WHERE id = ?;


