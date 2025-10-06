package task

import (
	"encoding/json"
	"os"
	"sync"
)

type FileStorage struct {
	filename string
	mu       sync.RWMutex
}

func NewFileStorage(filename string) *FileStorage {
	return &FileStorage{filename: filename}
}

func (s *FileStorage) Load(repo *Repo) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data, err := os.ReadFile(s.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Файл не существует - это нормально при первом запуске
		}
		return err
	}

	var tasks []*Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return err
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.items = make(map[int64]*Task)
	for _, task := range tasks {
		repo.items[task.ID] = task
		if task.ID > repo.seq {
			repo.seq = task.ID
		}
	}

	return nil
}

func (s *FileStorage) Save(repo *Repo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	repo.mu.RLock()
	tasks := make([]*Task, 0, len(repo.items))
	for _, task := range repo.items {
		tasks = append(tasks, task)
	}
	repo.mu.RUnlock()

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filename, data, 0644)
}
