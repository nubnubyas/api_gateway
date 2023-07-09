package database

import (
	"database/sql"
	"fmt"

	api "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/api"
	_ "github.com/go-sql-driver/mysql"
)

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
			gender TEXT NOT NULL,
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

	_, err = db.Exec("INSERT INTO students (num, name, gender) VALUES (?, ?, ?)", student.Num, student.Name, student.Gender)
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
	err = db.QueryRow("SELECT num, name, gender FROM students WHERE num = ?", num).Scan(&resp.Num, &resp.Name, &resp.Gender)
	if err != nil {
		fmt.Println("Error inserting table:", err)
		return nil, err
	}

	resp.Msg = "Student exists in server 1"
	return &resp, nil
}
