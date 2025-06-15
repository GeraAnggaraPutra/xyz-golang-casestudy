package payload

type DeleteCustomerRequest struct {
	GUID string `param:"guid" validate:"required"`
}
