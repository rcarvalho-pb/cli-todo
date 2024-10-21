-- name: GetAllTasks :many
SELECT * FROM tb_tasks;

-- name: GetAllUnfinishedTasks :many
SELECT * FROM tb_tasks WHERE is_completed = 0;

-- name: GetAllFinishedTasks :many
SELECT * FROM tb_tasks WHERE is_completed = 1;

-- name: NewTask :one
INSERT INTO tb_tasks (title) VALUES (?) RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tb_tasks WHERE id = ?;

-- name: FindTaskByTitle :many
SELECT * FROM tb_tasks WHERE title LIKE CONCAT('%', ?, '%');

-- name: FindTaskById :one
SELECT * FROM tb_tasks WHERE id = ?;

-- name: UpdateTaskTitle :exec
UPDATE tb_tasks SET title = ? WHERE id = ? RETURNING *;

-- name: ToogleTask :exec
UPDATE tb_tasks SET is_completed = ?, completed_at = ? WHERE id = ? RETURNING *;

