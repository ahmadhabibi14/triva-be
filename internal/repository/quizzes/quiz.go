package quizzes

import (
	"errors"
	"strings"
	"time"
	"triva/helper"
	"triva/internal/bootstrap/database"
)

const TABLE_Quiz string = `Quiz` 

type Quiz struct {
	Db *database.Database `db:"-" json:"-"`

	Id 				string 					`db:"id" json:"id"`
	Name 			string 					`db:"name" json:"name"`
	UserId 		string 					`db:"userId" json:"userId"`
	CreatedAt time.Time 			`db:"createdAt" json:"createdAt"`
	UpdatedAt time.Time 			`db:"updatedAt" json:"updatedAt"`
	DeletedAt time.Time 			`db:"deletedAt" json:"deletedAt"`
	Questions []QuizQuestion	`db:"-" json:"questions"`
}

func NewQuizMutator(Db *database.Database) *Quiz {
	return &Quiz{Db: Db}
}

func (q *Quiz) GetQuizzes() (quizzes []Quiz, err error) {
	query := `SELECT
		COALESCE(id, '') id, COALESCE(name, '') name
		FROM ` + TABLE_Quiz + ` WHERE userId = ` + q.UserId + `
		ORDER BY name DESC`
	
	err = q.Db.DB.Select(&quizzes, query)
	if err != nil {
		err = errors.New(`failed to get quizzes`)
	}

	return
}

func (q *Quiz) FindById(id string) error {
	query := `SELECT * FROM ` + TABLE_Quiz + ` WHERE id = $1 LIMIT 1`
	err := q.Db.DB.Get(q, query, strings.TrimSpace(id))
	if err != nil {
		return errors.New(`quiz not found`)
	}

	return nil
}

func (q *Quiz) Insert() error {
	query := `INSERT INTO ` + TABLE_Quiz + `
(id, name, userId, createdAt) VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (userId) DO NOTHING
RETURNING id, name, userId, createdAt, updatedAt`

	if err := q.Db.DB.QueryRowx(query,
		helper.RandString(35), q.Name, q.UserId, time.Now(),
	).StructScan(q); err != nil {
		return errors.New(`failed to insert a new quiz`)
	}

	return nil
}

func (q *Quiz) UpdateById() error {
	query := `UPDATE ` + TABLE_Quiz + `
SET name = $1, updatedAt = $2
WHERE id = $3
RETURNING id, name, userId, createdAt, updatedAt`

	if err := q.Db.DB.QueryRowx(query,
		q.Name, time.Now(), q.Id,
	).StructScan(q); err != nil {
		return errors.New(`failed to update quiz`)
	}

	return nil
}