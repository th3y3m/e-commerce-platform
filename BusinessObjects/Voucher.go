package BusinessObjects

import "time"

// Voucher represents the Vouchers table
type Voucher struct {
	VoucherID          string    `json:"voucher_id" db:"voucher_id"`
	VoucherCode        string    `json:"voucher_code" db:"voucher_code"`
	DiscountType       string    `json:"discount_type" db:"discount_type"`
	DiscountValue      float64   `json:"discount_value" db:"discount_value"`
	MinimumOrderAmount float64   `json:"minimum_order_amount" db:"minimum_order_amount"`
	MaxDiscountAmount  float64   `json:"max_discount_amount" db:"max_discount_amount"`
	StartDate          time.Time `json:"start_date" db:"start_date"`
	EndDate            time.Time `json:"end_date" db:"end_date"`
	UsageLimit         int       `json:"usage_limit" db:"usage_limit"`
	UsageCount         int       `json:"usage_count" db:"usage_count"`
	IsActive           bool      `json:"is_active" db:"is_active"`
}
