package main

import (
	"fmt"
	"log"

	"github.com/miceremwirigi/PLP-Family-Homework-Manager-Django/go-version/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(postgres.
		Open("host=localhost dbname=homework_manager user=homework_manager password=homework_manager port=5432 sslmode=disable timezone=Africa/Nairobi"),
		&gorm.Config{})
	if err != nil {
		log.Println("\n\nInitializing Database Failed")
		log.Fatal(err)
	} else {
		log.Println("\n\nInitializing Database Success")
	}
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;`)

	teacher := models.Teacher{
		Name:      "test_teacher",
		TSCNumber: "test_tsc_number",
	}
	db.Find(&teacher, "name = ?", "test_teacher")
	if teacher.ID == nil {
		db.Create(&teacher)
	} else {
		fmt.Printf("Teacher %s already exists", teacher.Name)
	}

	level := models.Level{
		LevelName: "test_level",
	}
	db.Find(&level, "level_name = ?", "test_level")
	if level.ID == nil {
		db.Create(&level)
	} else {
		fmt.Printf("Level %s already exists", teacher.Name)
	}

	subject := models.Subject{
		Name:     "test_subject",
		Optional: false,
	}
	db.Find(&subject, "name = ?", "test_subject")
	if subject.ID == nil {
		db.Create(&subject)
	} else {
		fmt.Printf("Subject %s already exists", teacher.Name)
	}

	student := models.Student{
		Name:               "test_student",
		RegistrationNumber: "test_registration_number",
		LevelID:            *level.ID,
		Subjects:           []models.Subject{subject},
	}
	db.Find(&student, "name = ?", "test_student")
	if student.ID == nil {
		db.Create(&student)
	} else {
		fmt.Printf("Student %s already exists", teacher.Name)
	}

	assignment := models.Assignment{
		SubjectID: *subject.ID,
		TeacherID: *teacher.ID,
		LevelID:   *level.ID,
		Question:  "test_Question",
	}
	db.Find(&assignment, "question = ?", "test_question")
	if assignment.ID == nil {
		db.Create(&assignment)
	} else {
		fmt.Printf("Assignment %s already exists", teacher.Name)
	}

	submission := models.Submission{
		AssignmentID: *assignment.ID,
		StudentID:    *student.ID,
		Reviewed:     false,
	}

	db.Find(&submission, "assignment_id = ?", assignment.ID)
	if submission.ID == nil {
		db.Create(&submission)
	} else {
		fmt.Printf("Submision %s already exists", submission.Assignment.Question)
	}

	message := models.Message{
		Body: "This is a test message",
	}
	db.Find(&message, "body = ?", "This is a test message")
	if message.ID == nil {
		db.Create(&message)
	} else {
		fmt.Printf("Message %s already exists", submission.Assignment.Question)
	}

}
