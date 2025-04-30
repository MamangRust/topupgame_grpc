package service

import (
	"database/sql"
	"errors"
	"strconv"
	"time"
	"topup_game/internal/domain/requests"
	"topup_game/internal/domain/response"
	response_service "topup_game/internal/mapper/response/service"
	"topup_game/internal/repository"
	"topup_game/pkg/auth"
	refreshtoken_errors "topup_game/pkg/errors/refresh_token_errors"
	"topup_game/pkg/errors/role_errors"
	"topup_game/pkg/errors/user_errors"
	userrole_errors "topup_game/pkg/errors/user_role_errors"
	"topup_game/pkg/hash"
	"topup_game/pkg/logger"

	"go.uber.org/zap"
)

type authService struct {
	auth         repository.UserRepository
	refreshToken repository.RefreshTokenRepository
	userRole     repository.UserRoleRepository
	role         repository.RoleRepository
	hash         hash.HashPassword
	token        auth.TokenManager
	logger       logger.LoggerInterface
	mapping      response_service.UserResponseMapper
}

func NewAuthService(auth repository.UserRepository, refreshToken repository.RefreshTokenRepository, role repository.RoleRepository, userRole repository.UserRoleRepository, hash hash.HashPassword, token auth.TokenManager, logger logger.LoggerInterface, mapping response_service.UserResponseMapper) *authService {
	return &authService{auth: auth, refreshToken: refreshToken, role: role, userRole: userRole, hash: hash, token: token, logger: logger, mapping: mapping}
}

func (s *authService) Register(request *requests.CreateUserRequest) (*response.UserResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting user registration",
		zap.String("email", request.Email),
		zap.String("firstname", request.FirstName),
		zap.String("lastname", request.LastName),
	)

	_, err := s.auth.FindByEmail(request.Email)

	if err == nil {
		s.logger.Error("Email already exists", zap.String("email", request.Email))
		return nil, user_errors.ErrUserEmailAlready
	}

	passwordHash, err := s.hash.HashPassword(request.Password)

	if err != nil {
		s.logger.Error("Failed to hash password", zap.Error(err))
		return nil, user_errors.ErrUserPassword
	}
	request.Password = passwordHash

	_, err = s.role.FindByName("Cashier")

	if err != nil {
		s.logger.Error("Failed to find role", zap.Error(err))
		return nil, role_errors.ErrRoleNotFoundRes
	}

	res, err := s.auth.CreateUser(request)
	if err != nil {
		s.logger.Error("Failed to create user", zap.Error(err))
		return nil, user_errors.ErrFailedCreateUser
	}

	_, err = s.userRole.AssignRoleToUser(&requests.CreateUserRoleRequest{
		UserId: res.ID,
		RoleId: 1,
	})

	if err != nil {
		s.logger.Error("Failed to assign role to user", zap.Error(err))
		return nil, userrole_errors.ErrFailedAssignRoleToUser
	}

	s.logger.Debug("User registered successfully", zap.String("email", request.Email))

	so := s.mapping.ToUserResponse(res)

	return so, nil
}

func (s *authService) Login(request *requests.AuthRequest) (*response.TokenResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting login process",
		zap.String("email", request.Email),
	)

	res, err := s.auth.FindByEmail(request.Email)

	if err != nil {
		s.logger.Error("Failed to get user", zap.Error(err))
		return nil, user_errors.ErrUserNotFoundRes
	}

	err = s.hash.ComparePassword(res.Password, request.Password)

	if err != nil {
		s.logger.Error("Failed to compare password", zap.Error(err))
		return nil, user_errors.ErrUserPassword
	}

	token, err := s.createAccessToken(res.ID)

	if err != nil {
		s.logger.Error("Failed to generate JWT token", zap.Error(err))
		return nil, refreshtoken_errors.ErrFailedCreateAccess
	}

	refreshToken, err := s.createRefreshToken(res.ID)

	if err != nil {
		s.logger.Error("Failed to generate refresh token", zap.Error(err))
		return nil, refreshtoken_errors.ErrFailedCreateRefresh
	}

	s.logger.Debug("User logged in successfully", zap.String("email", request.Email))

	return &response.TokenResponse{AccessToken: token, RefreshToken: refreshToken}, nil
}

func (s *authService) RefreshToken(token string) (*response.TokenResponse, *response.ErrorResponse) {
	s.logger.Debug("Refreshing token",
		zap.String("token", token),
	)

	userIdStr, err := s.token.ValidateToken(token)

	if err != nil {
		if errors.Is(err, auth.ErrTokenExpired) {
			if err := s.refreshToken.DeleteRefreshToken(token); err != nil {
				s.logger.Error("Failed to delete expired refresh token", zap.Error(err))

				return nil, refreshtoken_errors.ErrFailedDeleteRefreshToken
			}

			s.logger.Error("Refresh token has expired", zap.Error(err))

			return nil, refreshtoken_errors.ErrFailedExpire
		}
		s.logger.Error("Invalid refresh token", zap.Error(err))
		return nil, refreshtoken_errors.ErrRefreshTokenNotFound
	}

	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		s.logger.Error("Invalid user ID format in token", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Invalid user ID format in token",
		}
	}

	accessToken, err := s.createAccessToken(userId)
	if err != nil {
		s.logger.Error("Failed to generate new access token", zap.Error(err))

		return nil, refreshtoken_errors.ErrFailedCreateAccess
	}

	refreshToken, err := s.createRefreshToken(userId)
	if err != nil {
		s.logger.Error("Failed to generate new refresh token", zap.Error(err))

		return nil, refreshtoken_errors.ErrFailedCreateRefreshToken
	}

	expiryTime := time.Now().Add(24 * time.Hour)

	updateRequest := &requests.UpdateRefreshToken{
		UserId:    userId,
		Token:     refreshToken,
		ExpiresAt: expiryTime.Format("2006-01-02 15:04:05"),
	}

	if _, err = s.refreshToken.UpdateRefreshToken(updateRequest); err != nil {
		s.logger.Error("Failed to update refresh token in storage", zap.Error(err))

		return nil, refreshtoken_errors.ErrFailedUpdateRefreshToken
	}

	s.logger.Debug("Refresh token refreshed successfully")

	return &response.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) GetMe(token string) (*response.UserResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching user details",
		zap.String("token", token),
	)

	userIdStr, err := s.token.ValidateToken(token)

	if err != nil {
		s.logger.Error("Invalid access token", zap.Error(err))
		return nil, refreshtoken_errors.ErrFailedInValidToken
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		s.logger.Error("Invalid user ID format in token", zap.Error(err))
		return nil, refreshtoken_errors.ErrFailedInValidUserId
	}

	user, err := s.auth.FindById(userId)

	if err != nil {
		s.logger.Error("Failed to find user by ID", zap.Error(err))
		return nil, user_errors.ErrUserNotFoundRes
	}

	so := s.mapping.ToUserResponse(user)

	s.logger.Debug("User details fetched successfully",
		zap.Int("userID", userId),
	)

	return so, nil
}

func (s *authService) createAccessToken(id int) (string, error) {
	s.logger.Debug("Creating access token",
		zap.Int("userID", id),
	)

	res, err := s.token.GenerateToken(id, "access")

	if err != nil {
		s.logger.Error("Failed to create access token",
			zap.Int("userID", id),
			zap.Error(err))
		return "", err
	}

	s.logger.Debug("Access token created successfully",
		zap.Int("userID", id),
	)

	return res, nil
}

func (s *authService) createRefreshToken(id int) (string, error) {
	s.logger.Debug("Creating refresh token",
		zap.Int("userID", id),
	)

	res, err := s.token.GenerateToken(id, "refresh")

	if err != nil {
		s.logger.Error("Failed to create refresh token",
			zap.Int("userID", id),
			zap.Error(err),
		)

		return "", err
	}

	if err := s.refreshToken.DeleteRefreshTokenByUserId(id); err != nil && !errors.Is(err, sql.ErrNoRows) {
		s.logger.Error("Failed to delete existing refresh token", zap.Error(err))
		return "", err
	}

	_, err = s.refreshToken.CreateRefreshToken(&requests.CreateRefreshToken{Token: res, UserId: id, ExpiresAt: time.Now().Add(24 * time.Hour).Format("2006-01-02 15:04:05")})
	if err != nil {
		s.logger.Error("Failed to create refresh token", zap.Error(err))

		return "", err
	}

	s.logger.Debug("Refresh token created successfully",
		zap.Int("userID", id),
	)

	return res, nil
}
