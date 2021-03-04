package models

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Job  string `json:"job"`
}
