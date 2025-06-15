package payload

import (
	"mime/multipart"
)

type CreateFileUploadRequest struct {
	File *multipart.FileHeader `form:"file" validate:"required"`
	Path string                `form:"path" validate:"required"`
}
