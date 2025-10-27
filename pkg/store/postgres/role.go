package postgres

import (
	"context"
	"errors"
	models_api "go-cli-mgt/pkg/models/api"
	models_db "go-cli-mgt/pkg/models/db"

	pgxv4 "github.com/jackc/pgx/v4"
)

func (c *PgClient) GetRoleByUserId(userId uint) ([]models_db.Role, error) {
	query := `SELECT r.id, r.role_name, r.description, r.priority FROM "role" r JOIN user_role ur ON r.id = ur.role_id JOIN "user" u ON ur.user_id = u.id WHERE u.id = $1`
	rows, err := c.pool.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var roleList []models_db.Role
	for rows.Next() {
		var role models_db.Role
		err = rows.Scan(&role.Id, &role.RoleName, &role.Description, &role.Priority)
		if err != nil {
			return nil, err
		}
		roleList = append(roleList, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return roleList, nil
}

func (c *PgClient) GetRoleByName(roleName string) (*models_api.Role, error) {
	q := `SELECT id, role_name, description, priority FROM "role" WHERE role_name = $1`
	row := c.pool.QueryRow(context.Background(), q, roleName)

	var role models_api.Role
	err := row.Scan(&role.RoleId, &role.RoleName, &role.Description, &role.Priority)
	if errors.Is(err, pgxv4.ErrNoRows) {
		return nil, errors.New("role not found")
	} else if err != nil {
		return nil, err
	}
	return &role, nil
}

func (c *PgClient) CreateRole(role *models_api.Role) error {
	q := `INSERT INTO "role" (role_name, description, priority) VALUES ($1, $2, $3) RETURNING id`
	row := c.pool.QueryRow(context.Background(), q, role.RoleName, role.Description, role.Priority)
	var id uint
	err := row.Scan(&id)
	if err != nil {
		return err
	}
	role.RoleId = id
	return nil
}

func (c *PgClient) DeleteRole(role *models_api.Role) error {
	q := `DELETE FROM "role" WHERE role_name = $1 AND priority = $2`
	_ = c.pool.QueryRow(context.Background(), q, role.RoleName, role.Priority)
	return nil
}

func (c *PgClient) UpdateRole(role *models_api.Role) error {
	q := `UPDATE "role" SET role_name = $1, description = $2, priority = $3 WHERE id = $4 RETURNING id`
	row := c.pool.QueryRow(context.Background(), q, role.RoleName, role.Description, role.Priority, role.RoleId)
	var id uint
	err := row.Scan(&id)
	if err != nil {
		return err
	}
	role.RoleId = id
	return nil
}

func (c *PgClient) GetListRole() ([]models_api.Role, error) {
	q := `SELECT id, role_name, description, priority FROM "role"`
	rows, err := c.pool.Query(context.Background(), q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var roleList []models_api.Role
	for rows.Next() {
		var role models_api.Role
		err = rows.Scan(&role.RoleId, &role.RoleName, &role.Description, &role.Priority)
		if err != nil {
			return nil, err
		}
		roleList = append(roleList, role)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return roleList, nil
}
