package postgres

import (
	"context"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_db"
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
		err = rows.Scan(&role.Id, &role.RoleName)
		if err != nil {
			logger.Logger.Errorln("Error scanning", err)
			return nil, err
		}
		roleList = append(roleList, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return roleList, nil
}
