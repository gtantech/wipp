package pdm

import (
	"fmt"
	"uuid"
)

type ProjectNotFoundError struct {
	Id      uuid.UUID
	Message string
}

func NewProjectNotFoundError(id uuid.UUID) *ProjectNotFoundError {
	return &ProjectNotFoundError{Id: id, Message: fmt.Sprintf("Project not found for id: %v", id)}
}

func (p *ProjectNotFoundError) Error() string {
	return p.Message
}

type TaskNotFoundError struct {
	Id      uuid.UUID
	Message string
}

func NewTaskNotFoundError(id uuid.UUID) *TaskNotFoundError {
	return &TaskNotFoundError{Id: id, Message: fmt.Sprintf("Task not found for id: %v", id)}
}

func (p *TaskNotFoundError) Error() string {
	return p.Message
}
