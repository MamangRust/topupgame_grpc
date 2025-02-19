package response_service

type ResponseServiceMapper struct {
	RoleResponseMapper         RoleResponseMapper
	RefreshTokenResponseMapper RefreshTokenResponseMapper
	UserResponseMapper         UserResponseMapper
	BankResponseMapper         BankResponseMapper
	CategoryResponseMapper     CategoryResponseMapper
	MerchantResponseMapper     MerchantResponseMapper
	NominalResponseMapper      NominalResponseMapper
	TransactionResponseMapper  TransactionResponseMapper
	VoucherResponseMapper      VoucherResponseMapper
}

func NewResponseServiceMapper() *ResponseServiceMapper {
	return &ResponseServiceMapper{
		UserResponseMapper:         NewUserResponseMapper(),
		RefreshTokenResponseMapper: NewRefreshTokenResponseMapper(),
		RoleResponseMapper:         NewRoleResponseMapper(),
		BankResponseMapper:         NewBankResponseMapper(),
		CategoryResponseMapper:     NewCategoryResponseMapper(),
		MerchantResponseMapper:     NewMerchantResponseMapper(),
		NominalResponseMapper:      NewNominalResponseMapper(),
		TransactionResponseMapper:  NewTransactionResponseMapper(),
		VoucherResponseMapper:      NewVoucherResponseMapper(),
	}
}
