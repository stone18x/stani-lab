package model

import "testing"

func TestIfCorrectQuestionContainsData(t *testing.T) {
	answer := Answer{ID: 1234, Text: "Cat", Correct: true}

	if answer.ID != 1234 {
		t.Errorf("ID was incorrect, got %v, want %v", answer.ID, 1234)
	}
	if answer.Text != "Cat" {
		t.Errorf("Text was incorrect, got %v, want %v", answer.Text, "Cat")
	}
	if answer.Correct != true {
		t.Errorf("Correct was incorrect, got %v, want %v", answer.Correct, true)
	}
}

func TestIfIncorrectQuestionContainsData(t *testing.T) {
	answer := Answer{ID: 1234, Text: "Cat", Correct: false}

	if answer.ID != 1234 {
		t.Errorf("ID was incorrect, got %v, want %v", answer.ID, 1234)
	}
	if answer.Text != "Cat" {
		t.Errorf("Text was incorrect, got %v, want %v", answer.Text, "Cat")
	}
	if answer.Correct != false {
		t.Errorf("Correct was incorrect, got %v, want %v", answer.Correct, false)
	}
}
