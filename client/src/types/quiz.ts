// Quiz
export type Quiz = {
  id: string;
  name: string;
  description: string;
  questions: QuizQuestion[];
}

// QuizQuestion
export type QuizQuestion = {
  id: string;
  name: string;
  choices: QuizChoice[];
}

// QuizChoice
export type QuizChoice = {
  id: string;
  name: string;
  correct: boolean;
}

export type Player = {
  id: string;
  name: string;
}