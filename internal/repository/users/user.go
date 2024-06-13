package users

import (
	"errors"
	"time"
	"triva/internal/bootstrap/database"

	"github.com/lib/pq"
)

const TABLE_USER = `users`

type User struct {
	Db *database.Database `db:"-" json:"-"`

	Id 					uint64 		`db:"id" json:"id,omitempty"`
	Username 		string 		`db:"username" json:"username,omitempty"`
	FullName 		string 		`db:"full_name" json:"full_name,omitempty"`
	Email 			string 		`db:"email" json:"email,omitempty"`
	Password 		string 		`db:"password" json:"password,omitempty"`
	AvatarURL		string 		`db:"avatar_url" json:"avatar_url,omitempty"`
	GoogleId 		string 		`db:"google_id" json:"google_id,omitempty"`
	FacebookId	string 		`db:"facebook_id" json:"facebook_id,omitempty"`
	GithubId 		string 		`db:"github_id" json:"github_id,omitempty"`
	CreatedAt 	time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt 	time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt 	time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
} // @name User

func NewUserMutator(Db *database.Database) *User {
	return &User{Db: Db}
}

func (u *User) HideSecrets() {
	u.Password = ``
}

func (u *User) CreateUser() error {
	query := `INSERT INTO ` + TABLE_USER +` (username, full_name, email, password, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, username, full_name, email, avatar_url, created_at, updated_at`
	
	if err := u.Db.DB.QueryRowx(query,
		u.Username, u.FullName, u.Email, u.Password, time.Now(), time.Now(),
	).StructScan(u); err != nil {
		pgErr, ok := err.(*pq.Error)
		if ok && pgErr.Code == `23505` {
			return errors.New(`email or username is already in use`)
		}
		return err
	}

	return nil
}

func (u *User) FindByUsername() error {
	query := `SELECT id, username, full_name, password, email, avatar_url, created_at, updated_at
FROM ` + TABLE_USER + ` WHERE username = $1 LIMIT 1`
	err := u.Db.DB.Get(u, query, u.Username)
	if err != nil {
		return errors.New(`username not found`)
	}

	return nil
}

func (u *User) UpdateAvatarById() error {
	query := `UPDATE ` + TABLE_USER + `
SET avatar_url = $1, updated_at = $2
WHERE id = $3
RETURNING id, username, full_name, email, avatar_url,
	google_id, facebook_id, github_id, created_at, updated_at`

	if err := u.Db.DB.QueryRowx(query,
		u.AvatarURL, time.Now(), u.Id,
	).StructScan(u); err != nil {
		return errors.New(`failed to update avatar`)
	}

	return nil
}