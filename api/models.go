// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package api

type Blog struct {
	Blogid   int32
	Title    string
	Blogtext string
	Username string
}

type User struct {
	Username   string
	Password   string
	Bio        string
	Color      string
	Profession string
}
