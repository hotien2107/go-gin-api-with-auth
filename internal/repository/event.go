package repository

import (
	"gin-rest-api.com/basic/internal/db"
	"gin-rest-api.com/basic/internal/models"
)

type EventRepository struct{}

func NewEventRepository() *EventRepository {
	return &EventRepository{}
}

func (*EventRepository) GetAll() ([]models.Event, error) {
	// query string
	query := `
		SELECT * FROM events
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return []models.Event{}, err
	}
	defer rows.Close()

	var events []models.Event

	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return events, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (*EventRepository) GetById(eventID int64) (models.Event, error) {
	// query string
	query := `
		SELECT * FROM events
		WHERE id=?
	`

	row := db.DB.QueryRow(query, eventID)

	var event models.Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return models.Event{}, err
	}

	return event, nil
}

func (*EventRepository) Save(event *models.Event) (int64, error) {
	//query string
	query := `
		INSERT INTO events(name, description, location, dateTime, userId)
		VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserId)
	if err != nil {
		return 0, err
	}

	eventId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return eventId, nil
}

func (*EventRepository) Update(event *models.Event) error {
	//query string
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
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

func (*EventRepository) Delete(eventId int64) error {
	//query string
	query := `
		DELETE FROM events
		WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(eventId)
	if err != nil {
		return err
	}

	return nil
}
