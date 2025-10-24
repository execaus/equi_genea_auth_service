package app

import (
	"context"
	authpb "equi_genea_auth_service/internal/pb/api/auth"
	"equi_genea_auth_service/internal/service"

	"google.golang.org/protobuf/types/known/emptypb"
)

type AccountHandler struct {
	service *service.AuthService
	authpb.UnimplementedAuthServiceServer
}

func (h *AccountHandler) GenerateToken(ctx context.Context, in *authpb.GenerateTokenRequest) (*authpb.GenerateTokenResponse, error) {
	token, err := h.service.GenerateJWT(in.Id)
	if err != nil {
		return nil, err
	}

	return &authpb.GenerateTokenResponse{Token: token}, nil
}

func (h *AccountHandler) HashPassword(ctx context.Context, in *authpb.HashPasswordRequest) (*authpb.HashPasswordResponse, error) {
	hash, err := h.service.HashPassword(in.Password)
	if err != nil {
		return nil, err
	}

	return &authpb.HashPasswordResponse{Hash: hash}, nil
}

func (h *AccountHandler) GeneratePassword(ctx context.Context, in *emptypb.Empty) (*authpb.GeneratePasswordResponse, error) {
	password, err := h.service.GeneratePassword()
	if err != nil {
		return nil, err
	}

	return &authpb.GeneratePasswordResponse{Password: password}, nil
}

func (h *AccountHandler) GetClaimsFromToken(ctx context.Context, in *authpb.GetClaimsFromTokenRequest) (*authpb.GetClaimsFromTokenResponse, error) {
	claims, err := h.service.GetClaims(in.Token)
	if err != nil {
		return nil, err
	}

	return &authpb.GetClaimsFromTokenResponse{
		Claims: &authpb.AuthClaims{AccountId: claims.AccountID},
	}, nil
}

func NewAccountHandler(service *service.AuthService) *AccountHandler {
	return &AccountHandler{service: service}
}
