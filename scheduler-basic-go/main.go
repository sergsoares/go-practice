package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
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

func (s Scheduler) CheckInterval(ctx context.Context, duration time.Duration) {
	ticker := time.NewTicker(duration)
	go func() {
		for {
			select {
			case <-ctx.Done():
			case <-ticker.C:
				log.Println("Ticket Received...")
				events := s.CheckDueEvents()
				for _, e := range events {
					s.callListeners(e)
				}
			}
		}
	}()
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

func (s Scheduler) callListeners(event Event) {
	eventFn, ok := s.listeners[event.Name]
	if ok {
		go eventFn(event.Payload)
		_, err := s.db.Exec(`DELETE FROM jobs where id = $1`, event.ID)
		if err != nil {
			log.Print("Error", err)
		}
	} else {
		log.Println("Couldn't find the log attached with ", event.Name)
	}
}

func (s Scheduler) AddListener(event string, listenFunc ListenFunc) {
	s.listeners[event] = listenFunc
}

var eventListener = Listeners{
	"log": func(payload string) { log.Println("msg: ", payload) },
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
	ctx, cancel := context.WithCancel(context.Background())
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	db := newConnection()
	s := NewScheduler(db, eventListener)

	s.CheckInterval(ctx, time.Second)
	go func() {
		for range interrupt {
			log.Print("\nInterrupt received closing...")
			cancel()
		}
	}()

	<-ctx.Done()
}

func addf(payload string) {
	// eventId := uuid.New()
	db := newConnection()
	s := NewScheduler(db, eventListener)
	s.Schedule("log", payload, time.Now())
}
