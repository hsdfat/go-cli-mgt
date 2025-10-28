package db

// Role represents role table in database
type Role struct {
	ID          uint   `gorm:"column:id;primaryKey;autoIncrement"`
	RoleName    string `gorm:"column:role_name;type:varchar(100);uniqueIndex;not null"`
	Description string `gorm:"column:description;type:text"`
}

// TableName specifies the table name for Role model
func (Role) TableName() string {
	return "role"
}
