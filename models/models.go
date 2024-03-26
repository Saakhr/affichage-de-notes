package models

type Specialite struct {
	Name string
	ID   int
}
type SpecialiteWithModules struct {
	Name    string
	ID      int
	Modules []Modulec
}
type Modulec struct {
	Name        string
	ID          int
	Coefficient int
}
type Modules struct {
	Name string
	ID   int
}
type User struct {
	UserID   int
	Username string
	Password string
	Role     string
	GroupID  int
}
type UserSpecialite struct {
	UserID     int
	Username   string
	Password   string
	Role       string
	Specialite Specialite
	Group      Specialite
}
type Groupwithspecialite struct {
	Name       string
	Id         string
	Specialite Specialite
}
type Group struct {
	Name    string
	ID      int
	Modules []Modules
}
type Profe struct {
	ID   string
	Name string
}
type Prof struct {
	ID    string
	Name  string
	Group []Group
}
type ProfDB struct {
	ID     string
	Name   string
	Group  Specialite
	Module Modules
}
