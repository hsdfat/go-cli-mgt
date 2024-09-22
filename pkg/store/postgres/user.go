package postgres

import (
	"context"
	"fmt"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"

	"github.com/jackc/pgx/v5"
)

func (c *PgClient) CreateUser(user *models_api.User) error {
	query := `INSERT INTO "user" (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	row := c.pool.QueryRow(context.Background(), query, user.Username, user.Email, user.Password)

	var id uint
	err := row.Scan(&id)
	if err != nil {
		return err
	}
	user.Id = id
	return nil
}

func (c *PgClient) GetUserByID(id uint) (*models_api.User, error) {
	query := `SELECT id, username, email, password FROM "user" WHERE id = $1`
	row := c.pool.QueryRow(context.Background(), query, id)

	var user models_api.User
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *PgClient) ListUsers() ([]models_api.User, error) {
	query := `SELECT id, username, email FROM "user"`
	rows, err := c.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models_api.User
	for rows.Next() {
		var user models_api.User
		err = rows.Scan(&user.Id, &user.Username, &user.Email)
		if err != nil {
			logger.Logger.Errorln("Error scanning", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}