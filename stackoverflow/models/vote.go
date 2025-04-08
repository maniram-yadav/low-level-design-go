package models

type Vote struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"userId"`
	Value  int    `json:"value"`  // +1 or -1
	Type   string `json:"type"`   // "question" or "answer"
	ItemID uint   `json:"itemId"` // ID of question or answer
}
