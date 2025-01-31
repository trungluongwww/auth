// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGooseDbVersion = "goose_db_version"

// GooseDbVersion mapped from table <goose_db_version>
type GooseDbVersion struct {
	ID        uint64     `gorm:"column:id;primaryKey;autoIncrement:true;uniqueIndex:id,priority:1" json:"id"`
	VersionID int64      `gorm:"column:version_id;not null" json:"version_id"`
	IsApplied bool       `gorm:"column:is_applied;not null" json:"is_applied"`
	Tstamp    *time.Time `gorm:"column:tstamp;default:CURRENT_TIMESTAMP" json:"tstamp"`
}

// TableName GooseDbVersion's table name
func (*GooseDbVersion) TableName() string {
	return TableNameGooseDbVersion
}
