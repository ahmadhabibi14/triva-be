package request

type RegisterIn struct {
	Username string `json:"username" form:"username" validate:"required,omitempty,min=5"`
	FullName string `json:"full_name" form:"full_name" validate:"required,omitempty,min=5"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
} // @name RegisterIn


type LoginIn struct {
	Username string `json:"username" form:"username" validate:"required,omitempty,min=5"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
} // @name LoginIn