package db

// User represents user table in database
type User struct {
	ID              uint   `gorm:"column:id;primaryKey;autoIncrement"`
	Username        string `gorm:"column:username;type:varchar(255);uniqueIndex;not null"`
	Email           string `gorm:"column:email;type:varchar(255)"`
	Password        string `gorm:"column:password;type:varchar(255);not null"`
	Active          bool   `gorm:"column:active;default:true"`
	CreatedDateUnix uint64 `gorm:"column:created_date_unix"`
	DisableDateUnix uint64 `gorm:"column:disable_date_unix"`
	DeactivateBy    string `gorm:"column:deactivate_by;type:varchar(255)"`
}

// TableName specifies the table name for User model
func (User) TableName() string {
	return "user"
}
