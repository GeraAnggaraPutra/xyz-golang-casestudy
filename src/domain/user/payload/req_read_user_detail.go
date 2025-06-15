package payload

type ReadUserDetailRequest struct {
	GUID string `param:"guid" validate:"required"`
}
