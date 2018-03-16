package data

import "time"

// Todo is a single item that requires completion
type Todo struct {
	ID        int
	Text      string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
