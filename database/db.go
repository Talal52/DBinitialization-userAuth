// database/db.go
package database

import (
	"context"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func (db *DB) GetUserByUsername(username string) (User, error) {
	var user User
	err := db.QueryRow("SELECT \"Id\", \"username\", \"password\", \"email\" FROM \"user\" WHERE \"username\" = $1", username).Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	if err != nil {
		fmt.Println("error", err)
		return User{}, err
	}
	return user, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWTToken(userID int, secret string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24) // 24 hours from now

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

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