package postgres

import (
	"context"
	"errors"
	models_api "go-cli-mgt/pkg/models/api"

	"github.com/jackc/pgx/v4"
)

func (c *PgClient) CreateNetworkElement(ne *models_api.NeData) error {
	query := `INSERT INTO "network_element" (name, type, namespace, master_ip_config, master_port_config, slave_ip_config, slave_port_config, base_url, ip_command, port_command) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	row := c.pool.QueryRow(context.Background(), query, ne.Name, ne.Type, ne.Namespace, ne.MasterIpConfig, ne.MasterPortConfig, ne.SlaveIpConfig, ne.SlavePortConfig, ne.Url, ne.IpCommand, ne.PortCommand)

	var id uint
	err := row.Scan(&id)
	if err != nil {
		return err
	}
	ne.NeId = id
	return nil
}

func (c *PgClient) DeleteNetworkElementByName(neName string, namespace string) error {
	query := `DELETE FROM "network_element" WHERE name = $1 AND namespace = $2`
	_ = c.pool.QueryRow(context.Background(), query, neName, namespace)
	return nil
}

func (c *PgClient) GetNetworkElementByName(neName string, namespace string) (*models_api.NeData, error) {
	query := `SELECT id, name, type, namespace, master_ip_config, master_port_config, slave_ip_config, slave_port_config, base_url, ip_command, port_command FROM "network_element" WHERE name = $1 AND namespace = $2`
	row := c.pool.QueryRow(context.Background(), query, neName, namespace)

	var ne models_api.NeData
	err := row.Scan(&ne.NeId, &ne.Name, &ne.Type, &ne.Namespace, &ne.MasterIpConfig, &ne.MasterPortConfig, &ne.SlaveIpConfig, &ne.SlavePortConfig, &ne.Url, &ne.IpCommand, &ne.PortCommand)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("network element not found")
	} else if err != nil {
		return nil, err
	}
	return &ne, nil
}

func (c *PgClient) GetListNetworkElement() ([]models_api.NeData, error) {
	query := `SELECT id, name, type, namespace, master_ip_config, master_port_config, slave_ip_config, slave_port_config, base_url, ip_command, port_command FROM "network_element"`
	rows, err := c.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var neList []models_api.NeData
	for rows.Next() {
		var ne models_api.NeData
		err = rows.Scan(&ne.NeId, &ne.Name, &ne.Type, &ne.Namespace, &ne.MasterIpConfig, &ne.MasterPortConfig, &ne.SlaveIpConfig, &ne.SlavePortConfig, &ne.Url, &ne.IpCommand, &ne.PortCommand)
		if err != nil {
			return nil, err
		}
		neList = append(neList, ne)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return neList, nil
}

func (c *PgClient) GetNetworkElementByUserName(userName string) ([]models_api.NeData, error) {
	query := `SELECT ne.id, ne.name, ne.type, ne.namespace, ne.master_ip_config, ne.master_port_config, ne.slave_ip_config, ne.slave_port_config, ne.base_url, ne.ip_command, ne.port_command FROM "network_element" ne JOIN "user_ne" un ON ne.id = un.ne_id JOIN "user" u ON un.user_id = u.id WHERE u.username = $1`
	rows, err := c.pool.Query(context.Background(), query, userName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var neList []models_api.NeData
	for rows.Next() {
		var ne models_api.NeData
		err = rows.Scan(&ne.NeId, &ne.Name, &ne.Type, &ne.Namespace, &ne.MasterIpConfig, &ne.MasterPortConfig, &ne.SlaveIpConfig, &ne.SlavePortConfig, &ne.Url, &ne.IpCommand, &ne.PortCommand)
		if err != nil {
			return nil, err
		}
		neList = append(neList, ne)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return neList, nil
}
