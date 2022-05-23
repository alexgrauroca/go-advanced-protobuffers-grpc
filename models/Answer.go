package models

type Answer struct {
	StudentId  string `json:"student_id"`
	TestId     string `json:"test_id"`
	QuestionId string `json:"question_id"`
	Answer     string `json:"answer"`
	Correct    bool   `json:"correct"`
}
