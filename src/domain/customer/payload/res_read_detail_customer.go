package payload

import (
	"time"

	"kredit-plus/src/model"
	"kredit-plus/src/util"
)

type ReadDetailCustomerResponse struct {
	GUID        string          `json:"guid"`
	NIK         string          `json:"nik"`
	FullName    string          `json:"full_name"`
	LegalName   string          `json:"legal_name"`
	BirthPlace  string          `json:"birth_place"`
	BirthDate   time.Time       `json:"birth_date"`
	Salary      float64         `json:"salary"`
	PhotoKTP    string          `json:"photo_ktp"`
	PhotoSelfie string          `json:"photo_selfie"`
	CreatedAt   time.Time       `json:"created_at"`
	CreatedBy   *string         `json:"created_by"`
	UpdatedAt   *time.Time      `json:"updated_at"`
	UpdatedBy   *string         `json:"updated_by"`
	Tenor       []CustomerTenor `json:"tenors"`
}

func ToReadDetailCustomerResponse(entity model.Customer) (response ReadDetailCustomerResponse) {
	response.GUID = entity.GUID
	response.NIK = entity.NIK
	response.FullName = entity.FullName
	response.LegalName = entity.LegalName
	response.BirthPlace = entity.BirthPlace
	response.BirthDate = entity.BirthDate
	response.Salary = entity.Salary
	response.PhotoKTP = util.MakeFullURL(entity.PhotoKTP)
	response.PhotoSelfie = util.MakeFullURL(entity.PhotoSelfie)
	response.CreatedAt = entity.CreatedAt
	response.CreatedBy = &entity.CreatedBy

	if entity.UpdatedAt.Valid {
		response.UpdatedAt = &entity.UpdatedAt.Time
	}

	if entity.UpdatedBy.Valid {
		response.UpdatedBy = &entity.UpdatedBy.String
	}

	if len(entity.TenorLimits) > 0 {
		response.Tenor = make([]CustomerTenor, len(entity.TenorLimits))
		for i, limit := range entity.TenorLimits {
			response.Tenor[i] = CustomerTenor{
				TenorMonths: limit.TenorMonths,
				LimitAmount: limit.LimitAmount,
			}
		}
	}

	return
}
