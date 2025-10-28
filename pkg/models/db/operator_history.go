package db

import "time"

// OperationHistory represents operation_history table in database
type OperationHistory struct {
	ID           uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	Username     string    `gorm:"column:username;type:varchar(255);not null"`
	Command      string    `gorm:"column:command;type:text;not null"`
	ExecutedTime time.Time `gorm:"column:executed_time;not null"`
	UserIP       string    `gorm:"column:user_ip;type:varchar(50)"`
	Result       string    `gorm:"column:result;type:text"`
	NeName       string    `gorm:"column:ne_name;type:varchar(255)"`
	Mode         string    `gorm:"column:mode;type:varchar(50)"`
}

// TableName specifies the table name for OperationHistory model
func (OperationHistory) TableName() string {
	return "operation_history"
}
