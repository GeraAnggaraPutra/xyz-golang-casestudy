package payload

type CustomerPayload struct {
	NIK         string          `json:"nik" validate:"required"`
	FullName    string          `json:"full_name" validate:"required"`
	LegalName   string          `json:"legal_name" validate:"required"`
	BirthPlace  string          `json:"birth_place" validate:"required"`
	BirthDate   string          `json:"birth_date" validate:"required"`
	Salary      float64         `json:"salary" validate:"required"`
	PhotoKTP    string          `json:"photo_ktp" validate:"required"`
	PhotoSelfie string          `json:"photo_selfie" validate:"required"`
	Tenors      []CustomerTenor `json:"tenors" validate:"required"`
}

type CustomerTenor struct {
	TenorMonths int     `json:"tenor_months"`
	LimitAmount float64 `json:"limit_amount"`
}
