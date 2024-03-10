package model

type VideoInfo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Category    string `json:"category"`
	CreatedYear int    `json:"year"`
	Description string `json:"description"`
}
