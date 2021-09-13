package models

type User struct {
	ID   uint8  `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}
