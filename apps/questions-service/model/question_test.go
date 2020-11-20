package model

import "testing"

func TestQuestion(t *testing.T) {
	const expectedID = 1234
	const expectedText = "Is Go great?"

	question := Question{ID: expectedID, Text: expectedText}

	if question.ID != expectedID {
		t.Errorf("ID was incorrect, got %v, want %v", question.ID, expectedID)
	}

	if question.Text != expectedText {
		t.Errorf("Text was incorrect, got %v, want %v", question.Text, expectedText)
	}
}
