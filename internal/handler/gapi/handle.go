package gapi

import (
	protomapper "topup_game/internal/mapper/proto"
	"topup_game/internal/service"
)

type Deps struct {
	Service service.Service
	Mapper  protomapper.ProtoMapper
}

type Handler struct {
	Auth        AuthHandleGrpc
	Role        RoleHandleGrpc
	User        UserHandleGrpc
	Bank        BankHandleGrpc
	Category    CategoryHandleGrpc
	Merchant    MerchantHandleGrpc
	Nominal     NominalHandleGrpc
	Transaction TransactionHandleGrpc
	Voucher     VoucherHandleGrpc
}

func NewHandler(deps Deps) *Handler {
	return &Handler{
		Auth:        NewAuthHandleGrpc(deps.Service.Auth, deps.Mapper.AuthProtoMapper),
		Role:        NewRoleHandleGrpc(deps.Service.Role, deps.Mapper.RoleProtoMapper),
		User:        NewUserHandleGrpc(deps.Service.User, deps.Mapper.UserProtoMapper),
		Bank:        NewBankHandleGrpc(deps.Service.Bank, deps.Mapper.BankProtoMapper),
		Category:    NewCategoryHandleGrpc(deps.Service.Category, deps.Mapper.CategoryProtoMapper),
		Merchant:    NewMerchantHandleGrpc(deps.Service.Merchant, deps.Mapper.MerchantProtoMapper),
		Nominal:     NewNominalHandleGrpc(deps.Service.Nominal, deps.Mapper.NominalProtomapper),
		Transaction: NewTransactionHandleGrpc(deps.Service.Transaction, deps.Mapper.TransactionProtoMapper),
		Voucher:     NewVoucherHandleGrpc(deps.Service.Voucher, deps.Mapper.VoucherProtoMapper),
	}
}
