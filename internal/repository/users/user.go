package users

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	DB *sqlx.DB `db:"-" json:"-"`

	Id string `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	FullName string `db:"full_name" json:"full_name"`
	Email string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	Avatar string `db:"avatar" json:"avatar"`
	GoogleId string `db:"google_id" json:"google_id"`
	FacebookId string `db:"facebook_id" json:"facebook_id"`
	GithubId string `db:"github_id" json:"github_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at"`
}

func NewUserMutator(db *sqlx.DB) *User { return &User{DB: db} }