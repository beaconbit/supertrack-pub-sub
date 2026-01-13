package main

import (
    "context"
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "time"

    "github.com/nats-io/nats.go"
    _ "github.com/lib/pq"
)

// Your message struct
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

type UniqueMessage struct {
  UUID string
  Timestamp time.Time
  Label     string  
  Value     int    
  Factor    float32
  Reading   float32 
  Measurement float32
  Metric    int      
  DFI       int      
}

type Load struct {
  UUID 		string
  Entered 	time.Time
  Exited 	time.Time
  Delta 	time.Duration
}

func pocketStatusUpdater(
  incoming <-chan UniqueMessage, 
  outgoing chan<- UniqueMessage, 
  tableName string) {
    db, err := sql.Open("postgres", "postgres://admin:password123@postgres:5432/eventlog?sslmode=disable")
    if err != nil {
        log.Fatal("postgres connect:", err)
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        log.Fatal("postgres ping:", err)
    }
    fmt.Println("Connected to Postgres")

    // listen for messages
    for {
	select {
	case e := <-events:
	// update struct 
	// write to database
	// emit message to next channel

	}
    }



}

func main() {
    // --- Connect to Postgres ---
    db, err := sql.Open("postgres", "postgres://admin:password123@postgres:5432/eventlog?sslmode=disable")
    if err != nil {
        log.Fatal("postgres connect:", err)
    }
    defer db.Close()

    // Make sure DB is reachable
    if err := db.Ping(); err != nil {
        log.Fatal("postgres ping:", err)
    }
    fmt.Println("Connected to Postgres")


    // --- Connect to NATS ---
    nc, err := nats.Connect("nats://nats:4222")
    if err != nil {
        log.Fatal("nats connect:", err)
    }
    defer nc.Close()

    fmt.Println("Connected to NATS, waiting for messages...")

    // --- Subscribe ---
    _, err = nc.Subscribe("eventstream.>", func(m *nats.Msg) {
        var msg Message

        // Decode JSON
        if err := json.Unmarshal(m.Data, &msg); err != nil {
            log.Println("decode error:", err)
            return
        }
	fmt.Println("message read from nats")
	fmt.Printf("%v", msg)

	// send event over go chan to loop 1 subscriber

	// insert into postgres
	_, err := db.ExecContext(
	    context.Background(),
	    `INSERT INTO general (
		timestamp,
		source_mac,
		source_ip,
		value,
		factor,
		reading,
		measurement,
		metric,
		data_field_index
	    ) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9
	    )`,
	    msg.Timestamp.Unix(),
	    "00:00:00:00",
	    "0.0.0.0",
	    msg.Value,
	    msg.Factor,
	    msg.Reading,
	    msg.Measurement,
	    msg.Metric,
	    msg.DFI,
	)
	if err != nil {
	    log.Println("db insert error:", err)
	    return
	}

        fmt.Printf("Inserted event: %s (%s)\n", msg.Label, msg.Timestamp.Format(time.RFC3339))
    })
    if err != nil {
        log.Fatal("subscribe:", err)
    }

    go pocketStatusUpdater(cbw1pocket1Ch, cbw1pocket2Ch, "cbw1pocket1")
    go pocketStatusUpdater(cbw1pocket2Ch, cbw1pocket3Ch, "cbw1pocket2")
    go pocketStatusUpdater(cbw1pocket3Ch, cbw1pocket4Ch, "cbw1pocket3")
    go pocketStatusUpdater(cbw1pocket4Ch, cbw1pocket5Ch, "cbw1pocket4")
    go pocketStatusUpdater(cbw1pocket5Ch, cbw1pocket6Ch, "cbw1pocket5")
    go pocketStatusUpdater(cbw1pocket6Ch, cbw1pocket7Ch, "cbw1pocket6")
    go pocketStatusUpdater(cbw1pocket7Ch, cbw1pocket8Ch, "cbw1pocket7")
    go pocketStatusUpdater(cbw1pocket8Ch, cbw1pocket9Ch, "cbw1pocket8")
    go pocketStatusUpdater(cbw1pocket9Ch, cbw1pocket10Ch, "cbw1pocket9")
    go pocketStatusUpdater(cbw1pocket10Ch, cbw1pocket11Ch, "cbw1pocket10")
    go pocketStatusUpdater(cbw1pocket11Ch, cbw1pocket12Ch, "cbw1pocket11")
    go pocketStatusUpdater(cbw1pocket12Ch, cbw1pocket13Ch, "cbw1pocket12")
    go pocketStatusUpdater(cbw1pocket13Ch, cbw1pocket14Ch, "cbw1pocket13")
    go pocketStatusUpdater(cbw1pocket14Ch, cbw1pocket15Ch, "cbw1pocket14")
    go pocketStatusUpdater(cbw1pocket15Ch, cbw1pocket16Ch, "cbw1pocket15")
    go pocketStatusUpdater(cbw1pocket16Ch, cbw1pocket17Ch, "cbw1pocket16")
    go pocketStatusUpdater(cbw1pocket17Ch, cbw1pocket18Ch, "cbw1pocket17")
    go pocketStatusUpdater(cbw1pocket18Ch, cbw1pocket19Ch, "cbw1pocket18")
    go pocketStatusUpdater(cbw1pocket19Ch, cbw1pocket20Ch, "cbw1pocket19")
    go pocketStatusUpdater(cbw1pocket20Ch, cbw1pocket21Ch, "cbw1pocket20")
    go pocketStatusUpdater(cbw1pocket21Ch, cbw1pocket22Ch, "cbw1pocket21")
    go pocketStatusUpdater(cbw1pocket22Ch, cbw1pocket23Ch, "cbw1pocket22")
    go pocketStatusUpdater(cbw1pocket23Ch, cbw1pocket24Ch, "cbw1pocket23")
    go pocketStatusUpdater(cbw1pocket24Ch, cbw1pocket25Ch, "cbw1pocket24")
    go pocketStatusUpdater(cbw1pocket25Ch, cbw1pocket26Ch, "cbw1pocket25")
    go pocketStatusUpdater(cbw1pocket26Ch, cbw1pocket27Ch, "cbw1pocket26")
    go pocketStatusUpdater(cbw1pocket27Ch, cbw1pocket28Ch, "cbw1pocket27")
    go pocketStatusUpdater(cbw1pocket28Ch, cbw1pocket29Ch, "cbw1pocket28")
    go pocketStatusUpdater(cbw1pocket29Ch, cbw1pocket30Ch, "cbw1pocket29")
    go pocketStatusUpdater(cbw1pocket30Ch, cbw1pocket31Ch, "cbw1pocket30")
    go pocketStatusUpdater(cbw1pocket31Ch, cbw1pocket32Ch, "cbw1pocket31")
    go pocketStatusUpdater(cbw1pocket32Ch, dead, "cbw1pocket32")

    go pocketStatusUpdater(cbw2pocket1Ch, cbw2pocket2Ch, "cbw2pocket1")
    go pocketStatusUpdater(cbw2pocket2Ch, cbw2pocket3Ch, "cbw2pocket2")
    go pocketStatusUpdater(cbw2pocket3Ch, cbw2pocket4Ch, "cbw2pocket3")
    go pocketStatusUpdater(cbw2pocket4Ch, cbw2pocket5Ch, "cbw2pocket4")
    go pocketStatusUpdater(cbw2pocket5Ch, cbw2pocket6Ch, "cbw2pocket5")
    go pocketStatusUpdater(cbw2pocket6Ch, cbw2pocket7Ch, "cbw2pocket6")
    go pocketStatusUpdater(cbw2pocket7Ch, cbw2pocket8Ch, "cbw2pocket7")
    go pocketStatusUpdater(cbw2pocket8Ch, cbw2pocket9Ch, "cbw2pocket8")
    go pocketStatusUpdater(cbw2pocket9Ch, cbw2pocket10Ch, "cbw2pocket9")
    go pocketStatusUpdater(cbw2pocket10Ch, cbw2pocket11Ch, "cbw2pocket10")
    go pocketStatusUpdater(cbw2pocket11Ch, cbw2pocket12Ch, "cbw2pocket11")
    go pocketStatusUpdater(cbw2pocket12Ch, cbw2pocket13Ch, "cbw2pocket12")
    go pocketStatusUpdater(cbw2pocket13Ch, cbw2pocket14Ch, "cbw2pocket13")
    go pocketStatusUpdater(cbw2pocket14Ch, cbw2pocket15Ch, "cbw2pocket14")
    go pocketStatusUpdater(cbw2pocket15Ch, cbw2pocket16Ch, "cbw2pocket15")
    go pocketStatusUpdater(cbw2pocket16Ch, cbw2pocket17Ch, "cbw2pocket16")
    go pocketStatusUpdater(cbw2pocket17Ch, cbw2pocket18Ch, "cbw2pocket17")
    go pocketStatusUpdater(cbw2pocket18Ch, cbw2pocket19Ch, "cbw2pocket18")
    go pocketStatusUpdater(cbw2pocket19Ch, cbw2pocket20Ch, "cbw2pocket19")
    go pocketStatusUpdater(cbw2pocket20Ch, cbw2pocket21Ch, "cbw2pocket20")
    go pocketStatusUpdater(cbw2pocket21Ch, cbw2pocket22Ch, "cbw2pocket21")
    go pocketStatusUpdater(cbw2pocket22Ch, cbw2pocket23Ch, "cbw2pocket22")
    go pocketStatusUpdater(cbw2pocket23Ch, cbw2pocket24Ch, "cbw2pocket23")
    go pocketStatusUpdater(cbw2pocket24Ch, cbw2pocket25Ch, "cbw2pocket24")
    go pocketStatusUpdater(cbw2pocket25Ch, cbw2pocket26Ch, "cbw2pocket25")
    go pocketStatusUpdater(cbw2pocket26Ch, cbw2pocket27Ch, "cbw2pocket26")
    go pocketStatusUpdater(cbw2pocket27Ch, cbw2pocket28Ch, "cbw2pocket27")
    go pocketStatusUpdater(cbw2pocket28Ch, cbw2pocket29Ch, "cbw2pocket28")
    go pocketStatusUpdater(cbw2pocket29Ch, cbw2pocket30Ch, "cbw2pocket29")
    go pocketStatusUpdater(cbw2pocket30Ch, cbw2pocket31Ch, "cbw2pocket30")
    go pocketStatusUpdater(cbw2pocket31Ch, cbw2pocket32Ch, "cbw2pocket31")
    go pocketStatusUpdater(cbw2pocket32Ch, dead, "cbw2pocket32")

    go pocketStatusUpdater(cbw3pocket1Ch, cbw3pocket2Ch, "cbw3pocket1")
    go pocketStatusUpdater(cbw3pocket2Ch, cbw3pocket3Ch, "cbw3pocket2")
    go pocketStatusUpdater(cbw3pocket3Ch, cbw3pocket4Ch, "cbw3pocket3")
    go pocketStatusUpdater(cbw3pocket4Ch, cbw3pocket5Ch, "cbw3pocket4")
    go pocketStatusUpdater(cbw3pocket5Ch, cbw3pocket6Ch, "cbw3pocket5")
    go pocketStatusUpdater(cbw3pocket6Ch, cbw3pocket7Ch, "cbw3pocket6")
    go pocketStatusUpdater(cbw3pocket7Ch, cbw3pocket8Ch, "cbw3pocket7")
    go pocketStatusUpdater(cbw3pocket8Ch, cbw3pocket9Ch, "cbw3pocket8")
    go pocketStatusUpdater(cbw3pocket9Ch, cbw3pocket10Ch, "cbw3pocket9")
    go pocketStatusUpdater(cbw3pocket10Ch, cbw3pocket11Ch, "cbw3pocket10")
    go pocketStatusUpdater(cbw3pocket11Ch, cbw3pocket12Ch, "cbw3pocket11")
    go pocketStatusUpdater(cbw3pocket12Ch, cbw3pocket13Ch, "cbw3pocket12")
    go pocketStatusUpdater(cbw3pocket13Ch, cbw3pocket14Ch, "cbw3pocket13")
    go pocketStatusUpdater(cbw3pocket14Ch, cbw3pocket15Ch, "cbw3pocket14")
    go pocketStatusUpdater(cbw3pocket15Ch, cbw3pocket16Ch, "cbw3pocket15")
    go pocketStatusUpdater(cbw3pocket16Ch, cbw3pocket17Ch, "cbw3pocket16")
    go pocketStatusUpdater(cbw3pocket17Ch, cbw3pocket18Ch, "cbw3pocket17")
    go pocketStatusUpdater(cbw3pocket18Ch, cbw3pocket19Ch, "cbw3pocket18")
    go pocketStatusUpdater(cbw3pocket19Ch, cbw3pocket20Ch, "cbw3pocket19")
    go pocketStatusUpdater(cbw3pocket20Ch, cbw3pocket21Ch, "cbw3pocket20")
    go pocketStatusUpdater(cbw3pocket21Ch, cbw3pocket22Ch, "cbw3pocket21")
    go pocketStatusUpdater(cbw3pocket22Ch, cbw3pocket23Ch, "cbw3pocket22")
    go pocketStatusUpdater(cbw3pocket23Ch, cbw3pocket24Ch, "cbw3pocket23")
    go pocketStatusUpdater(cbw3pocket24Ch, cbw3pocket25Ch, "cbw3pocket24")
    go pocketStatusUpdater(cbw3pocket25Ch, cbw3pocket26Ch, "cbw3pocket25")
    go pocketStatusUpdater(cbw3pocket26Ch, cbw3pocket27Ch, "cbw3pocket26")
    go pocketStatusUpdater(cbw3pocket27Ch, cbw3pocket28Ch, "cbw3pocket27")
    go pocketStatusUpdater(cbw3pocket28Ch, cbw3pocket29Ch, "cbw3pocket28")
    go pocketStatusUpdater(cbw3pocket29Ch, cbw3pocket30Ch, "cbw3pocket29")
    go pocketStatusUpdater(cbw3pocket30Ch, cbw3pocket31Ch, "cbw3pocket30")
    go pocketStatusUpdater(cbw3pocket31Ch, cbw3pocket32Ch, "cbw3pocket31")
    go pocketStatusUpdater(cbw3pocket32Ch, dead, "cbw3pocket32")
    // Keep running forever
    select {}
}

