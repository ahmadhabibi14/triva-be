package request

type CreateQuizIn struct {
	Name 		string	`db:"name" json:"name" validate:"required"`
	UserId 	uint64	`db:"user_id" json:"user_id"`
} // @name CreateQuizIn