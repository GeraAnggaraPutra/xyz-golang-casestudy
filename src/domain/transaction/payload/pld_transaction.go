package payload

type TransactionPayload struct {
	CustomerGUID      string  `json:"customer_guid" validate:"required"`
	OTR               float64 `json:"otr" validate:"required"`
	AssetName         string  `json:"asset_name" validate:"required"`
	AssetType         string  `json:"asset_type" validate:"required,oneof=Motor Mobil WhiteGoods"`
	TenorMonths       int     `json:"tenor_months" validate:"required,gte=1,lte=30"`
}
