-- name: GetAllTodos :many
SELECT * FROM tb_todos;

-- name: GetAllUnfinishedTodos :many
SELECT * FROM tb_todos WHERE status <> 2;

-- name: GetAllFinishedTodos :many
SELECT * FROM tb_todos WHERE status = 2;

-- name: NewTodo :one
INSERT INTO tb_todos (title) VALUES (?) RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM tb_todos WHERE id = ?;

-- name: FindTodoByTitle :many
SELECT * FROM tb_todos WHERE title LIKE CONCAT('%', ?, '%');

-- name: UpdateTodoTitle :exec
UPDATE tb_todos SET title = ? WHERE id = ?;

-- name: UpdateTodoModifiedDate :exec
UPDATE tb_todos SET modified_at = CURRENT_TIMESTAMP WHERE id = ?;

-- name: CompleteTodo :exec
UPDATE tb_todos SET status = 2 WHERE id = ?;

-- name: ReopenTodo :exec
UPDATE tb_todos SET status = 1 WHERE id = ?;

-- name: StartTodo :exec
UPDATE tb_todos SET status = 1 WHERE id = ?;

