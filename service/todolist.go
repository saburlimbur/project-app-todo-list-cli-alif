package service

import (
	"errors"
	"todolist_cli/model"
	"todolist_cli/utils"
)

// blueprint contract struct
type TodoService interface {
	AddTodo(title, description, status, priority string) error

	// delete, done, update berdasarkan id
	DeleteTodo(id int) error
	UpdateTodoStatus(id int, status string) error

	ListTodo() ([]model.Todos, error)
}

type todoService struct{}

func NewTodoService() TodoService {
	return &todoService{}
}

func (td *todoService) AddTodo(title, description, status, priority string) error {

	var errMsg string

	switch {
	case title == "" || description == "":
		errMsg += "title and description required\n"
	}

	switch status {
	case "pending", "in-progress", "complete":
	default:
		errMsg += "status invalid\n"
	}

	switch priority {
	case "low", "medium", "high":
	default:
		errMsg += "priority invalid\n"
	}

	if errMsg != "" {
		return errors.New(errMsg)
	}

	// cek exist file
	todos, err := utils.LoadTodoFromFile()
	if err != nil {
		return err
	}

	// generate id otomatis
	newID := 1
	for _, tds := range todos {
		if int(tds.Id) >= newID {
			newID = tds.Id + 1
		}
	}

	newTodo := model.Todos{
		Id:          newID,
		Title:       title,
		Description: description,
		Status:      model.TodoStatus(status),
		Priority:    model.TodoPriority(priority),
	}

	todos = append(todos, newTodo)

	return utils.SavedTodoToFile(todos)
}

func (tds *todoService) ListTodo() ([]model.Todos, error) {
	// load todo di file
	todos, err := utils.LoadTodoFromFile()

	if err != nil {
		return nil, err
	}

	return todos, err
}

func (s *todoService) UpdateTodoStatus(id int, status string) error {
	todos, err := utils.LoadTodoFromFile()
	if err != nil {
		return err
	}

	if status != "pending" && status != "in-progress" && status != "complete" {
		return errors.New("invalid status: must be pending, in-progress, or complete")
	}

	foundData := false

	// loop semua todo
	for in, todo := range todos {
		if todo.Id == id {
			todos[in].Status = model.TodoStatus(status)

			foundData = true
			break
		}
	}

	if !foundData {
		return errors.New("todo with spesific id")
	}

	return utils.SavedTodoToFile(todos)
}

func (s *todoService) DeleteTodo(id int) error {
	todos, err := utils.LoadTodoFromFile()
	if err != nil {
		return err
	}

	// search todo berdasarkan id dan define slice updatedTodos
	foundId := false
	var updatedTodos []model.Todos

	for _, todo := range todos {
		if todo.Id == id {
			foundId = true
			continue // skip todo yang dihapus
		}
		updatedTodos = append(updatedTodos, todo)
	}

	if !foundId {
		return errors.New("todo with specified ID not found")
	}

	return utils.SavedTodoToFile(updatedTodos)
}
