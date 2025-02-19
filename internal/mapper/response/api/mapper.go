package response_api

type ResponseApiMapper struct {
	AuthResponseMapper        AuthResponseMapper
	RoleResponseMapper        RoleResponseMapper
	UserResponseMapper        UserResponseMapper
	BankResponseMapper        BankResponseMapper
	CategoryResponseMapper    CategoryResponseMapper
	MerchantResponseMapper    MerchantResponseMapper
	NominalResponseMapper     NominalResponseMapper
	TransactionResponseMapper TransactionResponseMapper
	VoucherResponseMapper     VoucherResponseMapper
}

func NewResponseApiMapper() *ResponseApiMapper {
	return &ResponseApiMapper{
		AuthResponseMapper:        NewAuthResponseMapper(),
		UserResponseMapper:        NewUserResponseMapper(),
		RoleResponseMapper:        NewRoleResponseMapper(),
		BankResponseMapper:        NewBankResponseMapper(),
		CategoryResponseMapper:    NewCategoryResponseMapper(),
		MerchantResponseMapper:    NewMerchantResponseMapper(),
		NominalResponseMapper:     NewNominalResponseMapper(),
		TransactionResponseMapper: NewTransactionResponseMapper(),
		VoucherResponseMapper:     NewVoucherResponseMapper(),
	}
}
