package user

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	userproto "github.com/s21platform/user-proto/user-proto"

	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/model"
)

type Service struct {
	client userproto.UserServiceClient
}

func NewService(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.User.Host, cfg.User.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	client := userproto.NewUserServiceClient(conn)
	return &Service{client: client}
}

func (s *Service) GetInfo(ctx context.Context, uuid string) (*userproto.GetUserInfoByUUIDOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	resp, err := s.client.GetUserInfoByUUID(ctx, &userproto.GetUserInfoByUUIDIn{
		Uuid: uuid,
	})
	if err != nil {
		log.Printf("failed to call: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Service) UpdateProfile(ctx context.Context, data model.ProfileData) (*userproto.UpdateProfileOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	resp, err := s.client.UpdateProfile(ctx, data.FromDTO())
	if err != nil {
		return nil, fmt.Errorf("failed to update user profile: %v", err)
	}
	return resp, nil
}
