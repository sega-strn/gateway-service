package avatar

import (
	"fmt"
	"net/http"

	avatar "github.com/s21platform/avatar-proto/avatar-proto"
)

type Usecase struct {
	aC AvatarClient
}

func New(aC AvatarClient) *Usecase {
	return &Usecase{aC: aC}
}

func (uc *Usecase) UploadAvatar(r *http.Request) (*avatar.SetAvatarOut, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return nil, fmt.Errorf("failed to parse multipart form: %w", err)
	}
	file, _, err := r.FormFile("avatar")
	if err != nil {
		return nil, fmt.Errorf("failed to get file from form: %w", err)
	}
	defer file.Close()

	filename := r.FormValue("filename")

	resp, err := uc.aC.SetAvatar(r.Context(), filename, file)
	if err != nil {
		return nil, fmt.Errorf("failed to set avatar: %w", err)
	}
	return resp, nil
}
