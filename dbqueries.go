package main

import (
	"database/sql"

	"github.com/Saakhr/go-affichage-de-notes/models"
)

func addGrade(db *sql.DB, teacherID, courseID, userID int, gradeValue [4]float64) {
	query := "INSERT INTO Grade (TeacherID, CourseID, UserID, Exam, TD, TP, Projet) VALUES (?,?, ?, ?, ?, ?, ?)"

	_, err := db.Exec(query, teacherID, courseID, userID, gradeValue[0], gradeValue[1], gradeValue[2], gradeValue[3])
	if err != nil {
		panic(err.Error())
	}
}
func updateGrade(db *sql.DB, GradeID int, gradeValue [4]float64) {
	query := `UPDATE Grade
SET Exam = ?, TD = ?, TP = ?, Projet = ?
WHERE GradeID = ?;
`

	_, err := db.Exec(query, gradeValue[0], gradeValue[1], gradeValue[2], gradeValue[3], GradeID)
	if err != nil {
		panic(err.Error())
	}
}
func deleteGrade(db *sql.DB, GradeID int) {
	query := "DELETE FROM Grade WHERE GradeID = ?"

	_, err := db.Exec(query, GradeID)
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
            G.GradeID,
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
            G.TeacherID = ?
    `

	rows, err := db.Query(query, teacherID)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var grades []models.GradeInfo
	for rows.Next() {
		var grade models.GradeInfo
		err := rows.Scan(&grade.GradeID, &grade.StudentName, &grade.GroupName, &grade.SpecialiteName, &grade.CourseName, &grade.Exam, &grade.TD, &grade.TP, &grade.Projet)
		if err != nil {
			panic(err.Error())
		}

		grades = append(grades, grade)
	}
	calculateMoyenne(grades)

	return grades
}

func getGradesWithSpecializationAndGroup(db *sql.DB, userID int) []models.GradeInfoStudent {
	query := `
        SELECT 
            G.GradeID,
            U.Username AS StudentName,
            T.Username AS TeacherName,
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
        JOIN 
            Users T ON TM.TeacherID = T.UserID
        WHERE 
            G.UserID = ?
    `

	rows, err := db.Query(query, userID)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var grades []models.GradeInfoStudent
	for rows.Next() {
		var grade models.GradeInfoStudent
		err := rows.Scan(
			&grade.GradeID,
			&grade.StudentName,
			&grade.TeacherName,
			&grade.GroupName,
			&grade.SpecialiteName,
			&grade.CourseName,
			&grade.Exam,
			&grade.TD,
			&grade.TP,
			&grade.Projet,
		)
		if err != nil {
			panic(err.Error())
		}
		grades = append(grades, grade)
	}
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
	calculateMoyenne2(grades)

	return grades
}

func calculateMoyenne2(g []models.GradeInfoStudent) {
	for i := 0; i < len(g); i++ {
		count := 0
		sum := 0.0

		if g[i].TD != -1 {
			count++
			sum += g[i].TD
		}

		if g[i].TP != -1 {
			count++
			sum += g[i].TP
		}

		if g[i].Projet != -1 {
			count++
			sum += g[i].Projet
		}

		if g[i].Exam != -1 {
			count += 2
			sum += g[i].Exam * 2
		}

		if count > 0 {
			g[i].Moyenne = sum / float64(count)
		} else {
			g[i].Moyenne = 0 // or any default value you prefer
		}
	}
}
func calculateMoyenne(g []models.GradeInfo) {
	for i := 0; i < len(g); i++ {
		count := 0
		sum := 0.0

		if g[i].TD != -1 {
			count++
			sum += g[i].TD
		}

		if g[i].TP != -1 {
			count++
			sum += g[i].TP
		}

		if g[i].Projet != -1 {
			count++
			sum += g[i].Projet
		}

		if g[i].Exam != -1 {
			count += 2
			sum += g[i].Exam * 2
		}

		if count > 0 {
			g[i].Moyenne = sum / float64(count)
		} else {
			g[i].Moyenne = 0 // or any default value you prefer
		}
	}
}
