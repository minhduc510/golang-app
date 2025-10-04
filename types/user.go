package types

type User struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6,max=20"`
	Age      int    `validate:"gte=18,lte=65"`
}