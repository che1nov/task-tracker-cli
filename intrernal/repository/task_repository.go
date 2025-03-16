package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"task-tracker/intrernal/entity"
)

// TaskRepository управляет хранением задач
type TaskRepository struct {
	filePath string
	mu       sync.Mutex
}

// NewTaskRepository создает новый экземпляр TaskRepository
func NewTaskRepository(filePath string) *TaskRepository {
	return &TaskRepository{filePath: filePath}
}

// Load загружает задачи из файла
func (r *TaskRepository) Load() ([]entity.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var tasks []entity.Task

	if _, err := os.Stat(r.filePath); os.IsNotExist(err) {
		return tasks, nil
	}

	data, err := ioutil.ReadFile(r.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return tasks, nil
}

// Save сохраняет задачи в файл
func (r *TaskRepository) Save(tasks []entity.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	err = ioutil.WriteFile(r.filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
