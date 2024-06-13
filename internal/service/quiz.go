package service

import (
	"triva/internal/bootstrap/database"
	"triva/internal/repository/quizzes"
	"triva/internal/request"
	"triva/internal/response"
)

type QuizService struct {
	Db *database.Database
}

func NewQuizService(Db *database.Database) *QuizService {
	return &QuizService{Db}
}

func (qs *QuizService) GetQuizzes() ([]quizzes.Quiz, error) {
	quiz := quizzes.NewQuizMutator(qs.Db)
	return quiz.GetQuizzes()
}

func (qs *QuizService) CreateQuiz(in request.CreateQuizIn) (out response.CreateQuizOut, err error) {
	quiz := quizzes.NewQuizMutator(qs.Db)
	quiz.Name = in.Name
	quiz.UserId = in.UserId

	err = quiz.Insert()
	if err != nil {
		return
	}

	out.Quiz = quiz
	return
}