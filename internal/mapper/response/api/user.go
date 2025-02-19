package response_api

import (
	"topup_game/internal/domain/response"
	"topup_game/internal/pb"
)

type userResponseMapper struct {
}

func NewUserResponseMapper() *userResponseMapper {
	return &userResponseMapper{}
}

func (u *userResponseMapper) ToResponseUser(user *pb.UserResponse) *response.UserResponse {
	return &response.UserResponse{
		ID:        int(user.Id),
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (u *userResponseMapper) ToResponsesUser(users []*pb.UserResponse) []*response.UserResponse {
	var mappedUsers []*response.UserResponse

	for _, user := range users {
		mappedUsers = append(mappedUsers, u.ToResponseUser(user))
	}

	return mappedUsers
}

func (u *userResponseMapper) ToResponseUserDeleteAt(user *pb.UserResponseDeleteAt) *response.UserResponseDeleteAt {
	return &response.UserResponseDeleteAt{
		ID:        int(user.Id),
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

func (u *userResponseMapper) ToResponsesUserDeleteAt(users []*pb.UserResponseDeleteAt) []*response.UserResponseDeleteAt {
	var mappedUsers []*response.UserResponseDeleteAt

	for _, user := range users {
		mappedUsers = append(mappedUsers, u.ToResponseUserDeleteAt(user))
	}

	return mappedUsers
}

func (u *userResponseMapper) ToApiResponseUser(pbResponse *pb.ApiResponseUser) *response.ApiResponseUser {
	return &response.ApiResponseUser{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    u.ToResponseUser(pbResponse.Data),
	}
}

func (u *userResponseMapper) ToApiResponseUserDeleteAt(pbResponse *pb.ApiResponseUserDeleteAt) *response.ApiResponseUserDeleteAt {
	return &response.ApiResponseUserDeleteAt{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    u.ToResponseUserDeleteAt(pbResponse.Data),
	}
}

func (u *userResponseMapper) ToApiResponsesUser(pbResponse *pb.ApiResponsesUser) *response.ApiResponsesUser {
	return &response.ApiResponsesUser{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    u.ToResponsesUser(pbResponse.Data),
	}
}

func (u *userResponseMapper) ToApiResponseUserDelete(pbResponse *pb.ApiResponseUserDelete) *response.ApiResponseUserDelete {
	return &response.ApiResponseUserDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (u *userResponseMapper) ToApiResponseUserAll(pbResponse *pb.ApiResponseUserAll) *response.ApiResponseUserAll {
	return &response.ApiResponseUserAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (u *userResponseMapper) ToApiResponsePaginationUserDeleteAt(pbResponse *pb.ApiResponsePaginationUserDeleteAt) *response.ApiResponsePaginationUserDeleteAt {
	return &response.ApiResponsePaginationUserDeleteAt{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       u.ToResponsesUserDeleteAt(pbResponse.Data),
		Pagination: *mapPaginationMeta(pbResponse.Pagination),
	}
}

func (u *userResponseMapper) ToApiResponsePaginationUser(pbResponse *pb.ApiResponsePaginationUser) *response.ApiResponsePaginationUser {
	return &response.ApiResponsePaginationUser{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       u.ToResponsesUser(pbResponse.Data),
		Pagination: *mapPaginationMeta(pbResponse.Pagination),
	}
}
