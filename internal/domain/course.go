package domain

import "time"

type Course struct {
	ID          int        `json:"id" gorm:"primaryKey,autoIncrement"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Language    string     `json:"language"`
	Icon        string     `json:"icon"`
	Grates      []Rating   `json:"grates" gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE"`
	Rating      int        `json:"rating" gorm:"default:0"`
	Exercises   []Exercise `json:"exercises" gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE"`
	Lectures    []Lecture  `json:"lectures" gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE"`
	AuthorID    int        `json:"author_id"`
	Author      *User      `json:"author" gorm:"foreignKey:AuthorID"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
}

type Rating struct {
	ID        int       `json:"id" gorm:"primaryKey,autoIncrement"`
	CourseID  int       `json:"course_id"`
	Course    *Course   `json:"course" gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE"`
	Grate     int       `json:"grate"`
	UserID    int       `json:"user_id"`
	User      *User     `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
