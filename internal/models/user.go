package models

type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Dob  string `json:"dob"`
	Age  int    `json:"age"`
}

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	Dob  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	Dob  string `json:"dob" validate:"required,datetime=2006-01-02"`
}
