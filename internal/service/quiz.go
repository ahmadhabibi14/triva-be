package service

import (
	repository "bwizz/internal/repository/quizzes"
)

type QuizService struct {
	quizRepository *repository.Quiz
}

func NewQuizService(qr *repository.Quiz) *QuizService {
	return &QuizService{quizRepository: qr}
}

func (s *QuizService) GetQuizzes() (quizzes []repository.Quiz, err error) {
	quizzes, err = s.quizRepository.GetQuizzes()

	return
}