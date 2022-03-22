package models

type Song struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Title    string `json:"title"`
	Singer   string `json:"singer" `
	Writer   string `json:"writer" `
	Director string `json:"director"`
}
