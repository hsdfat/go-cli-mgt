package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/models/models_error"
)

func (c *PgClient) UserNeAdd(userNe *models_api.UserNe) error {
	query := `INSERT INTO "user_ne" (user_id, ne_id) VALUES ($1, $2) RETURNING id`
	row := c.pool.QueryRow(context.Background(), query, userNe.UserId, userNe.NeId)

	var id uint
	err := row.Scan(&id)
	if err != nil {
		return err
	}
	userNe.Id = id
	return nil
}

func (c *PgClient) UserNeDelete(userId, neId uint) error {
	query := `DELETE FROM "user_ne" WHERE user_id = $1 AND ne_id = $2`
	_ = c.pool.QueryRow(context.Background(), query, userId, neId)
	return nil
}

func (c *PgClient) UserNeGet(userId, neId uint) (*models_api.UserNe, error) {
	q := `SELECT id, user_id, ne_id FROM "user_ne" WHERE user_id = $1 AND ne_id = $2`
	row := c.pool.QueryRow(context.Background(), q, userId, neId)

	var userNe models_api.UserNe
	err := row.Scan(&userNe.Id, &userNe.UserId, &userNe.NeId)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, models_error.ErrNotFoundUserNe
	} else if err != nil {
		return nil, err
	}

	return &userNe, nil
}

//func (c *PgClient) UserNeGetByUserId(userId uint) ([]models_api.UserNe, error) {
//	q := `SELECT id, user_id, ne_id FROM "user_ne" WHERE user_id = $1`
//	row := c.pool.QueryRow(context.Background(), q, userId)
//
//	var userNeList []models_api.UserNe
//	for
//}
