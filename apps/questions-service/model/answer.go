package model

// Answer is part of question and contains id, text and defines whether answer is correct
type Answer struct {
	ID      int    `json:"answerId"`
	Text    string `json:"text"`
	Correct bool   `json:"correct"`
}
