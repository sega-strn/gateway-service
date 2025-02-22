package society

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	societyproto "github.com/s21platform/society-proto/society-proto"
)

type UseCase struct {
	sC SocietyClient
}

func New(sC SocietyClient) *UseCase {
	return &UseCase{sC: sC}
}

type RequestData struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	IsPrivate     bool   `json:"is_private"`
	DirectionId   int64  `json:"direction_id"`
	AccessLevelId int64  `json:"access_level_id"`
}

type SocietyId struct {
	Id int64 `json:"id"`
}

func (u *UseCase) CreateSociety(r *http.Request) (*societyproto.SetSocietyOut, error) {
	requestData := RequestData{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}
	defer r.Body.Close()

	if len(body) == 0 {
		return nil, fmt.Errorf("request body is empty")
	}

	if err := json.Unmarshal(body, &requestData); err != nil {
		return nil, fmt.Errorf("failed to decode request body: %w", err)
	}

	resp, err := u.sC.CreateSociety(r.Context(), &requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to create society: %v", err)
	}
	return resp, nil
}

func (u *UseCase) GetAccessLevel(r *http.Request) (*societyproto.GetAccessLevelOut, error) {
	resp, err := u.sC.GetAccessLevel(r.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get access level: %v", err)
	}
	return resp, nil
}

func (u *UseCase) GetSocietyInfo(r *http.Request) (*societyproto.GetSocietyInfoOut, error) {
	id := SocietyId{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}

	if err := json.Unmarshal(body, &id); err != nil {
		return nil, fmt.Errorf("failed to decode request body: %w", err)
	}

	resp, err := u.sC.GetSocietyInfo(r.Context(), id.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get access level: %v", err)
	}

	return resp, nil
}
