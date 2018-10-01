package resolvers

import (
	"context"
	"fmt"

	"github.com/mitchellh/mapstructure"

	"github.com/tanema/vGoDoRPQl/api/data"
	"github.com/tanema/vGoDoRPQl/api/gql"
)

type todoresolver struct{}

// New will return a new todo resolver
func New() gql.Resolvers {
	return &todoresolver{}
}

func (r *todoresolver) Mutation_createTodo(ctx context.Context, text string, done *bool) (data.Todo, error) {
	newTodo := data.Todo{Text: text, Done: *done}
	fmt.Printf("Creating TODO: %v", text)
	return newTodo, data.DB.Create(&newTodo).Error
}

func (r *todoresolver) Mutation_updateTodos(ctx context.Context, ids []int, changes map[string]interface{}) ([]data.Todo, error) {
	todos := make([]data.Todo, len(ids))
	data.DB.Where("ID in (?)", ids).Find(&todos)

	fmt.Printf("Updating TODOs: %v", todos)
	for _, todo := range todos {
		if err := mapstructure.Decode(changes, &todo); err != nil {
			return nil, err
		}
		if err := data.DB.Save(todo).Error; err != nil {
			return todos, err
		}
	}
	return todos, nil
}

func (r *todoresolver) Mutation_deleteTodos(ctx context.Context, ids []int) (todo []data.Todo, err error) {
	todos := make([]data.Todo, len(ids))
	return todos, data.DB.Where("ID in (?)", ids).Delete(&todos).Error
}

func (r *todoresolver) Query_todo(ctx context.Context, id int) (todo *data.Todo, err error) {
	return todo, data.DB.Where("ID = ?", id).First(todo).Error
}

func (r *todoresolver) Query_todos(ctx context.Context, status *string) (todos []data.Todo, err error) {
	if status != nil && (*status == "SHOW_COMPLETED" || *status == "SHOW_ACTIVE") {
		return todos, data.DB.Where("Done = ?", *status == "SHOW_COMPLETED").Find(&todos).Error
	}
	return todos, data.DB.Find(&todos).Error
}
