package dto

// DTOに値をセットするタイミングでバリデーションを行う
// カスタムバリデーターがあればここで定義する
type User struct {
	Name                 string `json:"name"  validate:"required,gte=0,lte=30"`
	Email                string `json:"email" validate:"required"`
	Password             string `json:"password" validate:"required,gte=1,lte=50"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,gte=1,lte=50"`
}
