package repository

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"rental_mobile_fiber/internal/entity"
	"time"
)

type EventRepository struct {
}

func NewEventRepository() *EventRepository {
	return &EventRepository{}
}

func (e *EventRepository) Create(ctx context.Context, request entity.Event, db *sql.DB) (int, error) {
	sqlCommand := "insert into events (title, thumbnail, start_date, end_date, description, max_join, price, is_publish, location) values (?,?,?,?,?,?,?,?,?)"
	row, err := db.ExecContext(ctx, sqlCommand, request.Thumbnail, request.Title, request.StartDate, request.EndDate, request.Description, request.MaxJoin, request.Price, request.IsPublish, request.Location)
	if err != nil {
		return 0, err
	}

	eventId, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(eventId), nil
}

func (e *EventRepository) Update(ctx context.Context, event entity.Event, db *sql.DB) error {
	sqlCommand := "update events set  thumbnail = ?, title = ?, start_date = ?, end_date = ?, description = ?, max_join = ?, price = ?, is_publish = ?, location = ? where id = ?"

	_, err := db.ExecContext(
		ctx,
		sqlCommand,
		event.Thumbnail,
		event.Title,
		event.StartDate,
		event.EndDate,
		event.Description,
		event.MaxJoin,
		event.Price,
		event.IsPublish,
		event.Location,
		event.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (e *EventRepository) Delete(ctx context.Context, eventId int, db *sql.DB) error {
	sqlCommand := "delete from events where id = ?"

	_, err := db.ExecContext(ctx, sqlCommand, eventId)
	if err != nil {
		return err
	}

	return nil
}

func (e *EventRepository) Get(ctx context.Context, eventId int, db *sql.DB) (entity.Event, error) {
	SqlCommand := "select id, title, thumbnail, start_date, end_date, description, max_join, price, is_publish, location from events where id = ?"
	row, err := db.QueryContext(ctx, SqlCommand, eventId)

	var event entity.Event

	if err != nil {
		return event, err
	}

	if row.Next() {
		var startDateRaw []uint8
		var endDateRaw []uint8
		err := row.Scan(&event.ID, &event.Title, &event.Thumbnail, &startDateRaw, &endDateRaw, &event.Description, &event.MaxJoin, &event.Price, &event.IsPublish, &event.Location)
		if err != nil {
			return event, fiber.ErrInternalServerError
		}

		// Convert []uint8 to time.Time
		startDate, err := time.Parse("2006-01-02 15:04:05", string(startDateRaw))
		if err != nil {
			return event, err
		}

		event.StartDate = startDate

		// Convert []uint8 to time.Time
		endDate, err := time.Parse("2006-01-02 15:04:05", string(endDateRaw))
		if err != nil {
			return event, err
		}

		event.EndDate = endDate
	} else {
		return event, fiber.ErrNotFound
	}

	return event, nil
}

func (e *EventRepository) All(ctx context.Context, db *sql.DB) ([]entity.Event, error) {
	SqlCommand := "select id, title, thumbnail, start_date, end_date, description, max_join, price, is_publish, location from events"
	row, err := db.QueryContext(ctx, SqlCommand)

	var events []entity.Event

	if err != nil {
		return events, err
	}

	for row.Next() {
		var event entity.Event
		var startDateRaw []uint8
		var endDateRaw []uint8
		err := row.Scan(&event.ID, &event.Title, &event.Thumbnail, &startDateRaw, &endDateRaw, &event.Description, &event.MaxJoin, &event.Price, &event.IsPublish, &event.Location)
		if err != nil {
			return events, fiber.ErrInternalServerError
		}

		// Convert []uint8 to time.Time
		startDate, err := time.Parse("2006-01-02 15:04:05", string(startDateRaw))
		if err != nil {
			return events, err
		}

		event.StartDate = startDate

		// Convert []uint8 to time.Time
		endDate, err := time.Parse("2006-01-02 15:04:05", string(endDateRaw))
		if err != nil {
			return events, err
		}

		event.EndDate = endDate

		events = append(events, event)
	}

	return events, nil
}
