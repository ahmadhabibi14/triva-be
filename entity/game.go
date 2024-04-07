package entity

type Game struct {
	Id string
	Quiz Quiz
	CurrentQuestion int64
	Code string
}