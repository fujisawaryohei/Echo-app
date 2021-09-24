package dto

type User struct {
	Name                 string `json:"name"  validate:"required,gte=0,lte=30"`
	Email                string `json:"email" validate:"email"`
	Password             string `json:"password" validate:"required,gte=1,lte=50"`
	PasswordConfirmation string `json:"password_confirmation" validate:"eqfield=Password,required,gte=1,lte=50"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,gte=1,lte=50"`
}
