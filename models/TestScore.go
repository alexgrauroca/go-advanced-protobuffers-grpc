package models

type TestScore struct {
	TestId    string `json:"test_id"`
	StudentId string `json:"student_id"`
	Ok        int32  `json:"ok"`
	Ko        int32  `json:"ko"`
	Total     int32  `json:"total"`
	Score     int32  `json:"score"`
}
