package postgres

import (
	"context"
	"errors"
	pgxv4 "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v5"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
)

var (
	ErrNotFoundUser = errors.New("user not found")
	ErrDisableUser  = errors.New("user have been disable")
	ErrEnableUser   = errors.New("user have been active")
)

func (c *PgClient) CreateUser(user *models_api.User) error {
	query := `INSERT INTO "user" (username, email, password, active, created_date_unix) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	row := c.pool.QueryRow(context.Background(), query, user.Username, user.Email, user.Password, user.Active, user.CreatedDate)

	var id uint
	err := row.Scan(&id)
	if err != nil {
		return err
	}
	user.Id = id
	return nil
}

func (c *PgClient) DeleteUser(username string) error {
	query := `DELETE FROM "user" WHERE username = $1;`
	_ = c.pool.QueryRow(context.Background(), query, username)
	return nil
}

func (c *PgClient) UpdateUser(user *models_api.User) error {
	query := `UPDATE "user" SET password = $1, active = $2 where id = $3`
	_ = c.pool.QueryRow(context.Background(), query, user.Password, user.Active, user.Id)
	return nil
}

func (c *PgClient) GetUserByID(id uint) (*models_api.User, error) {
	query := `SELECT id, username, email, password FROM "user" WHERE id = $1`
	row := c.pool.QueryRow(context.Background(), query, id)

	var user models_api.User
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFoundUser
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *PgClient) GetUserByUsername(username string) (*models_api.User, error) {
	query := `SELECT "id", "username", "password", "active", "email", "created_date_unix"FROM "user" WHERE username = $1`
	row := c.pool.QueryRow(context.Background(), query, username)

	var user models_api.User
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Active, &user.Email, &user.CreatedDate)
	if errors.Is(err, pgxv4.ErrNoRows) {
		return nil, ErrNotFoundUser
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
