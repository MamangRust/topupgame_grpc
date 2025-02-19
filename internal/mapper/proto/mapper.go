package protomapper

type ProtoMapper struct {
	AuthProtoMapper        AuthProtoMapper
	RoleProtoMapper        RoleProtoMapper
	UserProtoMapper        UserProtoMapper
	BankProtoMapper        BankProtoMapper
	CategoryProtoMapper    CategoryProtoMapper
	MerchantProtoMapper    MerchantProtoMapper
	NominalProtomapper     NominalProtoMapper
	TransactionProtoMapper TransactionProtoMapper
	VoucherProtoMapper     VoucherProtoMapper
}

func NewProtoMapper() *ProtoMapper {
	return &ProtoMapper{
		AuthProtoMapper:        NewAuthProtoMapper(),
		RoleProtoMapper:        NewRoleProtoMapper(),
		UserProtoMapper:        NewUserProtoMapper(),
		BankProtoMapper:        NewBankProtoMapper(),
		CategoryProtoMapper:    NewCategoryProtoMapper(),
		MerchantProtoMapper:    NewMerchantProtoMaper(),
		NominalProtomapper:     NewNominalProtoMapper(),
		TransactionProtoMapper: NewTransactionProtoMapper(),
		VoucherProtoMapper:     NewVoucherProtoMapper(),
	}
}
