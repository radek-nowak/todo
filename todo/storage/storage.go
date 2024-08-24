package storage

import model "github.com/radek-nowak/todo/todo/model"

type Storage interface {
	FindAll() (*model.Tasks, error)
	FindTop(maxItems int) (*model.Tasks, error)
	AddNew(task string)
	Delete(taskId int) error
	DeleteRange(from, to int) error
	Complete(taksId int) error
	Update(taskId int, task string) error
}
