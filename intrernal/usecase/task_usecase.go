package usecase

import (
	"errors"
	"task-tracker/intrernal/entity"
	"task-tracker/intrernal/repository"
	"time"
)

// TaskUsecase содержит бизнес-логику для работы с задачами
type TaskUsecase struct {
	repo repository.TaskRepository
}

// NewTaskUsecase создает новый экземпляр TaskUsecase
func NewTaskUsecase(repo repository.TaskRepository) *TaskUsecase {
	return &TaskUsecase{repo: repo}
}

// AddTask добавляет новую задачу
func (uc *TaskUsecase) AddTask(description string) (int, error) {
	tasks, err := uc.repo.Load()
	if err != nil {
		return 0, err
	}

	newTask := entity.Task{
		ID:          generateTaskID(tasks),
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)
	err = uc.repo.Save(tasks)
	if err != nil {
		return 0, err
	}

	return newTask.ID, nil
}

// UpdateTask обновляет описание задачи
func (uc *TaskUsecase) UpdateTask(taskID int, newDescription string) error {
	tasks, err := uc.repo.Load()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now()
			return uc.repo.Save(tasks)
		}
	}

	return errors.New("task not found")
}

// DeleteTask удаляет задачу
func (uc *TaskUsecase) DeleteTask(taskID int) error {
	tasks, err := uc.repo.Load()
	if err != nil {
		return err
	}

	newTasks := []entity.Task{}
	for _, task := range tasks {
		if task.ID != taskID {
			newTasks = append(newTasks, task)
		}
	}

	if len(newTasks) == len(tasks) {
		return errors.New("task not found")
	}

	return uc.repo.Save(newTasks)
}

// MarkTaskStatus изменяет статус задачи
func (uc *TaskUsecase) MarkTaskStatus(taskID int, status string) error {
	tasks, err := uc.repo.Load()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			return uc.repo.Save(tasks)
		}
	}

	return errors.New("task not found")
}

// ListTasks возвращает список задач
func (uc *TaskUsecase) ListTasks(status string) ([]entity.Task, error) {
	tasks, err := uc.repo.Load()
	if err != nil {
		return nil, err
	}

	filteredTasks := []entity.Task{}
	if status != "" {
		for _, task := range tasks {
			if task.Status == status {
				filteredTasks = append(filteredTasks, task)
			}
		}
	} else {
		filteredTasks = tasks
	}

	return filteredTasks, nil
}

// Генерация уникального ID
func generateTaskID(tasks []entity.Task) int {
	if len(tasks) == 0 {
		return 1
	}

	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}
