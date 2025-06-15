package payload

import (
	"kredit-plus/src/handler/jwt"
)

type SessionPayload struct {
	SessionGUID string `json:"session_guid"`
	UserGUID    string `json:"user_guid"`
}

func (request *SessionPayload) ToAccessTokenRequest() (
	params jwt.AccessTokenPayload,
) {
	params = jwt.AccessTokenPayload{
		GUID:     request.SessionGUID,
		UserGUID: request.UserGUID,
	}

	return
}

func (request *SessionPayload) ToRefreshTokenRequest() (
	params jwt.RefreshTokenPayload,
) {
	params = jwt.RefreshTokenPayload{
		GUID: request.SessionGUID,
	}

	return
}
