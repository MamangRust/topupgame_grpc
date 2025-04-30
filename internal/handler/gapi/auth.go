package gapi

import (
	"context"
	"topup_game/internal/domain/requests"
	protomapper "topup_game/internal/mapper/proto"
	"topup_game/internal/pb"
	"topup_game/internal/service"
	"topup_game/pkg/errors/auth_errors"
	refreshtoken_errors "topup_game/pkg/errors/refresh_token_errors"
)

type authHandleGrpc struct {
	pb.UnimplementedAuthServiceServer
	authService service.AuthService
	mapping     protomapper.AuthProtoMapper
}

func NewAuthHandleGrpc(auth service.AuthService, mapping protomapper.AuthProtoMapper) *authHandleGrpc {
	return &authHandleGrpc{authService: auth, mapping: mapping}
}

func (s *authHandleGrpc) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.ApiResponseLogin, error) {
	request := &requests.AuthRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := s.authService.Login(request)
	if err != nil {
		return nil, auth_errors.ErrGrpcLogin
	}

	return s.mapping.ToProtoResponseLogin("success", "Login successful", res), nil
}

func (s *authHandleGrpc) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.ApiResponseRefreshToken, error) {
	res, err := s.authService.RefreshToken(req.RefreshToken)

	if err != nil {
		return nil, refreshtoken_errors.ErrGrpcRefreshToken
	}

	return s.mapping.ToProtoResponseRefreshToken("success", "Registration successful", res), nil
}

func (s *authHandleGrpc) GetMe(ctx context.Context, req *pb.GetMeRequest) (*pb.ApiResponseGetMe, error) {
	res, err := s.authService.GetMe(req.AccessToken)

	if err != nil {
		return nil, auth_errors.ErrGrpcGetMe
	}

	return s.mapping.ToProtoResponseGetMe("success", "Refresh token successful", res), nil
}

func (s *authHandleGrpc) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.ApiResponseRegister, error) {
	request := &requests.CreateUserRequest{
		FirstName:       req.Firstname,
		LastName:        req.Lastname,
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	}

	res, errResp := s.authService.Register(request)
	if errResp != nil {
		return nil, auth_errors.ErrGrpcRegisterToken
	}

	return s.mapping.ToProtoResponseRegister("success", "Get me successful", res), nil
}
