package middleware

import (
	"crud-app/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Sucessfully connected!!")

	return db
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Context-Type")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable decode body %v", err)
	}

	insertID := InsertUser(user)

	res := response{
		ID:      insertID,
		Message: "User created!",
	}

	json.NewEncoder(w).Encode(res)
}

func InsertUser(user models.User) int64 {
	db := CreateConnection()

	defer db.Close()

	statement := `INSERT INTO users (name, location, age) VALUES ($1,$2,$3) RETURNING userid`

	var id int64

	err := db.QueryRow(statement, user.Name, user.Location, user.Age).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to query %v", err)
	}

	fmt.Println("Inserted record  %v", id)

	return id
}
