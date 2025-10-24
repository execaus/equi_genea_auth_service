package app

import (
	"context"
	authpb "equi_genea_auth_service/internal/pb/api/auth"
	"equi_genea_auth_service/internal/service"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AccountHandler struct {
	service *service.AuthService
	authpb.UnimplementedAuthServiceServer
}

func (h *AccountHandler) GenerateToken(ctx context.Context, in *authpb.GenerateTokenRequest, opts ...grpc.CallOption) (*authpb.GenerateTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h *AccountHandler) HashPassword(ctx context.Context, in *authpb.HashPasswordRequest, opts ...grpc.CallOption) (*authpb.HashPasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h *AccountHandler) GeneratePassword(ctx context.Context, in *emptypb.Empty) (*authpb.GeneratePasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewAccountHandler(service *service.AuthService) *AccountHandler {
	return &AccountHandler{service: service}
}
