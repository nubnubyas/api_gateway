package database

import (
	"database/sql"
	"fmt"
	"strings"

	// "strconv"

	api "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/api"
	grader "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/grader"
	_ "github.com/go-sql-driver/mysql"
)

// to do : for student_api (add grades into table) and grader_api (get grades from table)

// type Student struct {
// 	Id     string
// 	Name   string
// 	Major  string
// 	Gender string
// 	Grades string
// }

// opens the database and creates the table if it does not exist
func OpenDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "sql6631394:jJ5LZYJR18@tcp(sql6.freesqldatabase.com:3306)/sql6631394")
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
		fmt.Println("Error creating table:", err)
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
		fmt.Println("Error inserting table:", err)
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

	/*
		gradeString := ""
		for _, grade := range req.Grades {
			gradeString += grade
			gradeString += ","
		}
	*/
	gradeString := strings.Join(req.Grades, ",")
	_, err = db.Exec("UPDATE students SET grades = ? WHERE num = ?", gradeString, fmt.Sprintf("%d", req.StudentId))
	if err != nil {
		fmt.Println("Error inserting table:", err)
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
		fmt.Println("Error to check exist:", err)
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
		fmt.Println("Error Query table:", err)
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
		fmt.Println("Error getting Student:", err)
		return nil, "", err
	}

	return &resp, grades, nil
}
