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
  question: string;
  choices: QuizChoice[];
  answer: QuizChoice;
}

// QuizChoice
export type QuizChoice = {
  id: string;
  text: string;
  correct: boolean;
}