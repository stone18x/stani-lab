package services

import "com.stani.questionsservice/model"

// QuestionsService contains all functions with regard to questions.
type questionsService struct {
	questions []model.Question
}

// NewQuestionsService creates new instance of QuestionsService.
func NewQuestionsService() *questionsService {
	return &questionsService{questions: createDefaultQuestions()}
}

func createDefaultQuestions() []model.Question {
	var questions []model.Question
	questions = append(questions, model.Question{ID: 1, Text: "Which is the best club in the world?",
		Answers: []model.Answer{{ID: 1, Text: "Tottenham", Correct: false}, {ID: 2, Text: "Arsenal", Correct: true}}})
	questions = append(questions, model.Question{ID: 2, Text: "Are cats great?",
		Answers: []model.Answer{{ID: 1, Text: "Yes", Correct: true}, {ID: 2, Text: "No", Correct: false}}})
	return questions
}

// Get returns all available questions
func (questionsService *questionsService) Get() []model.Question {
	return questionsService.questions
}
