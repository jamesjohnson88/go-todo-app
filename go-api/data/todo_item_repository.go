package data

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-todo-app/models"
)

var pool *pgxpool.Pool

func SetDbPool(p *pgxpool.Pool) {
	pool = p
}

func CreateTodoItem(item *models.TodoItem) (dbItem models.TodoItem, err error) {
	query := `
		INSERT INTO todo_items (title, description, priority)
		VALUES ($1, $2, $3)
		RETURNING id, title, description, completed, created_at, updated_at, due_date, priority, user_id;
	`
	err = pool.QueryRow(context.Background(), query, item.Title, item.Description, item.Priority).
		Scan(
			&dbItem.Id,
			&dbItem.Title,
			&dbItem.Description,
			&dbItem.Completed,
			&dbItem.CreatedAt,
			&dbItem.UpdatedAt,
			&dbItem.DueDate,
			&dbItem.Priority,
			&dbItem.UserId,
		)

	if err != nil {
		return dbItem, err
	}

	return dbItem, nil
}

func GetTodoItems() ([]models.TodoItem, error) {
	var dbItems []models.TodoItem

	query := `
		SELECT id, title, description, completed, created_at, updated_at, due_date, priority, user_id
		FROM todo_items;
	`
	rows, err := pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.TodoItem
		err := rows.Scan(
			&item.Id,
			&item.Title,
			&item.Description,
			&item.Completed,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.DueDate,
			&item.Priority,
			&item.UserId,
		)
		if err != nil {
			return nil, err
		}

		dbItems = append(dbItems, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dbItems, nil
}

func GetTodoItemById(id string) (*models.TodoItem, error) {
	dbItem := models.TodoItem{}
	query := `
		SELECT id, title, description, completed, created_at, updated_at, due_date, priority, user_id
		FROM todo_items
		WHERE id = $1;
	`
	err := pool.QueryRow(context.Background(), query, id).
		Scan(
			&dbItem.Id,
			&dbItem.Title,
			&dbItem.Description,
			&dbItem.Completed,
			&dbItem.CreatedAt,
			&dbItem.UpdatedAt,
			&dbItem.DueDate,
			&dbItem.Priority,
			&dbItem.UserId,
		)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &dbItem, nil
}

func CompleteTodoItem(id string) (bool, error) {
	query := `
		UPDATE todo_items
		SET completed = true, updated_at = NOW()
		WHERE id = $1;
	`
	result, err := pool.Exec(context.Background(), query, id)
	if err != nil {
		return false, err
	}

	return result.RowsAffected() > 0, nil
}

func UpdateTodoItem(id string, item *models.TodoItem) (dbItem *models.TodoItem, err error) {
	dbItem = &models.TodoItem{}
	query := `
		UPDATE todo_items
		SET title = $2, description = $3, due_date = $4, priority = $5, updated_at = NOW()
		WHERE id = $1
		RETURNING id, title, description, completed, created_at, updated_at, due_date, priority, user_id;
	`
	err = pool.QueryRow(context.Background(), query, id, item.Title, item.Description, item.DueDate, item.Priority).
		Scan(
			&dbItem.Id,
			&dbItem.Title,
			&dbItem.Description,
			&dbItem.Completed,
			&dbItem.CreatedAt,
			&dbItem.UpdatedAt,
			&dbItem.DueDate,
			&dbItem.Priority,
			&dbItem.UserId,
		)
	if err != nil {
		return nil, err
	}

	return dbItem, nil
}

func DeleteTodoItem(id string) (bool, error) {
	query := `DELETE FROM todo_items WHERE id = $1;`
	result, err := pool.Exec(context.Background(), query, id)
	if err != nil {
		return false, err
	}

	return result.RowsAffected() > 0, nil
}
