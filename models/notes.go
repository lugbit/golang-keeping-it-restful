// Package models stores data models
package models

// Note model with JSON tags
type Note struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created"`
	UpdatedAt string `json:"updated"`
	Title     string `json:"title"`
	Note      string `json:"note"`
	Colour    string `json:"colour"`
	Archived  bool   `json:"archived"`
	UserID    int    `json:"userID,omitempty"`
}
