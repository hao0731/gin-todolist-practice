package model

type Todo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required,max=20"`
	Description string `json:"description" binding:"max=100"`
}

type CreateTodoResponse struct {
	Todo Todo `json:"todo"`
}

type GetTodosResponse struct {
	Todos []Todo `json:"todos"`
}

type GetTodoResponse struct {
	Todo Todo `json:"todo"`
}
