package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"context"
	"database/sql"

        "github.com/nats-io/nats.go"
	_ "github.com/microsoft/go-mssqldb"
)

// Message is what gets sent back to the main goroutine.
type Message struct {
	Timestamp time.Time `json:"timestamp"`
	Label     string    `json:"label"`
	Value     int       `json:"value"`
	Factor    float32   `json:"factor"`
	Reading   float32   `json:"reading"`
	Measurement float32 `json:"measurement"`
	Metric    int       `json:"metric"`
	DFI       int       `json:"dfi"`
}

// Listener periodically polls an HTTP endpoint and
// sends a message to msgCh if some condition is met.
type Listener struct {
	msgCh     chan<- Message // send-only channel to main program
	activeInterface string
	interval time.Duration // poll interval
	errorCount int
}

type Row struct { 
    TimeDate 	time.Time
    Category 	int
    Machine 	int
}

// NewListener constructs a Listener.
// msgCh: channel to send messages on
// ipAddress: address to poll (e.g. "http://192.168.1.50/status")
// label: identifying label for this listener
func NewListener(
	msgCh chan<- Message, 
	interval time.Duration) *Listener {
	return &Listener{
		msgCh: msgCh,
		interval: interval,
		errorCount: 0,
	}
}

// Run starts the polling loop.
// Call this in a goroutine:  go listener.Run()
// All the logic for handling strange count behaviour from the iot device, and starting from 0 regardless of the current count from last session
func (l *Listener) Run() {
    mostRecent := time.Now()
    for {
        previous, err := l.QueryDatabase(mostRecent)
	if err != nil {
	    fmt.Println(err)
	}
	mostRecent = previous
	time.Sleep(l.interval)
    }
}

func (l *Listener) QueryDatabase(previous time.Time) (time.Time, error) {
    connectionString := "sqlserver://SUP:1234@10.8.4.220:1433?database=SUPDB"

    db, err := sql.Open("sqlserver", connectionString)
    if err != nil {
	log.Fatal(err)
    }
    defer db.Close()

    ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
    defer cancel()

    rows, err := db.QueryContext(ctx, `
	SELECT TimeDate, Category, Machine
	FROM Supertrack.tblReports
	WHERE TimeDate > @p1
	`,
	previous,
    )
    if err != nil {
	log.Fatal(err)
    }
    defer rows.Close()

    latest := previous // in case no rows are returned
    fmt.Println("\nlatest time:", latest)
    count := 0
    for rows.Next() {
	count++
	var r Row
	if err := rows.Scan(&r.TimeDate, &r.Category, &r.Machine); err != nil {
	  log.Fatal(err)
	}
	latest = r.TimeDate
	fmt.Println(r)
	l.msgCh <- Message{
	    Timestamp:  r.TimeDate,
	    Label:	fmt.Sprintf("%d", r.Machine),
	    Value:	r.Category,
	    Factor:	1.0, 
	    Reading:	0.0,
	    Measurement: 0,
	    Metric:	r.Machine,
	    DFI:	0,
	}
    }
    fmt.Println("returned row count:", count)
    return latest, nil
}

func main() {
	msgCh := make(chan Message, 6)
	//username := "SUP"
	//password := "1234"

	interval := 1 * time.Minute
	listener := NewListener(msgCh, interval)
	fmt.Printf("listeners created")

	go listener.Run()

	fmt.Printf("listeners started")

	// graceful shutdown on SIGINT/SIGTERM
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		nc, err := nats.Connect("nats://localhost:4222")
		if err != nil {
		    log.Fatal(err)
		}
		defer nc.Close()

		for {
			select {
			case msg := <-msgCh:
				subject := fmt.Sprintf("eventstream.%s", msg.Label)
				data, _ := json.Marshal(msg)

				if err := nc.Publish(subject, data); err != nil {
				    log.Fatal(err)
				}
			}
		}
	}()

	<-sigs
	fmt.Println("shutting down...")
	fmt.Println("done")
}

