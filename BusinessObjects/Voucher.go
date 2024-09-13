package BusinessObjects

import "time"

// Voucher represents the Vouchers table
type Voucher struct {
	VoucherID          string    `gorm:"primaryKey;column:voucher_id"`
	VoucherCode        string    `gorm:"unique;not null;column:voucher_code"`
	DiscountType       string    `gorm:"column:discount_type"`
	DiscountValue      float64   `gorm:"column:discount_value"`
	MinimumOrderAmount float64   `gorm:"column:minimum_order_amount"`
	MaxDiscountAmount  float64   `gorm:"column:max_discount_amount"`
	StartDate          time.Time `gorm:"column:start_date"`
	EndDate            time.Time `gorm:"column:end_date"`
	UsageLimit         int       `gorm:"column:usage_limit"`
	UsageCount         int       `gorm:"column:usage_count"`
	Status             bool      `gorm:"column:status"`

	Orders []Order `gorm:"foreignKey:VoucherID"`
}
type NewVoucher struct {
	VoucherCode        string    `gorm:"unique;not null;column:voucher_code"`
	DiscountType       string    `gorm:"column:discount_type"`
	DiscountValue      float64   `gorm:"column:discount_value"`
	MinimumOrderAmount float64   `gorm:"column:minimum_order_amount"`
	MaxDiscountAmount  float64   `gorm:"column:max_discount_amount"`
	StartDate          time.Time `gorm:"column:start_date"`
	EndDate            time.Time `gorm:"column:end_date"`
	UsageLimit         int       `gorm:"column:usage_limit"`
	Status             bool      `gorm:"column:status"`
}
