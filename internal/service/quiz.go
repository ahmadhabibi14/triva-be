package service

import (
	"bwizz/internal/repository/quizzes"
)

type QuizService struct {
	quizRepository *quizzes.Quiz
}

func NewQuizService(qr *quizzes.Quiz) *QuizService {
	return &QuizService{quizRepository: qr}
}

func (s *QuizService) GetQuizzes() ([]quizzes.Quiz, error) {
	return s.quizRepository.GetQuizzes()
}