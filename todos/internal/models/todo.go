package db

type Todo struct {
	ID uint32 `json:"id" gorm:"primary_key"`
	Content string `json:"content"`
	IsComplete bool `json:"isComplete"`
}