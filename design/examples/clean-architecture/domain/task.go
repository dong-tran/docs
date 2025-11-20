package domain

import (
"errors"
"time"
)

// Task represents the core business entity
// This is the innermost layer with no dependencies on other layers
type Task struct {
	ID          int64
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Business rules and validations belong in the domain layer

var (
ErrEmptyTitle         = errors.New("task title cannot be empty")
ErrTitleTooLong       = errors.New("task title cannot exceed 200 characters")
ErrDescriptionTooLong = errors.New("task description cannot exceed 1000 characters")
)

// NewTask creates a new task with validation
func NewTask(title, description string) (*Task, error) {
	if err := ValidateTitle(title); err != nil {
		return nil, err
	}
	if err := ValidateDescription(description); err != nil {
		return nil, err
	}

	now := time.Now()
	return &Task{
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

// ValidateTitle validates the task title
func ValidateTitle(title string) error {
	if title == "" {
		return ErrEmptyTitle
	}
	if len(title) > 200 {
		return ErrTitleTooLong
	}
	return nil
}

// ValidateDescription validates the task description
func ValidateDescription(description string) error {
	if len(description) > 1000 {
		return ErrDescriptionTooLong
	}
	return nil
}

// MarkAsCompleted marks the task as completed
func (t *Task) MarkAsCompleted() {
	t.Completed = true
	t.UpdatedAt = time.Now()
}

// MarkAsIncomplete marks the task as incomplete
func (t *Task) MarkAsIncomplete() {
	t.Completed = false
	t.UpdatedAt = time.Now()
}

// Update updates the task with new values
func (t *Task) Update(title, description string, completed bool) error {
	if err := ValidateTitle(title); err != nil {
		return err
	}
	if err := ValidateDescription(description); err != nil {
		return err
	}

	t.Title = title
	t.Description = description
	t.Completed = completed
	t.UpdatedAt = time.Now()
	return nil
}

// TaskRepository defines the interface for task persistence
// This is defined in the domain layer but implemented in outer layers
type TaskRepository interface {
	Create(task *Task) error
	GetByID(id int64) (*Task, error)
	GetAll() ([]*Task, error)
	Update(task *Task) error
	Delete(id int64) error
}
