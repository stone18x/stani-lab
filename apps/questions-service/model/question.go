package model

// Question defines its text and answers
type Question struct {
	ID      int      `json:"questionId"`
	Text    string   `json:"text"`
	Answers []Answer `json:"answers"`
}
