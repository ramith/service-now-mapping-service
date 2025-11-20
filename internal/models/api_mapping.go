package models

import (
	"time"

	"gorm.io/gorm"
)

// APIMapping represents the mapping between ServiceNow system ID and API configuration
type APIMapping struct {
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	SysID      string         `gorm:"type:varchar(32);uniqueIndex:uk_sys_id;not null" json:"sys_id" binding:"required"`
	APIID      string         `gorm:"type:varchar(100);index:idx_api_id;not null" json:"api_id" binding:"required"`
	APIName    string         `gorm:"type:varchar(255);uniqueIndex:uk_api_name_version;not null" json:"api_name" binding:"required"`
	APIVersion string         `gorm:"type:varchar(50);uniqueIndex:uk_api_name_version;not null" json:"api_version" binding:"required"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index:idx_deleted_at" json:"deleted_at,omitempty"`
}

// TableName specifies the table name for GORM
func (APIMapping) TableName() string {
	return "api_mappings"
}

// CreateAPIMapping represents the request payload for creating a new mapping
type CreateAPIMapping struct {
	SysID      string `json:"sys_id" binding:"required,min=1,max=32"`
	APIID      string `json:"api_id" binding:"required,min=1,max=100"`
	APIName    string `json:"api_name" binding:"required,min=1,max=255"`
	APIVersion string `json:"api_version" binding:"required,min=1,max=50"`
}

// UpdateAPIMapping represents the request payload for updating a mapping
type UpdateAPIMapping struct {
	APIID      string `json:"api_id" binding:"omitempty,min=1,max=100"`
	APIName    string `json:"api_name" binding:"omitempty,min=1,max=255"`
	APIVersion string `json:"api_version" binding:"omitempty,min=1,max=50"`
}

// SearchByNameVersion represents search parameters for API name and version
type SearchByNameVersion struct {
	APIName    string `form:"api_name" binding:"required"`
	APIVersion string `form:"api_version" binding:"required"`
}
