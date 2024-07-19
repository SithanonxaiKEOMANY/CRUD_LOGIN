package models

import "time"

//type Student struct {
//	ID        uint
//	StudentID *string
//	Firstname *string
//	Lastname  *string
//	Phone     *string
//	Email     *string
//	Password  string
//	Birthday  *time.Time // Use pointer to time.Time to allow NULL
//	Gender    *string
//	Status    *bool
//	Image     *string
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt *time.Time // Use pointer to time.Time to allow NULL
//}

type Student struct {
	ID        uint
	StudentID string
	Firstname string
	Lastname  string
	Phone     string `gorm:"unique"`
	Email     string
	Password  string
	Birthday  time.Time // You can Use pointer to time.Time to allow NULL
	Gender    string
	Status    int
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time // Use pointer to time.Time to allow NULL
	Token     string
}

type Classroom struct {
	ID          uint   `json:"id"`
	ClassName   string `json:"className"`
	ClassYear   int    `json:"class_year"`
	SubjectName string `json:"subject_name"`
}

type StudentClassroom struct {
	ID          uint
	StudentID   uint `gorm:"foreignKey:StudentID;references:ID"`
	ClassroomID uint `gorm:"foreignKey:ClassroomID;references:ID"`
	Classroom   Classroom
	Student     Student
}

type Teacher struct {
	ID        uint
	Phone     string `gorm:"unique"`
	Firstname string
	Lastname  string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
