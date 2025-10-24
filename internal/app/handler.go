package app

import (
	"context"
	authpb "equi_genea_auth_service/internal/pb/api/auth"
	"equi_genea_auth_service/internal/service"
)

type AccountHandler struct {
	service *service.AuthService
	authpb.UnimplementedAuthServiceServer
}

func (h *AccountHandler) GenerateToken(ctx context.Context, in *authpb.GenerateTokenRequest) (*authpb.GenerateTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewAccountHandler(service *service.AuthService) *AccountHandler {
	return &AccountHandler{service: service}
}
