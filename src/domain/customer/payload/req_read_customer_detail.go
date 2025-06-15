package payload

type ReadCustomerDetailRequest struct {
	GUID string `param:"guid" validate:"required"`
}
