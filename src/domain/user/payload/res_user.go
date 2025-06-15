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

func ToUserResponse(entity model.User) (response UserResponse) {
	response.GUID = entity.GUID
	response.Email = entity.Email
	response.CreatedAt = entity.CreatedAt.Format(time.RFC3339)

	if entity.CreatedBy.Valid {
		response.CreatedBy = &entity.CreatedBy.String
	}

	if entity.UpdatedAt.Valid {
		updatedAt := entity.UpdatedAt.Time.Format(time.RFC3339)
		response.UpdatedAt = &updatedAt
	}

	if entity.UpdatedBy.Valid {
		response.UpdatedBy = &entity.UpdatedBy.String
	}

	return
}

func ToUserResponses(entities []model.User) (response []UserResponse) {
	response = make([]UserResponse, len(entities))

	for i := range entities {
		response[i] = ToUserResponse(entities[i])
	}

	return
}
