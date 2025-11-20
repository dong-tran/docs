package handler

import (
"net/http"
"strconv"

"github.com/dong-tran/docs/clean-architecture-example/domain"
"github.com/dong-tran/docs/clean-architecture-example/usecase"
"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskUseCase *usecase.TaskUseCase
}

func NewTaskHandler(taskUseCase *usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{
		taskUseCase: taskUseCase,
	}
}

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TaskResponse struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func toResponse(task *domain.Task) TaskResponse {
	return TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		CreatedAt:   task.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   task.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	var req CreateTaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
"error": "invalid request body",
})
	}

	task, err := h.taskUseCase.CreateTask(usecase.CreateTaskInput{
Title:       req.Title,
Description: req.Description,
})
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, toResponse(task))
}

func (h *TaskHandler) GetTask(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
"error": "invalid task id",
})
	}

	task, err := h.taskUseCase.GetTask(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
"error": "task not found",
})
	}

	return c.JSON(http.StatusOK, toResponse(task))
}

func (h *TaskHandler) GetAllTasks(c echo.Context) error {
	tasks, err := h.taskUseCase.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
"error": "failed to retrieve tasks",
})
	}

	responses := make([]TaskResponse, len(tasks))
	for i, task := range tasks {
		responses[i] = toResponse(task)
	}

	return c.JSON(http.StatusOK, responses)
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
"error": "invalid task id",
})
	}

	var req UpdateTaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
"error": "invalid request body",
})
	}

	task, err := h.taskUseCase.UpdateTask(usecase.UpdateTaskInput{
ID:          id,
Title:       req.Title,
Description: req.Description,
Completed:   req.Completed,
})
	if err != nil {
		if err == usecase.ErrTaskNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
"error": "task not found",
})
		}
		return c.JSON(http.StatusBadRequest, map[string]string{
"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, toResponse(task))
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
"error": "invalid task id",
})
	}

	if err := h.taskUseCase.DeleteTask(id); err != nil {
		if err == usecase.ErrTaskNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
"error": "task not found",
})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
"error": "failed to delete task",
})
	}

	return c.NoContent(http.StatusNoContent)
}
