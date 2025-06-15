package model

import (
	"database/sql"
	"time"
)

type User struct {
	GUID      string         `json:"guid" gorm:"primaryKey"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	CreatedBy sql.NullString `json:"created_by"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
	UpdatedBy sql.NullString `json:"updated_by"`
	DeletedAt sql.NullTime   `json:"deleted_at"`
	DeletedBy sql.NullString `json:"deleted_by"`
}
