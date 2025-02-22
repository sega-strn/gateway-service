package auth

import (
	"context"

	"github.com/s21platform/gateway-service/internal/rpc/auth"
)

type Usecase struct {
	aC AuthClient
}

func New(aC AuthClient) *Usecase {
	return &Usecase{aC: aC}
}

func (uc *Usecase) Login(ctx context.Context, username string, password string) (*auth.JWT, error) {
	return uc.aC.DoLogin(ctx, username, password)
}
