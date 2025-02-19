package record

type VoucherRecord struct {
	ID         int     `json:"id"`
	MerchantID int     `json:"merchant_id"`
	CategoryID int     `json:"category_id"`
	Name       string  `json:"name"`
	ImageName  string  `json:"image_name"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	DeletedAt  *string `json:"deleted_at"`
}
