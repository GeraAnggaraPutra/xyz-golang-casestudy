package payload

import "kredit-plus/src/util"

type LoginRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	UserAgent string
	IPAddress string
}

func (request *LoginRequest) ToSessionPayload(userGUID string) (
	params SessionPayload,
) {
	params = SessionPayload{
		SessionGUID: util.GenerateUUID(),
		UserGUID:    userGUID,
		UserAgent:   request.UserAgent,
		IPAddress:   request.IPAddress,
	}

	return
}
