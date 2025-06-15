package model

import (
	"database/sql"
	"time"
)

type Customer struct {
	GUID        string               `json:"guid" gorm:"primaryKey"`
	NIK         string               `json:"nik"`
	FullName    string               `json:"full_name"`
	LegalName   string               `json:"legal_name"`
	BirthPlace  string               `json:"birth_place"`
	BirthDate   time.Time            `json:"birth_date"`
	Salary      float64              `json:"salary"`
	PhotoKTP    string               `json:"photo_ktp"`
	PhotoSelfie string               `json:"photo_selfie"`
	CreatedAt   time.Time            `json:"created_at"`
	CreatedBy   string               `json:"created_by"`
	UpdatedAt   sql.NullTime         `json:"updated_at"`
	UpdatedBy   sql.NullString       `json:"updated_by"`
	DeletedAt   sql.NullTime         `json:"deleted_at"`
	DeletedBy   sql.NullString       `json:"deleted_by"`
	TenorLimits []CustomerTenorLimit `json:"customer_tenor_limits" gorm:"foreignKey:CustomerGUID;references:GUID"`
}
