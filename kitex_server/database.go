package database

import (
	"database/sql"
	"fmt"
	"strings"

	api "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/api"
	grader "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/grader"
	"github.com/cloudwego/kitex/pkg/klog"
	_ "github.com/go-sql-driver/mysql"
)

// opens the database and creates the table if it does not exist
func OpenDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "sql6635221:tBf8YkXzVh@tcp(sql6.freesqldatabase.com:3306)/sql6635221")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS students (
			id INT NOT NULL AUTO_INCREMENT,
			num TEXT NOT NULL,
			name TEXT NOT NULL,
			major TEXT NOT NULL,
			gender TEXT NOT NULL,
			grades TEXT NULL,
			PRIMARY KEY (id),
			UNIQUE KEY num_idx (num(255))
		);
	`)
	if err != nil {
		klog.Errorf("Error creating table: %v", err)
		return nil, err
	}

	return db, nil
}

// includes sql command to insert a student
func InsertStudentDB(student *api.InsertStudentRequest) error {
	db, err := OpenDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO students (num, name, major, gender) VALUES (?, ?, ?, ?)", student.Id, student.Name, student.Major, student.Gender)
	if err != nil {
		klog.Errorf("Error inserting table: %v", err)
		return err
	}

	return nil
}

// includes sql command to insert a student
func InsertGradesDB(req *grader.InsertGradeRequest) error {
	db, err := OpenDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	gradeString := strings.Join(req.Grades, ",")
	_, err = db.Exec("UPDATE students SET grades = ? WHERE num = ?", gradeString, fmt.Sprintf("%d", req.StudentId))
	if err != nil {
		klog.Errorf("Error inserting table: %v", err)
		return err
	}

	return nil
}

// check if student exist
func NumExists(num string) (bool, error) {
	db, err := OpenDatabase()
	if err != nil {
		return false, err
	}
	defer db.Close()

	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM students WHERE num = ?)", num).Scan(&exists)
	if err != nil {
		klog.Errorf("Error to check if num exist: %v", err)
		return false, err
	}

	return exists, nil
}

// includes sql command to query a student
func QueryStudentDB(num string) (*api.QueryStudentResponse, error) {
	db, err := OpenDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var resp api.QueryStudentResponse
	err = db.QueryRow("SELECT num, name, major, gender FROM students WHERE num = ?", num).Scan(&resp.Id, &resp.Name, &resp.Major, &resp.Gender)
	if err != nil {
		klog.Errorf("Error Query Table: %v", err)
		return nil, err
	}

	return &resp, nil
}

func GetGradesDB(id int) (*grader.GetCapResponse, string, error) {
	db, err := OpenDatabase()
	if err != nil {
		return nil, "", err
	}
	defer db.Close()

	var resp grader.GetCapResponse
	var grades string
	err = db.QueryRow("SELECT num, name, major, gender, grades FROM students WHERE num = ?", fmt.Sprintf("%d", id)).Scan(&resp.Id, &resp.Name, &resp.Major, &resp.Gender, &grades)
	if err != nil {
		klog.Errorf("Error getting student: %v", err)
		return nil, "", err
	}

	return &resp, grades, nil
}
