package response

import "time"

type UpdateAvatarOut struct {
	Id 					uint64 		`json:"id"`
	Username 		string 		`json:"username"`
	FullName 		string 		`json:"full_name"`
	Email 			string 		`json:"email"`
	AvatarURL		string 		`json:"avatar_url"`
	GoogleId 		string 		`json:"google_id"`
	FacebookId	string 		`json:"facebook_id"`
	GithubId 		string 		`json:"github_id"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
} // @name UpdateAvatarOut