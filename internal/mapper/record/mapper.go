package recordmapper

type RecordMapper struct {
	UserRecordMapper         UserRecordMapping
	RoleRecordMapper         RoleRecordMapping
	UserRoleRecordMapper     UserRoleRecordMapping
	RefreshTokenRecordMapper RefreshTokenRecordMapping
	BankRecordMapper         BankRecordMapping
	CategoryRecordMapper     CategoryRecordMapping
	MerchantRecordMapper     MerchantRecordMapping
	NominalRecordMapper      NominalRecordMapping
	TransactionRecordMapper  TransactionRecordMapping
	VoucherRecordMapper      VoucherRecordMapping
}

func NewRecordMapper() *RecordMapper {
	return &RecordMapper{
		UserRecordMapper:         NewUserRecordMapper(),
		RoleRecordMapper:         NewRoleRecordMapper(),
		UserRoleRecordMapper:     NewUserRoleRecordMapper(),
		RefreshTokenRecordMapper: NewRefreshTokenRecordMapper(),
		BankRecordMapper:         NewBankRecordMapper(),
		CategoryRecordMapper:     NewCategoryRecordMapper(),
		MerchantRecordMapper:     NewMerchantRecordMapper(),
		NominalRecordMapper:      NewNominalRecordMapper(),
		TransactionRecordMapper:  NewTransactionRecordMapper(),
		VoucherRecordMapper:      NewVoucherRecordMapper(),
	}
}
