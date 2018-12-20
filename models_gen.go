// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package loft

type Echo struct {
	Time   string `json:"time"`
	Format string `json:"format"`
}

type Event struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Loft struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Members  []Member  `json:"members"`
	Tasks    []Task    `json:"tasks"`
	Events   []Event   `json:"events"`
	Requests []Request `json:"requests"`
}

type Member struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NewEvent struct {
	Name string `json:"name"`
}

type NewRequest struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	LoftID  string `json:"loftId"`
}

type NewTask struct {
	Title string `json:"title"`
}

type Request struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}