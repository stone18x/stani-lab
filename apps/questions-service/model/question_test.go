package model

import "testing"

func TestQuestionWithoutAnswers(t *testing.T) {
	question := Question{ID: 1234, Text: "Is Go great?"}

	if question.ID != 1234 {
		t.Errorf("ID was incorrect, got %v, want %v", question.ID, 1234)
	}
	if question.Text != "Is Go great?" {
		t.Errorf("Text was incorrect, got %v, want %v", question.Text, "Is Go great?")
	}
	if len(question.Answers) != 0 {
		t.Errorf("Expected %v answers", 0)
	}
}

func TestQuestionsWithAnswers(t *testing.T) {
	answers := []Answer{{ID: 1, Text: "No", Correct: false}, {ID: 2, Text: "Yes", Correct: true}}
	question := Question{ID: 1234, Text: "Is Go great?", Answers: answers}

	if question.ID != 1234 {
		t.Errorf("ID was incorrect, got %v, want %v", question.ID, 1234)
	}
	if question.Text != "Is Go great?" {
		t.Errorf("Text was incorrect, got %v, want %v", question.Text, "Is Go great?")
	}
	if len(question.Answers) != 2 {
		t.Errorf("Expected %v answers", 2)
	}
}
