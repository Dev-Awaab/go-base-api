// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
