package gapi

import "topup_game/internal/pb"

type AuthHandleGrpc interface {
	pb.AuthServiceServer
}

type RoleHandleGrpc interface {
	pb.RoleServiceServer
}

type UserHandleGrpc interface {
	pb.UserServiceServer
}

type BankHandleGrpc interface {
	pb.BankServiceServer
}

type CategoryHandleGrpc interface {
	pb.CategoryServiceServer
}

type MerchantHandleGrpc interface {
	pb.MerchantServiceServer
}

type NominalHandleGrpc interface {
	pb.NominalServiceServer
}

type TransactionHandleGrpc interface {
	pb.TransactionServiceServer
}

type VoucherHandleGrpc interface {
	pb.VoucherServiceServer
}
