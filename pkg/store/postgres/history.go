package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/models/models_error"
)

func (c *PgClient) SaveHistory(history *models_api.History) error {
	query := `INSERT INTO "operation_history" (username, command, executed_time, user_ip, result, ne_name, mode) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	row := c.pool.QueryRow(context.Background(), query, history.Username, history.Command, history.ExecutedTime, history.UserIp, history.Result, history.NeName, history.Mode)

	var id uint64
	err := row.Scan(&id)
	if err != nil {
		return err
	}
	history.Id = id
	return nil
}

func (c *PgClient) GetHistoryById(id uint64) (*models_api.History, error) {
	query := `SELECT id, username, command, executed_time, user_ip, result, ne_name, mode FROM "operation_history" WHERE id = $1`
	row := c.pool.QueryRow(context.Background(), query, id)
	var history models_api.History
	err := row.Scan(&history.Id, &history.Username, &history.Command, &history.ExecutedTime, &history.UserIp, &history.Result, &history.NeName, &history.Mode)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, models_error.ErrNotFoundUser
	} else if err != nil {
		return nil, err
	}

	return &history, nil
}

func (c *PgClient) DeleteHistoryById(id uint64) error {
	query := `DELETE FROM "operation_history" WHERE id = $1`
	_ = c.pool.QueryRow(context.Background(), query, id)
	return nil
}

func (c *PgClient) GetHistoryListByMode(mode string) ([]models_api.History, error) {
	q := `SELECT id, username, command, executed_time, user_ip, result, ne_name, mode FROM "operation_history" WHERE mode = $1`
	rows, err := c.pool.Query(context.Background(), q, mode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var histories []models_api.History
	for rows.Next() {
		var history models_api.History
		err = rows.Scan(&history.Id, &history.Username, &history.Command, &history.ExecutedTime, &history.UserIp, &history.Result, &history.NeName, &history.Mode)
		if err != nil {
			return nil, err
		}
		histories = append(histories, history)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return histories, nil
}
