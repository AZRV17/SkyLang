package domain

import "time"

type Exercise struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	FirstVariant  string    `json:"firstVariant"`
	SecondVariant string    `json:"secondVariant"`
	ThirdVariant  string    `json:"thirdVariant"`
	FourthVariant string    `json:"fourthVariant"`
	CorrectAnswer string    `json:"correctAnswer"`
	Difficulty    string    `json:"difficulty"`
	CourseID      uint      `json:"course_id"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
