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

type IResponseApiMapper interface {
	GetAuthMapper() AuthResponseMapper
	GetRoleMapper() RoleResponseMapper
	GetUserMapper() UserResponseMapper
	GetBankMapper() BankResponseMapper
	GetCategoryMapper() CategoryResponseMapper
	GetMerchantMapper() MerchantResponseMapper
	GetNominalMapper() NominalResponseMapper
	GetTransactionMapper() TransactionResponseMapper
	GetVoucherMapper() VoucherResponseMapper
}

func NewResponseApiMapper() IResponseApiMapper {
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

func (r *ResponseApiMapper) GetAuthMapper() AuthResponseMapper {
	return r.AuthResponseMapper
}

func (r *ResponseApiMapper) GetBankMapper() BankResponseMapper {
	return r.BankResponseMapper
}

func (r *ResponseApiMapper) GetCategoryMapper() CategoryResponseMapper {
	return r.CategoryResponseMapper
}

func (r *ResponseApiMapper) GetMerchantMapper() MerchantResponseMapper {
	return r.MerchantResponseMapper
}

func (r *ResponseApiMapper) GetNominalMapper() NominalResponseMapper {
	return r.NominalResponseMapper
}

func (r *ResponseApiMapper) GetRoleMapper() RoleResponseMapper {
	return r.RoleResponseMapper
}

func (r *ResponseApiMapper) GetTransactionMapper() TransactionResponseMapper {
	return r.TransactionResponseMapper
}

func (r *ResponseApiMapper) GetUserMapper() UserResponseMapper {
	return r.UserResponseMapper
}

func (r *ResponseApiMapper) GetVoucherMapper() VoucherResponseMapper {
	return r.VoucherResponseMapper
}
