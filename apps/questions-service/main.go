package main

import (
	"fmt"

	"com.stani.questionsservice/model"
)

func main() {
	question := model.Question{ID: 123, Text: "How old  are you?"}
	fmt.Printf(question.Text)
}
