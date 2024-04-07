package entity

type Quiz struct {
	Id string
	Name string
	Questions []QuizQuestion
}

type QuizQuestion struct {
	Id string
	Name string
	Choices []QuizChoice
}

type QuizChoice struct {
	Id string
	Name string
	Correct bool
}