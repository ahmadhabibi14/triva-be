package users

import (
	"errors"
	"time"
	"triva/internal/bootstrap/database"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const TABLE_User = `Users`

type User struct {
	Db *database.Database `db:"-" json:"-"`

	Id 					string 		`db:"id" json:"id"`
	Username 		string 		`db:"username" json:"username"`
	FullName 		string 		`db:"fullName" json:"fullName"`
	Email 			string 		`db:"email" json:"email"`
	Password 		string 		`db:"password" json:"password"`
	Avatar 			string 		`db:"avatar" json:"avatar"`
	GoogleId 		string 		`db:"googleId" json:"googleId"`
	FacebookId	string 		`db:"facebookId" json:"facebookId"`
	GithubId 		string 		`db:"githubId" json:"githubId"`
	CreatedAt 	time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt 	time.Time `db:"updatedAt" json:"updatedAt"`
	DeletedAt 	time.Time `db:"deletedAt" json:"deletedAt"`
}

func NewUserMutator(Db *database.Database) *User {
	return &User{Db: Db}
}

func (u *User) CreateUser() error {
	query := `INSERT INTO ` + TABLE_User +` (id, username, fullName, email, password, createdAt, updatedAt)
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING username`
	
	uid := uuid.New().String()
	if err := u.Db.DB.QueryRowx(query,
		uid, u.Username, u.FullName, u.Email, u.Password,
		time.Now(), time.Now(),
	).StructScan(u); err != nil {
		pgErr, ok := err.(*pq.Error)
		if ok && pgErr.Code == `23505` {
			return errors.New(`email or username is already in use`)
		}
		return err
	}

	return nil
}

func (u *User) FindUsernamePassword() error {
	query := `SELECT username, password FROM ` + TABLE_User + ` WHERE username = $1 LIMIT 1`
	err := u.Db.DB.Get(u, query, u.Username)
	if err != nil {
		return errors.New(`username not found`)
	}

	return nil
}