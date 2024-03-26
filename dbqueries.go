package main

import (
	"database/sql"

	"github.com/Saakhr/go-affichage-de-notes/models"
)

func addGrade(db *sql.DB, courseID int, userID int, gradeValue [4]float64) {
	query := "INSERT INTO Grade (CourseID, UserID, Exam, TD, TP, Projet) VALUES (?, ?, ?, ?, ?, ?)"

	_, err := db.Exec(query, courseID, userID, gradeValue[0], gradeValue[1], gradeValue[2], gradeValue[3])
	if err != nil {
		panic(err.Error())
	}
}
func getStudentsByTeacherID(db *sql.DB, teacherID int) []models.StudentGroupCourse {
	query := `
        SELECT 
            U.UserID,
            U.Username AS StudentName,
            G.GroupName,
            C.CourseID,
            C.CourseName
        FROM 
            Users U
        JOIN 
            ` + "`Group`" + ` G ON U.GroupID = G.GroupID
        JOIN 
            TeacherModule TM ON G.GroupID = TM.GroupID
        JOIN 
            Course C ON TM.CourseID = C.CourseID
        WHERE 
            TM.TeacherID = ? AND U.Role = 'student'
            AND NOT EXISTS (
                SELECT 1
                FROM Grade Gd
                WHERE Gd.UserID = U.UserID AND Gd.CourseID = C.CourseID
            )
    `
	rows, err := db.Query(query, teacherID)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var studentGroupCourses []models.StudentGroupCourse
	for rows.Next() {
		var sgc models.StudentGroupCourse
		err := rows.Scan(&sgc.UserID, &sgc.StudentName, &sgc.GroupName, &sgc.CourseID, &sgc.CourseName)
		if err != nil {
			panic(err.Error())
		}
		studentGroupCourses = append(studentGroupCourses, sgc)
	}
	return studentGroupCourses
}

func getGradesByTeacherID(db *sql.DB, teacherID int) []models.GradeInfo {
	query := `
        SELECT 
            U.Username AS StudentName,
            Gp.GroupName,
            Sp.SpecialiteName,
            C.CourseName,
            G.Exam,
            G.TD,
            G.TP,
            G.Projet
        FROM 
            Grade G
        JOIN 
            Users U ON G.UserID = U.UserID
        JOIN 
            Course C ON G.CourseID = C.CourseID
        JOIN 
            TeacherModule TM ON G.CourseID = TM.CourseID
        JOIN 
            ` + "`Group`" + ` Gp ON TM.GroupID = Gp.GroupID
        JOIN 
            Specialite Sp ON Gp.SpecialiteID = Sp.SpecialiteID
        WHERE 
            TM.TeacherID = ?
    `

	rows, err := db.Query(query, teacherID)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var grades []models.GradeInfo
	for rows.Next() {
		var grade models.GradeInfo
		err := rows.Scan(&grade.StudentName, &grade.GroupName, &grade.SpecialiteName, &grade.CourseName, &grade.Exam, &grade.TD, &grade.TP, &grade.Projet)
		if err != nil {
			panic(err.Error())
		}
		grades = append(grades, grade)
	}
	return grades
}
