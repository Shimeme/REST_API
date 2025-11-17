package database

import "database/sql"

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
