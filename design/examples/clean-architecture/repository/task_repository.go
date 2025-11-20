package repository

import (
"database/sql"
"github.com/dong-tran/docs/clean-architecture-example/domain"
"github.com/jmoiron/sqlx"
)

type TaskRepositoryImpl struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) domain.TaskRepository {
	return &TaskRepositoryImpl{db: db}
}

func (r *TaskRepositoryImpl) Create(task *domain.Task) error {
	query := `
		INSERT INTO tasks (title, description, completed, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`
	result, err := r.db.Exec(query,
task.Title,
task.Description,
task.Completed,
task.CreatedAt,
task.UpdatedAt,
)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	task.ID = id
	return nil
}

func (r *TaskRepositoryImpl) GetByID(id int64) (*domain.Task, error) {
	query := `
		SELECT id, title, description, completed, created_at, updated_at
		FROM tasks
		WHERE id = ?
	`
	var task domain.Task
	err := r.db.Get(&task, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepositoryImpl) GetAll() ([]*domain.Task, error) {
	query := `
		SELECT id, title, description, completed, created_at, updated_at
		FROM tasks
		ORDER BY created_at DESC
	`
	var tasks []*domain.Task
	err := r.db.Select(&tasks, query)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepositoryImpl) Update(task *domain.Task) error {
	query := `
		UPDATE tasks
		SET title = ?, description = ?, completed = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.Exec(query,
task.Title,
task.Description,
task.Completed,
task.UpdatedAt,
task.ID,
)
	return err
}

func (r *TaskRepositoryImpl) Delete(id int64) error {
	query := `DELETE FROM tasks WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
