package model

import (
	"database/sql"
	"time"
)

type CustomerTenorLimit struct {
	CustomerGUID string         `json:"customer_guid" gorm:"primaryKey"`
	TenorMonths  int            `json:"tenor_months" gorm:"primaryKey"`
	LimitAmount  float64        `json:"limit_amount"`
	CreatedAt    time.Time      `json:"created_at"`
	CreatedBy    string         `json:"created_by"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
	UpdatedBy    sql.NullString `json:"updated_by"`
	Customer     Customer       `json:"-"`
}
