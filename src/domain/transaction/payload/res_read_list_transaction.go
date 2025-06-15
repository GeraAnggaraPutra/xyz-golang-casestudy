package payload

import (
	"time"

	"kredit-plus/src/model"
)

type ReadListTransactionResponse struct {
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
}

func ToReadListTransactionResponse(entity model.Transaction) (response ReadListTransactionResponse) {
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

	if entity.UpdatedAt.Valid {
		response.UpdatedAt = &entity.UpdatedAt.Time
	}

	if entity.UpdatedBy.Valid {
		response.UpdatedBy = &entity.UpdatedBy.String
	}

	return
}

func ToReadListTransactionResponses(entities []model.Transaction) (response []ReadListTransactionResponse) {
	response = make([]ReadListTransactionResponse, len(entities))

	for i := range entities {
		response[i] = ToReadListTransactionResponse(entities[i])
	}

	return
}
