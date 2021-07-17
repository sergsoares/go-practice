package middleware

import (
	"crud-app/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
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

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Context-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalln("Invalid id", id)
	}

	user, err := getUser(int64(id))

	if err != nil {
		log.Fatalln("Failed query", err)
	}

	json.NewEncoder(w).Encode(user)

}

func getUser(id int64) (models.User, error) {
	db := CreateConnection()

	defer db.Close()

	statement := "SELECT * FROM users where userid=$1"

	row := db.QueryRow(statement, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

	switch err {
	case sql.ErrNoRows:
		log.Println("No rows returned")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalln("Unable to scan row", err)

	}

	return user, err
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
