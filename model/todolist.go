package model

type TodoStatus string
type TodoPriority string

const (
	StatusPending  TodoStatus = "pending"
	StatusProgress TodoStatus = "in-progress"
	StatusComplete TodoStatus = "complete" // done
)

const (
	PriorityLow    TodoPriority = "low"
	PriorityMedium TodoPriority = "medium"
	PriorityHigh   TodoPriority = "high"
)

type Todos struct {
	Id          int
	Title       string
	Description string
	Status      TodoStatus
	Priority    TodoPriority
}
