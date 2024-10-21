CREATE TABLE IF NOT EXISTS tb_tasks(
    id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    is_completed BOOLEAN DEFAULT 0 NOT NULL,
    created_at TIMESTAMP DEFAULT (datetime('now', 'localtime')) NOT NULL,
    completed_at TIMESTAMP,
    CONSTRAINT todo_pk PRIMARY KEY(id)
);

