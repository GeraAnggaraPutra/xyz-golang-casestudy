package payload

import "kredit-plus/src/model"

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (request *RefreshTokenRequest) ToSessionPayload(session model.Session) (
	params SessionPayload,
) {
	params = SessionPayload{
		SessionGUID: session.GUID,
		UserGUID:    session.UserGUID,
	}

	return
}
