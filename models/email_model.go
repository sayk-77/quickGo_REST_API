package models

type Email struct {
	Question   string `json:"question"`
	Customer   string `json:"customer"`
	Email      string `json:"email"`
	NameEmploy string `json:"nameEmploy"`
	Solution   string `json:"solution"`
}
