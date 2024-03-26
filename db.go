package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ExecSQL(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
	}
}
func setupDB() *sql.DB {
	dataSourceName := "user:user@tcp(localhost:3306)/db1"

	// Open a connection to the database
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to the MySQL database")

	ExecSQL(db, `
        CREATE TABLE IF NOT EXISTS Specialite (
            SpecialiteID INT AUTO_INCREMENT PRIMARY KEY,
            SpecialiteName VARCHAR(100)
        );
    `)

	// Create Group table
	ExecSQL(db, `
        CREATE TABLE IF NOT EXISTS `+"`Group`"+` (
            GroupID INT AUTO_INCREMENT PRIMARY KEY,
            GroupName VARCHAR(100),
            SpecialiteID INT,
            FOREIGN KEY (SpecialiteID) REFERENCES Specialite(SpecialiteID)
        );
    `)

	// Create Users table
	ExecSQL(db, `
        CREATE TABLE IF NOT EXISTS Users (
            UserID INT AUTO_INCREMENT PRIMARY KEY,
            Username VARCHAR(50),
            Password VARCHAR(100),
            Session VARCHAR(100),
            Role VARCHAR(20),
            GroupID INT,
            FOREIGN KEY (GroupID) REFERENCES `+"`Group`"+`(GroupID)
        );
    `)

	// Create Course table
	ExecSQL(db, `
        CREATE TABLE IF NOT EXISTS Course (
            CourseID INT AUTO_INCREMENT PRIMARY KEY,
            CourseName VARCHAR(100)
        );
    `)

	// Create CourseSpecialite table
	ExecSQL(db, `
        CREATE TABLE IF NOT EXISTS CourseSpecialite (
            CourseSpecialiteID INT AUTO_INCREMENT PRIMARY KEY,
            CourseID INT,
            SpecialiteID INT,
            Coefficient INT,
            FOREIGN KEY (CourseID) REFERENCES Course(CourseID),
            FOREIGN KEY (SpecialiteID) REFERENCES Specialite(SpecialiteID),
            UNIQUE (CourseID, SpecialiteID)
        );
    `)

	// Create TeacherModule table
	ExecSQL(db, `
        CREATE TABLE IF NOT EXISTS TeacherModule (
            TeacherModuleID INT AUTO_INCREMENT PRIMARY KEY,
            TeacherID INT,
            CourseID INT,
            GroupID INT,
            FOREIGN KEY (TeacherID) REFERENCES Users(UserID),
            FOREIGN KEY (CourseID) REFERENCES Course(CourseID),
            FOREIGN KEY (GroupID) REFERENCES `+"`Group`"+`(GroupID),
            UNIQUE (TeacherID, CourseID, GroupID)
        );
    `)

	// Create Grade table
	ExecSQL(db, `
CREATE TABLE IF NOT EXISTS Grade (
    GradeID INT AUTO_INCREMENT PRIMARY KEY,
    CourseID INT,
    UserID INT,
    Exam DECIMAL(5, 2),
    TD DECIMAL(5, 2),
    TP DECIMAL(5, 2),
    Projet DECIMAL(5, 2),
    FOREIGN KEY (CourseID) REFERENCES Course(CourseID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);
    `)
	//admin
	db.Exec(`
INSERT INTO Users (UserID, Username, Password,Session, Role)
VALUES (1, 'admin', 'admin',NULL, 'admin');
    `)
	fmt.Println("Initialized Tables")
	return db
}
