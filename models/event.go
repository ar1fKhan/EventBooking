package models

import (
	"EventBooking/db"
	"database/sql"
	"time"
)

// Event struct
type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int       `json:"user_id"`
}

// Save inserts an event into the database
func (e *Event) Save() error {
	query := "INSERT INTO events (name, description, date_time, user_id) VALUES (?, ?, ?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Convert time.Time to string format
	dateTimeStr := e.DateTime.Format("2006-01-02 15:04:05")

	result, err := stmt.Exec(e.Name, e.Description, dateTimeStr, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = id
	return nil
}

// GetAvailableEvents retrieves all events from the database
func GetAvailableEvents() ([]Event, error) {
	query := "SELECT id, name, description, date_time, user_id FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event
		var dateTimeStr string

		err := rows.Scan(&e.ID, &e.Name, &e.Description, &dateTimeStr, &e.UserID)
		if err != nil {
			return nil, err
		}

		// Convert string to time.Time
		/*e.DateTime, err = time.Parse("2006-01-02 15:04:05.999999-07:00", dateTimeStr)
		if err != nil {
			log.Printf("Failed to parse date_time: %v", err)
			return nil, err
		}*/

		events = append(events, e)
	}
	return events, nil
}

// GetEventByID retrieves a single event by ID
func GetEventByID(id int64) (*Event, error) {
	query := "SELECT id, name, description, date_time, user_id FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var e = &Event{}
	var dateTimeStr string

	err := row.Scan(&e.ID, &e.Name, &e.Description, &dateTimeStr, &e.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	// Convert string to time.Time
	/*e.DateTime, err = time.Parse("2006-01-02 15:04:05", dateTimeStr)
	if err != nil {
		log.Printf("Failed to parse date_time: %v", err)
		return nil, err
	}*/

	return e, nil
}
func (e Event) Update() error {
	query := "UPDATE events SET name=?, description=?, date_time=?, user_id=? WHERE id=?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	dateTimeStr := e.DateTime.Format("2006-01-02 15:04:05")
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, dateTimeStr, e.UserID, e.ID)
	return err
}

func (e Event) DeleteEvent() error {
	query := "DELETE FROM events WHERE id=?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	return err
}
