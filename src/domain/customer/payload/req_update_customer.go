package payload

type UpdateCustomerRequest struct {
	GUID string `param:"guid" validate:"required"`
	CustomerPayload
}
