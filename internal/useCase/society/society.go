package society

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	society_proto "github.com/s21platform/society-proto/society-proto"
)

type UseCase struct {
	sC SocietyClient
}

func New(sC SocietyClient) *UseCase {
	return &UseCase{sC: sC}
}

type RequestData struct {
	Name          string `json:"Name"`
	Description   string `json:"Description"`
	IsPrivate     bool   `json:"IsPrivate"`
	DirectionId   int64  `json:"DirectionId"`
	AccessLevelId int64  `json:"AccessLevelId"`
}

func (u *UseCase) CreateSociety(r *http.Request) (*society_proto.SetSocietyOut, error) {
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
