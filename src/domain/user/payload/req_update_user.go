package payload

type UpdateUserRequest struct {
	GUID string `param:"guid" validate:"required"`
	UserPayload
}
