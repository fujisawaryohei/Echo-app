package dto

type User struct {
	Name  string `json:"name"  validate:"required,gte=0,lte=30"`
	Email string `json:"email" validate:"email"`
}
