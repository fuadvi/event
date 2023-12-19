package repository

import (
	"context"
	"database/sql"
	"rental_mobile_fiber/internal/entity"
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

func (e EventRepository) Update(ctx context.Context, event entity.Event, db *sql.DB) error {
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

func (e EventRepository) Delete(ctx context.Context, eventId int, db *sql.DB) error {
	sqlCommand := "delete from events where id = ?"

	_, err := db.ExecContext(ctx, sqlCommand, eventId)
	if err != nil {
		return err
	}

	return nil
}
