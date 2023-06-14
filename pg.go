package sharedtimer

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

type PostgresClient struct {
	db *sql.DB
}

func NewPostgresClient(host, port, user, password, dbName string) (PostgresClient, error) {
	var err error

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return PostgresClient{}, err
	}

	err = db.Ping()
	if err != nil {
		return PostgresClient{}, err
	}

	fmt.Printf("connected to postgres database %s at %s:%s\n", dbName, host, port)

	return PostgresClient{
		db: db,
	}, nil
}

type TimerModel struct {
	Timer
	ownerID string
}

func (cl *PostgresClient) CreateUser(id string) error {
	queryStr := fmt.Sprintf("INSERT INTO Users(\"userID\") VALUES($1)")
	_, err := cl.db.Exec(queryStr, id)
	if err != nil {
		return err
	}
	return nil
}

func (cl *PostgresClient) GetUser(id string) (string, error) {
	var userID string
	queryStr := "SELECT * FROM \"Users\" WHERE userID=$1"
	err := cl.db.QueryRow(queryStr, id).Scan(&userID)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func (cl *PostgresClient) DeleteUser(id string) error {
	queryStr := "DELETE FROM \"Users\" WHERE userID=$1"
	_, err := cl.db.Exec(queryStr, id)
	if err != nil {
		return err
	}
	return nil
}

func (cl *PostgresClient) CreateTimer(timer Timer, ownerID string) error {
	queryStr := "INSERT INTO \"Timers\"(\"timerID\" \"name\" \"initialSeconds\" \"remainingSeconds\" \"isRunning\" \"ownerID\") VALUES($1 $2 $3 $4 $5 $6)"
	_, err := cl.db.Exec(queryStr, timer.Id, timer.Name, timer.InitialSeconds, timer.RemainingSeconds, timer.IsRunning, ownerID)
	if err != nil {
		return err
	}
	return nil
}

func (cl *PostgresClient) GetTimer(id uuid.UUID) (Timer, error) {
	var t Timer
	queryStr := "SELECT * FROM \"Timers\" WHERE timerID=$1"
	err := cl.db.QueryRow(queryStr, id).Scan(&t.Id, &t.Name, &t.InitialSeconds, &t.RemainingSeconds, &t.IsRunning)
	if err != nil {
		return Timer{}, err
	}
	return t, nil
}

// Query by:
// timers I own
// timers I can edit
// timers I can read
// timers with gt, geq, eq, leq, lt remaining time
// timers I share with another user (query by multiple user ID)
// paused/unpaused timers
func (cl *PostgresClient) QueryTimers() error {
	return nil
}

func (cl *PostgresClient) UpdateTimer() error {
	return nil
}

func (cl *PostgresClient) DeleteTimer(id string) error {
	queryStr := "DELETE FROM \"Timers\" WHERE timerID=$1"
	_, err := cl.db.Exec(queryStr, id)
	if err != nil {
		return err
	}
	return nil
}

func (cl *PostgresClient) AddEditAccess(timerID uuid.UUID, userID string) error {
	return nil
}

func (cl *PostgresClient) RemoveEditAccess(timerID uuid.UUID, userID string) error {
	return nil
}

func (cl *PostgresClient) AddReadAccess(timerID uuid.UUID, userID string) error {
	return nil
}

func (cl *PostgresClient) RemoveReadAccess(timerID uuid.UUID, userID string) error {
	return nil
}
