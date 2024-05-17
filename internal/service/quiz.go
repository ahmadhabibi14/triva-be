package service

import (
	"triva/internal/database"
	"triva/internal/repository/quizzes"
)

type QuizService struct {
	Db *database.Database
}

func NewQuizService(Db *database.Database) *QuizService {
	return &QuizService{Db}
}

func (s *QuizService) GetQuizzes() ([]quizzes.Quiz, error) {
	quiz := quizzes.NewQuizMutator(s.Db)
	return quiz.GetQuizzes()
}