package models

type UserRequest struct {
	Name string `json:"name" validate:"required"`
	DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  string `json:"age,omitempty"`
}
