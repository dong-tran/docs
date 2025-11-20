package usecase

import (
"errors"
"github.com/dong-tran/docs/clean-architecture-example/domain"
)

var (
ErrTaskNotFound = errors.New("task not found")
)

type TaskUseCase struct {
	taskRepo domain.TaskRepository
}

func NewTaskUseCase(taskRepo domain.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		taskRepo: taskRepo,
	}
}

type CreateTaskInput struct {
	Title       string
	Description string
}

type UpdateTaskInput struct {
	ID          int64
	Title       string
	Description string
	Completed   bool
}

func (uc *TaskUseCase) CreateTask(input CreateTaskInput) (*domain.Task, error) {
	task, err := domain.NewTask(input.Title, input.Description)
	if err != nil {
		return nil, err
	}

	if err := uc.taskRepo.Create(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (uc *TaskUseCase) GetTask(id int64) (*domain.Task, error) {
	task, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return nil, ErrTaskNotFound
	}
	return task, nil
}

func (uc *TaskUseCase) GetAllTasks() ([]*domain.Task, error) {
	return uc.taskRepo.GetAll()
}

func (uc *TaskUseCase) UpdateTask(input UpdateTaskInput) (*domain.Task, error) {
	task, err := uc.taskRepo.GetByID(input.ID)
	if err != nil {
		return nil, ErrTaskNotFound
	}

	if err := task.Update(input.Title, input.Description, input.Completed); err != nil {
		return nil, err
	}

	if err := uc.taskRepo.Update(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (uc *TaskUseCase) DeleteTask(id int64) error {
	_, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return ErrTaskNotFound
	}

	return uc.taskRepo.Delete(id)
}

func (uc *TaskUseCase) CompleteTask(id int64) (*domain.Task, error) {
	task, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return nil, ErrTaskNotFound
	}

	task.MarkAsCompleted()

	if err := uc.taskRepo.Update(task); err != nil {
		return nil, err
	}

	return task, nil
}
