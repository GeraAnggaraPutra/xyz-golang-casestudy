package payload

import (
	"time"

	"kredit-plus/src/model"
	"kredit-plus/src/util"
)

type ReadDetailTransactionResponse struct {
	GUID              string     `json:"guid"`
	CustomerGUID      string     `json:"customer_guid"`
	ContractNo        string     `json:"contract_no"`
	OTR               float64    `json:"otr"`
	AdminFee          float64    `json:"admin_fee"`
	InstallmentAmount float64    `json:"installment_amount"`
	InterestAmount    float64    `json:"interest_amount"`
	AssetName         string     `json:"asset_name"`
	AssetType         string     `json:"asset_type"`
	TenorMonths       int        `json:"tenor_months"`
	CreatedAt         time.Time  `json:"created_at"`
	CreatedBy         *string    `json:"created_by"`
	UpdatedAt         *time.Time `json:"updated_at"`
	UpdatedBy         *string    `json:"updated_by"`
	Customer          Customer   `json:"customer"`
}

type Customer struct {
	GUID        string    `json:"guid"`
	NIK         string    `json:"nik"`
	FullName    string    `json:"full_name"`
	LegalName   string    `json:"legal_name"`
	BirthPlace  string    `json:"birth_place"`
	BirthDate   time.Time `json:"birth_date"`
	Salary      float64   `json:"salary"`
	PhotoKTP    string    `json:"photo_ktp"`
	PhotoSelfie string    `json:"photo_selfie"`
}

func ToReadDetailTransactionResponse(entity model.Transaction) (response ReadDetailTransactionResponse) {
	response.GUID = entity.GUID
	response.CustomerGUID = entity.CustomerGUID
	response.ContractNo = entity.ContractNo
	response.OTR = entity.OTR
	response.AdminFee = entity.AdminFee
	response.InstallmentAmount = entity.InstallmentAmount
	response.InterestAmount = entity.InterestAmount
	response.AssetName = entity.AssetName
	response.AssetType = entity.AssetType
	response.TenorMonths = entity.TenorMonths
	response.CreatedAt = entity.CreatedAt
	response.CreatedBy = &entity.CreatedBy
	response.Customer.GUID = entity.Customer.GUID
	response.Customer.NIK = entity.Customer.NIK
	response.Customer.FullName = entity.Customer.FullName
	response.Customer.LegalName = entity.Customer.LegalName
	response.Customer.BirthPlace = entity.Customer.BirthPlace
	response.Customer.BirthDate = entity.Customer.BirthDate
	response.Customer.Salary = entity.Customer.Salary
	response.Customer.PhotoKTP = util.MakeFullURL(entity.Customer.PhotoKTP)
	response.Customer.PhotoSelfie = util.MakeFullURL(entity.Customer.PhotoSelfie)

	if entity.UpdatedAt.Valid {
		response.UpdatedAt = &entity.UpdatedAt.Time
	}

	if entity.UpdatedBy.Valid {
		response.UpdatedBy = &entity.UpdatedBy.String
	}

	return
}
