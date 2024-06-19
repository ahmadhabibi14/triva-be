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

func (qs *QuizService) GetQuestions(quizId int) ([]quizzes.QuizQuestion, error) {
    quizQuestion := quizzes.NewQuizQuestionMutator(qs.Db)
    return quizQuestion.Select(quizId)
}

func (qs *QuizService) CreateQuestion(in request.CreateQuestionIn) (out response.CreateQuestionOut, err error) {
    quiz := quizzes.NewQuizMutator(qs.Db)
    err = quiz.FindById(in.QuizId)
    if err != nil {
        return
    }

    quizQuestion := quizzes.NewQuizQuestionMutator(qs.Db)
    quizQuestion.Name = in.Name
    quizQuestion.QuizId = in.QuizId
    //hardcode false for now
    quizQuestion.IsUseImage = false

    err = quizQuestion.Insert()
    if err != nil {
        return
    }

    //creating []quizChoice
    quizchoices := make([]quizzes.QuizChoice, len(in.Choices))
    for i, c := range in.Choices{
        var qc quizzes.QuizChoice = quizzes.QuizChoice{
        	QuestionId: quizQuestion.Id,
        	Name:       c.Name,
        	IsCorrect:  c.IsCorrect,
        }
        quizchoices[i] = qc
    }

    //insert options
    quizChoiceMutator := quizzes.NewQuizChoiceMutator(qs.Db)
    result, err := quizChoiceMutator.InsertMany(quizchoices)
    if err != nil {
        return
    }

    quizQuestion.Choices = *result
    out.Question = quizQuestion
    return
}
