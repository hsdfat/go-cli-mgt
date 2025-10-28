package db

// UserNe represents user_ne table (many-to-many relationship)
type UserNe struct {
	ID     uint `gorm:"column:id;primaryKey;autoIncrement"`
	UserID uint `gorm:"column:user_id;not null;index:idx_user_ne"`
	NeID   uint `gorm:"column:ne_id;not null;index:idx_user_ne"`
}

// TableName specifies the table name for UserNe model
func (UserNe) TableName() string {
	return "user_ne"
}
