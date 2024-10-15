package dto

type Task struct {
	TaskId      string `json:"TaskId"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Status      string `json:"Status"`
}
