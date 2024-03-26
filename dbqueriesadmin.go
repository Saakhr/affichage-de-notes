package main

import (
	"database/sql"

	"github.com/Saakhr/go-affichage-de-notes/models"
)

func getGroups(db *sql.DB) []models.Specialite {
	query := "SELECT GroupID,GroupName FROM `Group`"

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Create a slice to hold the results
	var specialites []models.Specialite

	// Process the results and append them to the slice
	for rows.Next() {
		var specialite models.Specialite

		err := rows.Scan(&specialite.ID, &specialite.Name)
		if err != nil {
			panic(err.Error())
		}

		specialites = append(specialites, specialite)
	}
	return specialites
}
func getSpecialities(db *sql.DB) []models.Specialite {
	query := "SELECT * FROM Specialite"

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Create a slice to hold the results
	var specialites []models.Specialite

	// Process the results and append them to the slice
	for rows.Next() {
		var specialite models.Specialite

		err := rows.Scan(&specialite.ID, &specialite.Name)
		if err != nil {
			panic(err.Error())
		}

		specialites = append(specialites, specialite)
	}
	return specialites
}
func getASpecialitie(db *sql.DB, name string) models.Specialite {
	query := "SELECT * FROM Specialite WHERE SpecialiteName = ?"

	// Execute the query
	rows, err := db.Query(query, name)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Create a slice to hold the results

	// Process the results and append them to the slice
	var specialite models.Specialite
	for rows.Next() {

		err := rows.Scan(&specialite.ID, &specialite.Name)
		if err != nil {
			panic(err.Error())
		}

	}
	return specialite
}

func getMods(db *sql.DB) []models.Modules {
	query := "SELECT * FROM Course"

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Create a slice to hold the results
	var specialites []models.Modules

	// Process the results and append them to the slice
	for rows.Next() {
		var specialite models.Modules

		err := rows.Scan(&specialite.ID, &specialite.Name)
		if err != nil {
			panic(err.Error())
		}

		specialites = append(specialites, specialite)
	}
	return specialites
}
func getAMod(db *sql.DB, input int) models.Modules {
	query := "SELECT * FROM Course WHERE CourseID=?"

	// Execute the query
	rows, err := db.Query(query, input)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Create a slice to hold the results
	var specialites models.Modules

	// Process the results and append them to the slice
	for rows.Next() {

		err := rows.Scan(&specialites.ID, &specialites.Name)
		if err != nil {
			panic(err.Error())
		}

	}
	return specialites
}
func updateMod(db *sql.DB, input string, course int) {
	query := "UPDATE Course SET CourseName = ? WHERE CourseID = ?"

	_, err := db.Exec(query, input, course)
	if err != nil {
		panic(err.Error())
	}
}
func deleteMod(db *sql.DB, userID int) {
	query := "DELETE FROM Course WHERE CourseID = ?"

	_, err := db.Exec(query, userID)
	if err != nil {
		panic(err.Error())
	}
}
func updateSession(db *sql.DB, userID int, session string) error {
	stmt, err := db.Prepare("UPDATE Users SET Session = ? WHERE UserID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(session, userID)
	if err != nil {
		return err
	}

	return nil
}

func addSpec(db *sql.DB, name string) {
	query := "INSERT INTO Specialite (SpecialiteName) VALUES (?)"

	// Execute the SQL query
	_, err := db.Exec(query, name)
	if err != nil {
		panic(err.Error())
	}
}
func deleteSpec(db *sql.DB, Specid int) {
	query := "DELETE FROM Specialite WHERE SpecialiteID = ?;"
	query2 := "DELETE FROM CourseSpecialite WHERE SpecialiteID = ?;"

	_, err := db.Exec(query2, Specid)
	if err != nil {
		_, err = db.Exec(query, Specid)
		if err != nil {
			panic(err.Error())
		}
	}
	_, err = db.Exec(query, Specid)
	if err != nil {
		panic(err.Error())
	}
}

func updateSpec(db *sql.DB, input string, courseID int) {
	query := "UPDATE Specialite SET SpecialiteName = ? WHERE SpecialiteID = ?"

	_, err := db.Exec(query, input, courseID)
	if err != nil {
		panic(err.Error())
	}
}

func deleteSpecCourses(db *sql.DB, Specid int) {
	query := "DELETE FROM CourseSpecialite WHERE SpecialiteID = ?;"

	_, err := db.Exec(query, Specid)
	if err != nil {
		panic(err.Error())
	}
}

func addSpecCourse(db *sql.DB, input [3]int) {
	query := "INSERT INTO CourseSpecialite (CourseID,SpecialiteID,Coefficient) VALUES (?,?,?)"

	// Execute the SQL query
	_, err := db.Exec(query, input[0], input[1], input[2])
	if err != nil {
		panic(err.Error())
	}
}

func addMod(db *sql.DB, input string) {
	query := "INSERT INTO Course (CourseName) VALUES (?)"

	_, err := db.Exec(query, input)
	if err != nil {
		panic(err.Error())
	}
}
func addGrp(db *sql.DB, input string, id int) {
	query := "INSERT INTO `Group` (GroupName,SpecialiteID) VALUES (?,?)"

	_, err := db.Exec(query, input, id)
	if err != nil {
		panic(err.Error())
	}
}
func updateGrp(db *sql.DB, input string, groupid, specid int) {
	query := "UPDATE `Group` SET GroupName = ?, SpecialiteID = ? WHERE GroupID = ?"

	_, err := db.Exec(query, input, specid, groupid)
	if err != nil {
		panic(err.Error())
	}
}
func deleteGrp(db *sql.DB, userID int) {
	query := "DELETE FROM `Group` WHERE GroupID = ?"

	_, err := db.Exec(query, userID)
	if err != nil {
		panic(err.Error())
	}
}
func addStudant(db *sql.DB, input [2]string, id int) {
	query := "INSERT INTO Users (Username, Password, Role, GroupID) VALUES (?, ?, ?,?)"

	_, err := db.Exec(query, input[0], input[1], "student", id)
	if err != nil {
		panic(err.Error())
	}
}
func updateStudant(db *sql.DB, input [2]string, groupid, userid int) {
	query := "UPDATE Users SET Username = ?, Password = ?, GroupID = ? WHERE UserID = ?"

	_, err := db.Exec(query, input[0], input[1], groupid, userid)
	if err != nil {
		panic(err.Error())
	}
}
func deleteStudant(db *sql.DB, userID int) {
	query := "DELETE FROM Users WHERE UserID = ?"

	_, err := db.Exec(query, userID)
	if err != nil {
		panic(err.Error())
	}
}
func addProf(db *sql.DB, input [2]string) {
	query := "INSERT INTO Users (Username, Password, Role, GroupID) VALUES (?, ?, ?,NULL)"

	_, err := db.Exec(query, input[0], input[1], "prof")
	if err != nil {
		panic(err.Error())
	}
}

func updateProfUsername(db *sql.DB, userID int, newUsername string) {
	query := "UPDATE Users SET Username = ? WHERE UserID = ? AND Role = 'prof'"

	_, err := db.Exec(query, newUsername, userID)
	if err != nil {
		panic(err.Error())
	}
}
func deleteProfModules(db *sql.DB, Specid int) {
	query := "DELETE FROM TeacherModule WHERE TeacherID = ?;"
	_, err := db.Exec(query, Specid)
	if err != nil {
		panic(err.Error())
	}
}
func deleteProf(db *sql.DB, Specid int) {
	query := "DELETE FROM Users WHERE UserID = ?;"
	query2 := "DELETE FROM TeacherModule WHERE TeacherID = ?;"

	_, err := db.Exec(query2, Specid)
	if err != nil {
		_, err = db.Exec(query, Specid)
		if err != nil {
			panic(err.Error())
		}
	}
	_, err = db.Exec(query, Specid)
	if err != nil {
		panic(err.Error())
	}
}
func addTeacherModule(db *sql.DB, teacherID int, courseID int, groupID int) {
	query := "INSERT INTO TeacherModule (TeacherID, CourseID, GroupID) VALUES (?, ?, ?)"

	_, err := db.Exec(query, teacherID, courseID, groupID)
	if err != nil {
		panic(err.Error())
	}
}
func getAllGroupsAndCourses(db *sql.DB) []models.Group {
	query := `SELECT 
    g.GroupID,
    g.GroupName,
    c.CourseID,
    c.CourseName
FROM 
    ` + "`Group`" + `g
JOIN 
    CourseSpecialite cs ON g.SpecialiteID = cs.SpecialiteID
JOIN 
    Course c ON cs.CourseID = c.CourseID;
  `
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Map to store Group data
	groupsMap := make(map[int]models.Group)

	for rows.Next() {
		var groupID int
		var groupName string
		var courseID int
		var courseName string

		err := rows.Scan(&groupID, &groupName, &courseID, &courseName)
		if err != nil {
			panic(err.Error())
		}

		// Check if Group exists in map, if not create a new Group
		group, ok := groupsMap[groupID]
		if !ok {
			group = models.Group{Name: groupName, ID: groupID}
		}

		// Append the course to the Group's Modules
		group.Modules = append(group.Modules, models.Modules{Name: courseName, ID: courseID})

		// Update the Group in the map
		groupsMap[groupID] = group
	}

	// Convert map to slice and return
	groups := make([]models.Group, 0, len(groupsMap))
	for _, group := range groupsMap {
		groups = append(groups, group)
	}
	return groups
}
func getAllProfs(db *sql.DB) []models.Profe {
	query := `SELECT 
    UserID,
    Username
FROM 
    Users
WHERE 
    Role = 'prof';
  `
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var profs []models.Profe
	for rows.Next() {
		var prof models.Profe
		err := rows.Scan(&prof.ID, &prof.Name)
		if err != nil {
			panic(err.Error())
		}
		profs = append(profs, prof)
	}
	return profs

}
func getAllProfsWithModulesAndGroups(db *sql.DB) []models.Prof {
	query := `SELECT 
    T.UserID,
    T.Username,
    G.GroupID,
    G.GroupName,
    C.CourseID,
    C.CourseName
FROM 
    TeacherModule TM
JOIN 
    Users T ON TM.TeacherID = T.UserID
JOIN 
    ` + "`Group`" + ` G ON TM.GroupID = G.GroupID
JOIN 
    Course C ON TM.CourseID = C.CourseID
WHERE 
    T.Role = 'prof';
  `
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Map to store Prof data
	profsMap := make(map[string]models.Prof)

	for rows.Next() {
		var userID string
		var username string
		var groupID int
		var groupName string
		var courseID int
		var courseName string

		err := rows.Scan(&userID, &username, &groupID, &groupName, &courseID, &courseName)
		if err != nil {
			panic(err.Error())
		}

		// Check if Prof exists in map, if not create a new Prof
		prof, ok := profsMap[userID]
		if !ok {
			prof = models.Prof{ID: userID, Name: username}
		}

		// Check if Group exists in Prof, if not create a new Group
		var groupExists bool
		var groupIndex int
		for i, g := range prof.Group {
			if g.ID == groupID {
				groupExists = true
				groupIndex = i
				break
			}
		}

		if !groupExists {
			prof.Group = append(prof.Group, models.Group{Name: groupName, ID: groupID})
			groupIndex = len(prof.Group) - 1
		}

		// Append the course to the appropriate Group
		prof.Group[groupIndex].Modules = append(prof.Group[groupIndex].Modules, models.Modules{Name: courseName, ID: courseID})

		// Update the Prof in the map
		profsMap[userID] = prof
	}

	// Convert map to slice and return
	profs := make([]models.Prof, 0, len(profsMap))
	for _, prof := range profsMap {
		profs = append(profs, prof)
	}
	return profs

}
func getAllStudents(db *sql.DB) []models.User {
	query := "SELECT UserID, Username, Password, Role, SpecialiteID FROM Users WHERE Role = ?"
	rows, err := db.Query(query, "student")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Iterate over the result set and store each row in a slice of User structs
	var students []models.User
	for rows.Next() {
		var student models.User
		err := rows.Scan(&student.UserID, &student.Username, &student.Password, &student.Role, &student.GroupID)
		if err != nil {
			panic(err.Error())
		}
		students = append(students, student)
	}
	return students
}
func getAStudentwithspec(db *sql.DB, id int) models.UserSpecialite {
	query := `SELECT 
    u.UserID,
    u.Username,
    u.Password,
    g.GroupName,
    g.GroupID,
    s.SpecialiteName,
    s.SpecialiteID
FROM 
    Users u
JOIN
    ` + "`Group`" + ` g ON u.GroupID = g.GroupID
JOIN 
    Specialite s ON g.SpecialiteID = s.SpecialiteID
WHERE 
    u.UserID = ?;
`

	rows, err := db.Query(query, id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var student models.UserSpecialite
	for rows.Next() {
		// Iterate over the result set and store each row in a slice of User structs
		err := rows.Scan(&student.UserID, &student.Username, &student.Password, &student.Group.Name, &student.Group.ID, &student.Specialite.Name, &student.Specialite.ID)
		if err != nil {
			panic(err.Error())
		}
	}
	return student
}
func getAllStudentswithspec(db *sql.DB) []models.UserSpecialite {
	query := `SELECT 
    u.UserID,
    u.Username,
    u.Password,
    g.GroupName,
    g.GroupID,
    s.SpecialiteName,
    s.SpecialiteID
FROM 
    Users u
JOIN
    ` + "`Group`" + ` g ON u.GroupID = g.GroupID
JOIN 
    Specialite s ON g.SpecialiteID = s.SpecialiteID
WHERE 
    u.Role = 'student';
`

	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Iterate over the result set and store each row in a slice of User structs
	var students []models.UserSpecialite
	for rows.Next() {
		var student models.UserSpecialite
		err := rows.Scan(&student.UserID, &student.Username, &student.Password, &student.Group.Name, &student.Group.ID, &student.Specialite.Name, &student.Specialite.ID)
		if err != nil {
			panic(err.Error())
		}
		students = append(students, student)
	}
	return students
}
func getAllSpecsWithModules(db *sql.DB) []models.SpecialiteWithModules {
	query := `
    SELECT 
        Specialite.SpecialiteID,
        Specialite.SpecialiteName,
        Course.CourseID,
        Course.CourseName,
        CourseSpecialite.Coefficient
    FROM 
        Specialite
    JOIN 
        CourseSpecialite ON Specialite.SpecialiteID = CourseSpecialite.SpecialiteID
    JOIN 
        Course ON CourseSpecialite.CourseID = Course.CourseID;
`
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Iterate over the result set and store each row in a slice of User structs
	specialites := make(map[int]*models.SpecialiteWithModules)

	for rows.Next() {
		var (
			specialiteID   int
			specialiteName string
			courseID       int
			courseName     string
			Coefficient    int
		)
		if err := rows.Scan(&specialiteID, &specialiteName, &courseID, &courseName, &Coefficient); err != nil {
			panic(err.Error())
		}

		// Check if the SpecialiteWithModules instance exists for the current specialty ID
		specialite, ok := specialites[specialiteID]
		if !ok {
			// If it doesn't exist, create a new instance and store it in the map
			specialite = &models.SpecialiteWithModules{
				ID:      specialiteID,
				Name:    specialiteName,
				Modules: make([]models.Modulec, 0), // Initialize an empty slice for modules
			}
			specialites[specialiteID] = specialite
		}

		// Append the current course to the modules slice of the corresponding specialty
		specialite.Modules = append(specialite.Modules, models.Modulec{
			ID:          courseID,
			Name:        courseName,
			Coefficient: Coefficient,
		})
	}

	if err := rows.Err(); err != nil {
		panic(err.Error())
	}

	// Convert the map values to a slice of SpecialiteWithModules instances
	var result []models.SpecialiteWithModules
	for _, specialite := range specialites {
		result = append(result, *specialite)
	}
	return result
}
func getAGroupwithspec(db *sql.DB, id int) models.Groupwithspecialite {
	query := `
    SELECT g.GroupID, g.GroupName,g.SpecialiteID,s.SpecialiteName
    FROM ` + "`Group`" + ` g
    JOIN Specialite s ON g.SpecialiteID = s.SpecialiteID 
    WHERE g.GroupID=?;
`

	rows, err := db.Query(query, id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var student models.Groupwithspecialite
	for rows.Next() {
		err := rows.Scan(&student.Id, &student.Name, &student.Specialite.ID, &student.Specialite.Name)
		if err != nil {
			panic(err.Error())
		}
	}
	return student
}
func getAllGroupswithspecNAME(db *sql.DB) []models.Groupwithspecialite {
	query := `
    SELECT g.GroupID, g.GroupName,g.SpecialiteID,s.SpecialiteName
    FROM ` + "`Group`" + ` g
    JOIN Specialite s ON g.SpecialiteID = s.SpecialiteID;
`
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Iterate over the result set and store each row in a slice of User structs
	var students []models.Groupwithspecialite
	for rows.Next() {
		var student models.Groupwithspecialite
		err := rows.Scan(&student.Id, &student.Name, &student.Specialite.ID, &student.Specialite.Name)
		if err != nil {
			panic(err.Error())
		}
		students = append(students, student)
	}
	return students

}

func getAllProf(db *sql.DB) []models.User {
	query := "SELECT UserID, Username, Password, Role FROM Users WHERE Role = ?"
	rows, err := db.Query(query, "prof")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Iterate over the result set and store each row in a slice of User structs
	var students []models.User
	for rows.Next() {
		var student models.User
		err := rows.Scan(&student.UserID, &student.Username, &student.Password, &student.Role)
		if err != nil {
			panic(err.Error())
		}
		students = append(students, student)
	}
	return students
}
