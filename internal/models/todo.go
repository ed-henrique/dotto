package models

type ToDo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`

	Attributes map[string]any `json:"attributes"`
}

var ToDos = []ToDo{
	{0, "a", true, nil},
}
