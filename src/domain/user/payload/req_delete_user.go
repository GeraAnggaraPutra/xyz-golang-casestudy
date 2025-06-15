package payload

type DeleteUserRequest struct {
	GUID string `param:"guid" validate:"required"`
}
