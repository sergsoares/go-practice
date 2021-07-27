package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Event struct {
	ID      uint
	Name    string
	Payload string
}

type Listeners map[string]ListenFunc

type ListenFunc func(string)

type Scheduler struct {
	db        *sql.DB
	listeners Listeners
}

func NewScheduler(db *sql.DB, listeners Listeners) Scheduler {
	return Scheduler{
		db:        db,
		listeners: listeners,
	}
}

func (s Scheduler) Schedule(event, payload string, runAt time.Time) {
	log.Println("Scheduling event", event, "to run at", runAt)
	_, err := s.db.Exec(`INSERT INTO "public"."jobs" ("name", "payload", "runAt") VALUES ($1, $2, $3)`, event, payload, runAt)

	if err != nil {
		log.Println("Scheduler with errors", err)
	}
}

func (s Scheduler) AddListener(event string, listenFunc ListenFunc) {
	s.listeners[event] = listenFunc
}

var eventListener = Listeners{
	"log": func(payload string) { log.Println("executed job") },
}

func newConnection() *sql.DB {
	url := "postgresql://postgres:postgres@localhost:5432/db?sslmode=disable"

	db, err := sql.Open("postgres", url)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	db := newConnection()
	s := NewScheduler(db, eventListener)
	s.Schedule("test", "bla", time.Now())
}
