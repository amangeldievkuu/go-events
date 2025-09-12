package database

import (
	"context"
	"database/sql"
	"time"
)

type AttendeeModel struct {
	DB *sql.DB
}

type Attendee struct {
	Id      int `json:"id"`
	UserId  int `json:"userId"`
	EventId int `json:"eventId"`
}

func (m *AttendeeModel) Insert(attendee *Attendee) (*Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `INSERT INTO attendees (user_id, event_id) VALUES (?, ?)`

	err := m.DB.QueryRowContext(ctx, query, attendee.UserId, attendee.EventId).Scan(&attendee.Id)
	if err != nil {
		return nil, err
	}

	return attendee, nil

}

func (m *AttendeeModel) GetByEventAndAttendee(eventId, userId int) (*Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `SELECT * FROM attendees where event_id = ? AND user_id = ?`

	var attendee Attendee

	err := m.DB.QueryRowContext(ctx, query, eventId, userId).Scan(&attendee.EventId, &attendee.UserId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &attendee, nil
}
