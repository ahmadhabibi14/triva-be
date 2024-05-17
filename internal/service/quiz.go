package service

import (
	"triva/internal/repository/quizzes"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type QuizService struct {
	db *sqlx.DB
	rd *redis.Client
}

func NewQuizService(db *sqlx.DB, rd *redis.Client) *QuizService {
	return &QuizService{db, rd}
}

func (s *QuizService) GetQuizzes() ([]quizzes.Quiz, error) {
	quiz := quizzes.NewQuizMutator(s.db)
	return quiz.GetQuizzes()
}