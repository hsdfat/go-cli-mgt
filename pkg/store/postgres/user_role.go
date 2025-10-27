package postgres

import (
	"context"
	"errors"
	models_api "go-cli-mgt/pkg/models/api"

	"github.com/jackc/pgx/v4"
)

func (c *PgClient) UserRoleAdd(userRole *models_api.UserRole) error {
	query := `INSERT INTO "user_role" (user_id, role_id) VALUES ($1, $2) RETURNING id`
	row := c.pool.QueryRow(context.Background(), query, userRole.UserId, userRole.RoleId)

	var id uint
	err := row.Scan(&id)
	if err != nil {
		return err
	}
	userRole.Id = id
	return nil
}

func (c *PgClient) UserRoleGet(userId, roleId uint) (*models_api.UserRole, error) {
	q := `SELECT id, user_id, role_id FROM "user_role" WHERE user_id = $1 AND role_id = $2`
	row := c.pool.QueryRow(context.Background(), q, userId, roleId)

	var userRole models_api.UserRole
	err := row.Scan(&userRole.Id, &userRole.UserId, &userRole.RoleId)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("user do not have this role")
	} else if err != nil {
		return nil, err
	}

	return &userRole, nil
}

func (c *PgClient) UserRoleDelete(userId, roleId uint) {
	query := `DELETE FROM "user_role" WHERE user_id = $1 AND role_id = $2`
	_ = c.pool.QueryRow(context.Background(), query, userId, roleId)
}
