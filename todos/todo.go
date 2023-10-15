package todo

import (
	"errors"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	list := *t

	err := list.CheckConstraints(index)

	if err != nil {
		return err
	}

	list[index-1].CompletedAt = time.Now()
	list[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {
	list := *t

	err := list.CheckConstraints(index)

	if err != nil {
		return err
	}

	*t = append(list[:index-1], list[index:]...)

	return nil
}

func (t *Todos) Clear() error {
	list := *t

	if len(list) == 0 {
		return errors.New("Todo list is already empty..")
	}

	*t = list[:0]

	return nil
}

func (t *Todos) Total() int {
	total := 0

	for _, item := range *t {
		if !item.Done {
			total++
		}
	}

	return total
}
