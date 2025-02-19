package recordmapper

import (
	"topup_game/internal/domain/record"
	db "topup_game/pkg/database/schema"
)

type refreshTokenRecordMapper struct {
}

func NewRefreshTokenRecordMapper() *refreshTokenRecordMapper {
	return &refreshTokenRecordMapper{}
}

func (r *refreshTokenRecordMapper) ToRefreshTokenRecord(refreshToken *db.RefreshToken) *record.RefreshTokenRecord {
	return &record.RefreshTokenRecord{
		ID:        int(refreshToken.RefreshTokenID),
		UserID:    int(refreshToken.UserID),
		Token:     refreshToken.Token,
		ExpiredAt: refreshToken.Expiration.String(),
		CreatedAt: refreshToken.CreatedAt.Time.String(),
		UpdatedAt: refreshToken.UpdatedAt.Time.String(),
	}
}

func (r *refreshTokenRecordMapper) ToRefreshTokensRecord(refreshTokens []*db.RefreshToken) []*record.RefreshTokenRecord {
	var refreshTokenRecords []*record.RefreshTokenRecord
	for _, refreshToken := range refreshTokens {
		refreshTokenRecords = append(refreshTokenRecords, r.ToRefreshTokenRecord(refreshToken))
	}
	return refreshTokenRecords
}
