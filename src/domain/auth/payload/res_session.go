package payload

import (
	"time"

	"kredit-plus/src/model"
)

type UserResponse struct {
	GUID      string  `json:"guid"`
	Email     string  `json:"email"`
	CreatedAt string  `json:"created_at"`
	CreatedBy *string `json:"created_by"`
	UpdatedAt *string `json:"updated_at"`
	UpdatedBy *string `json:"updated_by"`
}

type SessionResponse struct {
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiredAt  time.Time    `json:"access_token_expired_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiredAt time.Time    `json:"refresh_token_expired_at"`
	User                  UserResponse `json:"user"`
}

func ToSessionResponse(entity model.Session, user model.User) (response SessionResponse) {
	response.AccessToken = entity.AccessToken
	response.AccessTokenExpiredAt = entity.AccessTokenExpiredAt
	response.RefreshToken = entity.RefreshToken
	response.RefreshTokenExpiredAt = entity.RefreshTokenExpiredAt
	response.User.GUID = user.GUID
	response.User.Email = user.Email
	response.User.CreatedAt = user.CreatedAt.Format(time.RFC3339)

	if user.CreatedBy.Valid {
		response.User.CreatedBy = &user.CreatedBy.String
	}

	if user.UpdatedAt.Valid {
		updatedAt := user.UpdatedAt.Time.Format(time.RFC3339)
		response.User.UpdatedAt = &updatedAt
	}

	if user.UpdatedBy.Valid {
		response.User.UpdatedBy = &user.UpdatedBy.String
	}

	return
}
