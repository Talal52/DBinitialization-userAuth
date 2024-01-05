// database/db.go
package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// DB represents the database connection.
type DB struct {
	*sql.DB
}

// InitializeDB initializes the database connection.
func InitializeDB() (*DB, error) {
	// Replace these values with your actual database connection details
	db, err := sql.Open("postgres", "user=your_username dbname=your_database sslmode=disable")
	if err != nil {
		fmt.Println("error 1", err)
		return nil, err
	}

	return &DB{db}, nil
}

// CreateUser creates a new user in the database.
func (db *DB) CreateUser(user User) (int, error) {
	var userID int
	err := db.QueryRow("INSERT INTO users(username, email) VALUES($1, $2) RETURNING id", user.Username, user.Email).Scan(&userID)
	if err != nil {
		fmt.Println("error 2", err)
		return 0, err
	}
	return userID, nil
}

// GetUserByID retrieves a user from the database by ID.
func (db *DB) GetUserByID(userID int) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, email FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		fmt.Println("error", err)
		return User{}, err
	}
	return user, nil
}
