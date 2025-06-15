package payload

import (
	"fmt"
	"os"
)

type ReadFileUploadResponse struct {
	URL string `json:"url"`
}

func ToReadFileUploadResponse(path string) (res ReadFileUploadResponse) {
	url := fmt.Sprintf("%s/storage/file/%s", os.Getenv("BACKEND_URL"), path)

	res = ReadFileUploadResponse{
		URL: url,
	}

	return
}
