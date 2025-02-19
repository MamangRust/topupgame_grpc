package api

import (
	response_api "topup_game/internal/mapper/response/api"
	"topup_game/internal/pb"
	"topup_game/pkg/auth"
	"topup_game/pkg/logger"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type Deps struct {
	Conn    *grpc.ClientConn
	Token   auth.TokenManager
	E       *echo.Echo
	Logger  logger.LoggerInterface
	Mapping response_api.ResponseApiMapper
}

func NewHandler(deps Deps) {

	clientAuth := pb.NewAuthServiceClient(deps.Conn)
	clientRole := pb.NewRoleServiceClient(deps.Conn)
	clientUser := pb.NewUserServiceClient(deps.Conn)
	clientBank := pb.NewBankServiceClient(deps.Conn)
	clientCategory := pb.NewCategoryServiceClient(deps.Conn)
	clientMerchant := pb.NewMerchantServiceClient(deps.Conn)
	clientNominal := pb.NewNominalServiceClient(deps.Conn)
	clientTransaction := pb.NewTransactionServiceClient(deps.Conn)
	clientVoucher := pb.NewVoucherServiceClient(deps.Conn)

	NewHandlerAuth(deps.E, clientAuth, deps.Logger, deps.Mapping.AuthResponseMapper)
	NewHandlerRole(deps.E, clientRole, deps.Logger, deps.Mapping.RoleResponseMapper)
	NewHandlerUser(deps.E, clientUser, deps.Logger, deps.Mapping.UserResponseMapper)
	NewHandlerBank(deps.E, clientBank, deps.Logger, deps.Mapping.BankResponseMapper)
	NewHandlerCategory(deps.E, clientCategory, deps.Logger, deps.Mapping.CategoryResponseMapper)
	NewHandlerMerchant(deps.E, clientMerchant, deps.Logger, deps.Mapping.MerchantResponseMapper)
	NewHandlerNominal(deps.E, clientNominal, deps.Logger, deps.Mapping.NominalResponseMapper)
	NewHandlerTransaction(deps.E, clientTransaction, deps.Logger, deps.Mapping.TransactionResponseMapper)
	NewHandlerVoucher(deps.E, clientVoucher, deps.Logger, deps.Mapping.VoucherResponseMapper)
}
