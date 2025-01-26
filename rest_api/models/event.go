package models

import (
	"rest_api/data_access"
	"time"
)

// Event is a struct that represents an event.

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	query := `INSERT INTO events (name, description, location, date_time, user_id) 
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := data_access.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
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

func GetAll() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := data_access.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := []Event{}
	for rows.Next() {
		e := Event{}
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}

func Get(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	stmt, err := data_access.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(query, id)

	e := Event{}
	err = row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (event Event) Update() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, date_time = ?
		WHERE id = ?`

	stmt, err := data_access.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	if err != nil {
		return err
	}

	return nil
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"

	stmt, err := data_access.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	if err != nil {
		return err
	}

	return nil
}

func (e *Event) Register(userId int64) error {
	query := `INSERT INTO event_registrations (event_id, user_id) VALUES (?, ?)`
	stmt, err := data_access.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	if err != nil {
		return err
	}

	return nil
}

func (e *Event) CancelRegister(userId int64) error {
	query := `DELETE FROM event_registrations WHERE event_id = ? AND user_id = ?`
	stmt, err := data_access.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	if err != nil {
		return err
	}

	return nil
}
