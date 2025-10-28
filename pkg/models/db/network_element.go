package db

// NetworkElement represents network_element table in database
type NetworkElement struct {
	ID               uint   `gorm:"column:id;primaryKey;autoIncrement"`
	Name             string `gorm:"column:name;type:varchar(255);not null"`
	Type             string `gorm:"column:type;type:varchar(100)"`
	Namespace        string `gorm:"column:namespace;type:varchar(255)"`
	MasterIpConfig   string `gorm:"column:master_ip_config;type:varchar(50)"`
	MasterPortConfig string `gorm:"column:master_port_config;type:varchar(10)"`
	SlaveIpConfig    string `gorm:"column:slave_ip_config;type:varchar(50)"`
	SlavePortConfig  string `gorm:"column:slave_port_config;type:varchar(10)"`
	BaseURL          string `gorm:"column:base_url;type:text"`
	IpCommand        string `gorm:"column:ip_command;type:varchar(50)"`
	PortCommand      string `gorm:"column:port_command;type:varchar(10)"`
}

// TableName specifies the table name for NetworkElement model
func (NetworkElement) TableName() string {
	return "network_element"
}
