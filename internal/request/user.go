package request

import "mime/multipart"

type UpdateAvatarIn struct {
	Avatar 			multipart.Form 		`json:"avatar" form:"avatar"`
} // @name UpdateAvatarIn