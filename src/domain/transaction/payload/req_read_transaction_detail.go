package payload

type ReadTransactionDetailRequest struct {
	GUID string `param:"guid" validate:"required"`
}
