package errors_custom

import (
	"encoding/json"
	"topup_game/internal/pb"
)

func GrpcErrorToJson(err *pb.ErrorResponse) string {
	jsonData, _ := json.Marshal(err)
	return string(jsonData)
}
