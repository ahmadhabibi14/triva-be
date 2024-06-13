package response

import "triva/internal/repository/users"

type RegisterOut struct {
	User 		*users.User	`json:"user" form:"user"`
} // @name RegisterOut

type LoginOut struct {
	SessionKey 	string 			`json:"session_key" form:"session_key"`
	User 				*users.User	`json:"user" form:"user"`
} // @name LoginOut