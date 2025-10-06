package task

import (
	"errors"
	"log"
	"sync"
	"time"
)

var ErrNotFound = errors.New("task not found")

type Repo struct {
	mu      sync.RWMutex
	seq     int64
	items   map[int64]*Task
	storage *FileStorage
}

func NewRepo(storage *FileStorage) *Repo {
	repo := &Repo{
		items:   make(map[int64]*Task),
		storage: storage,
	}
	// Загружаем данные при старте
	if err := storage.Load(repo); err != nil {
		// Логируем ошибку, но продолжаем работу с пустым репозиторием
		log.Printf("Failed to load data: %v", err)
	}
	return repo
}

func (r *Repo) saveToDisk() {
	if err := r.storage.Save(r); err != nil {
		log.Printf("Failed to save data: %v", err)
	}
}

// List возвращает все задачи (без пагинации и фильтрации) для обратной совместимости
func (r *Repo) List() []*Task {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*Task, 0, len(r.items))
	for _, t := range r.items {
		out = append(out, t)
	}
	return out
}

// ListWithPagination возвращает задачи с пагинацией и фильтрацией
func (r *Repo) ListWithPagination(opts ListOptions) ([]*Task, int) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Фильтрация
	var filtered []*Task
	for _, t := range r.items {
		if opts.Done != nil && t.Done != *opts.Done {
			continue
		}
		filtered = append(filtered, t)
	}

	total := len(filtered)

	// Пагинация
	start := (opts.Page - 1) * opts.Limit
	if start >= total {
		return []*Task{}, total
	}
	end := start + opts.Limit
	if end > total {
		end = total
	}

	return filtered[start:end], total
}

func (r *Repo) Get(id int64) (*Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.items[id]
	if !ok {
		return nil, ErrNotFound
	}
	return t, nil
}

func (r *Repo) Create(title string) *Task {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.seq++
	now := time.Now()
	t := &Task{ID: r.seq, Title: title, CreatedAt: now, UpdatedAt: now, Done: false}
	r.items[t.ID] = t
	go r.saveToDisk() // Асинхронное сохранение
	return t
}

func (r *Repo) Update(id int64, title string, done bool) (*Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	t, ok := r.items[id]
	if !ok {
		return nil, ErrNotFound
	}
	t.Title = title
	t.Done = done
	t.UpdatedAt = time.Now()
	go r.saveToDisk() // Асинхронное сохранение
	return t, nil
}

func (r *Repo) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[id]; !ok {
		return ErrNotFound
	}
	delete(r.items, id)
	go r.saveToDisk() // Асинхронное сохранение
	return nil
}
