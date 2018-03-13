package graph_api

import (
	"context"
	"time"

	"github.com/mitchellh/mapstructure"
)

type Todo struct {
	ID        int
	Text      string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type resolver struct{}

func New() *resolver {
	return &resolver{}
}

func (r *resolver) Mutation_createTodo(ctx context.Context, text string, done *bool) (Todo, error) {
	newTodo := Todo{Text: text, Done: *done}
	return newTodo, DB.Create(&newTodo).Error
}

func (r *resolver) Mutation_updateTodos(ctx context.Context, ids []int, changes map[string]interface{}) ([]Todo, error) {
	todos := make([]Todo, len(ids))
	DB.Where("ID in (?)", ids).Find(&todos)
	for _, todo := range todos {
		if err := mapstructure.Decode(changes, &todo); err != nil {
			return nil, err
		}
		if err := DB.Save(todo).Error; err != nil {
			return todos, err
		}
	}
	return todos, nil
}

func (r *resolver) Mutation_deleteTodos(ctx context.Context, ids []int) (todo []Todo, err error) {
	todos := make([]Todo, len(ids))
	return todos, DB.Where("ID in (?)", ids).Delete(&todos).Error
}

func (r *resolver) Query_todo(ctx context.Context, id int) (todo *Todo, err error) {
	return todo, DB.Where("ID = ?", id).First(todo).Error
}

func (r *resolver) Query_todos(ctx context.Context, status *string) (todos []Todo, err error) {
	if status != nil && (*status == "SHOW_COMPLETED" || *status == "SHOW_ACTIVE") {
		return todos, DB.Where("Done = ?", *status == "SHOW_COMPLETED").Find(&todos).Error
	}
	return todos, DB.Find(&todos).Error
}
