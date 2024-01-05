// database/db.go
package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type DB struct {
	*sql.DB
}

func InitializeDB() (*DB, error) {

	os.Unsetenv("PGLOCALEDIR")

	db, err := sql.Open("postgres", "user=postgres password=12345 dbname=MyDatabase sslmode=disable")
	if err != nil {
		fmt.Println("error 1", err)
		return nil, err
	}

	return &DB{db}, nil
}

// func (db *DB) CreateUser(user User) (int, error) {
// 	var userID int
// 	err := db.QueryRow("INSERT INTO \"user\" (\"Id\", \"username\", \"password\", \"email\") VALUES ($1, $2, $3, $4) RETURNING \"Id\"", user.Id, user.Username, user.Password, user.Email).Scan(&userID)
// 	if err != nil {
// 		fmt.Println("error 2", err)
// 		return 0, err
// 	}
// 	return userID, nil
// }

func (db *DB) CreateUser(user User) (int, error) {
	var userID int
	query := "INSERT INTO \"user\" (\"username\", \"password\", \"email\") VALUES ($1, $2, $3) RETURNING \"Id\""

	err := db.QueryRowContext(context.Background(), query, user.Username, user.Password, user.Email).Scan(&userID)

	if err != nil {
		fmt.Println("Error creating user:", err)
		return 0, err
	}

	return userID, nil
}

func (db *DB) GetUserByID(userID int) (User, error) {
	var user User
	err := db.QueryRow("SELECT \"Id\", \"username\", \"email\" FROM \"user\" WHERE \"Id\" = $1", userID).Scan(&user.Id, &user.Username, &user.Email)
	if err != nil {
		fmt.Println("error", err)
		return User{}, err
	}
	return user, nil
}
