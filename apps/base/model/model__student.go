package model

//StudentModel is a model for student request
type StudentModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//StudentIDModel is a model for student request by ID
type StudentIDModel struct {
	ID int `json:"id"`
}
