package database

import (
	"context"
	"database/sql"
	"time"
)

type EventModel struct {
	DB *sql.DB
}

type Event struct {
	Id          int    `json:"id"`
	OwnerId     int    `json:"ownerid" binding:"required"`
	Name        string `json:"name" binding:"required,min=3"`
	Description string `json:"description" binding:"required,min=3,"`
	Date        string `json:"date" binding:"required,datetime=2025-11-14"`
	Location    string `json:"location" binding:"required,min=3"`
}

func (m *EventModel) Insert(event *Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO events (owner_id, name, description, date, location) VALUES ($1, $2, $3, $4, $5)"

	return m.DB.QueryRowContext(ctx, query, event.OwnerId, event.Name, event.Description, event.Date, event.Location).Scan(&event.Id)
}

func (m *EventModel) getAll() ([]*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := "SELECT * FROM events:"

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	events := []*Event{}
	for rows.Next() {

		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Date, &event.Location)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}
