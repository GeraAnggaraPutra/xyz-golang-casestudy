package model

import (
	"database/sql"
	"time"
)

type Transaction struct {
	GUID              string         `json:"guid" gorm:"primaryKey"`
	CustomerGUID      string         `json:"customer_guid"`
	ContractNo        string         `json:"contract_no"`
	OTR               float64        `json:"otr"`
	AdminFee          float64        `json:"admin_fee"`
	InstallmentAmount float64        `json:"installment_amount"`
	InterestAmount    float64        `json:"interest_amount"`
	AssetName         string         `json:"asset_name"`
	AssetType         string         `json:"asset_type"`
	CreatedAt         time.Time      `json:"created_at"`
	CreatedBy         string         `json:"created_by"`
	UpdatedAt         sql.NullTime   `json:"updated_at"`
	UpdatedBy         sql.NullString `json:"updated_by"`
	DeletedAt         sql.NullTime   `json:"deleted_at"`
	DeletedBy         sql.NullString `json:"deleted_by"`
	Customer          Customer       `json:"-"`
}
