package models

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        *string     `gorm:"primarykey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	base.ID = &id
	return nil
}

type Teacher struct {
	BaseModel
	Name      string    `json:"name"`
	TSCNumber string    `json:"tsc_number"`
	Subjects  []Subject `gorm:"many2many:teacher_subjects" json:"subjects"`
}

type Student struct {
	BaseModel
	Name               string    `json:"name"`
	RegistrationNumber string    `json:"registration_number"`
	LevelID            string    `gorm:"foreignKey:ID" json:"level_id"`
	Level              Level     `gorm:"foreignKey:LevelID;references:id" json:"level"`
	Subjects           []Subject `gorm:"many2many:student_subjects" json:"subjects"`

	//Subjects []string `json:"subjects" gorm:"type:text"`
}

type Level struct {
	BaseModel
	LevelName string `json:"level_name"`
}

type Subject struct {
	BaseModel
	Name     string `json:"name"`
	Optional bool   `json:"optional"`
}

type Assignment struct {
	BaseModel
	SubjectID    string    `gorm:"foreignKey:ID" json:"subject_id"`
	Subject      Subject   `gorm:"foreignKey:SubjectID;references:id" json:"subject"`
	TeacherID    string    `gorm:"foreignKey:ID" json:"teacher_id"`
	Teacher      Teacher   `gorm:"foreignKey:TeacherID;references:id" json:"teacher"`
	LevelID      string    `gorm:"foreignKey:ID" json:"level_id"`
	Level        Level     `gorm:"foreignKey:LevelID;references:id" json:"level"`
	Question     string    `json:"question"`
	DueDate      time.Time `gorm:"autoCreateTime:false" json:"due_date"`
	AverageScore float64   `json:"average_score"`
}

type Submission struct {
	BaseModel
	AssignmentID string     `gorm:"foreignKey:ID" json:"assignment_id"`
	Assignment   Assignment `gorm:"foreignKey:AssignmentID" json:"assignment"`
	StudentID    string     `gorm:"foreignKey:ID" json:"student_id"`
	Student      Student    `gorm:"foreignKey:StudentID" json:"student"`
	SubmittedAt  time.Time  `gorm:"autoCreateTime:true" json:"submitted_at"`
	Reviewed     bool       `gorm:"default:false" json:"reviewed"`
	Score        string     `json:"score"`
}
type Message struct {
	BaseModel
	Body      string    `json:"body"`
	ExpiresAt time.Time `gorm:"autoCreateTime:false" json:"expires_at"`
}

type models interface{}

var My_models = [7]models{
	&Teacher{},
	&Student{},
	&Level{},
	&Subject{},
	&Assignment{},
	&Submission{},
	&Message{},
}

func MakeMigrations(db *gorm.DB, models [7]models) {
	for _, model := range models {
		if err := db.AutoMigrate(&model); err != nil {
			log.Fatal(err)
		}
	}
	tx := db.Begin()
	var tables []string
	if err := tx.Table("information_schema.tables").
		Where("table_schema = ?", "public").Pluck("tables", &tables).
		Error; err != nil {
		fmt.Printf("\nError counting created tables: %s", err)
	} else {
		fmt.Printf("\nMigration Complete. %d tables created:\n\n%s", len(tables), tables)
	}
}
