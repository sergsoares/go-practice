package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
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

func (s Scheduler) CheckDueEvents() []Event {
	log.Println("Checking due events")
	rows, err := s.db.Query(`SELECT "id", "name", "payload" FROM "public"."jobs" WHERE "runAt" < $1`, time.Now())
	if err != nil {
		log.Print("Error when quering")
	}

	events := []Event{}
	for rows.Next() {
		evt := Event{}
		rows.Scan(&evt.ID, &evt.Name, &evt.Payload)
		events = append(events, evt)
	}

	return events
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

const (
	add    string = "add"
	worker string = "worker"
)

func main() {
	var payload string
	var ty string
	flag.StringVar(&payload, "m", "", "")
	flag.StringVar(&ty, "type", "add", "")
	flag.Parse()

	switch ty {
	case add:
		addf(payload)
	case worker:
		workerf()
	default:
		fmt.Println("type not found")
	}

}

func workerf() {
	db := newConnection()
	s := NewScheduler(db, eventListener)
	events := s.CheckDueEvents()

	for _, v := range events {
		log.Println("Showing", v)
	}
}

func addf(payload string) {
	eventId := uuid.New()
	db := newConnection()
	s := NewScheduler(db, eventListener)
	s.Schedule(eventId.String(), payload, time.Now())
}
