package domain

import "time"

type Exercise struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	CorrectAnswer string    `json:"correctAnswer"`
	Difficulty    string    `json:"difficulty"`
	CourseID      uint      `json:"courseID"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
