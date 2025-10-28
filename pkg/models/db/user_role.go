package db

// UserRole represents user_role table (many-to-many relationship)
type UserRole struct {
	ID     uint `gorm:"column:id;primaryKey;autoIncrement"`
	UserID uint `gorm:"column:user_id;not null;index:idx_user_role"`
	RoleID uint `gorm:"column:role_id;not null;index:idx_user_role"`
}

// TableName specifies the table name for UserRole model
func (UserRole) TableName() string {
	return "user_role"
}
