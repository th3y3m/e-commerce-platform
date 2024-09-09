package BusinessObjects

import "time"

// VoucherRedemption represents the VoucherRedemptions table
type VoucherRedemption struct {
	RedemptionID string    `json:"redemption_id" db:"redemption_id"`
	VoucherID    string    `json:"voucher_id" db:"voucher_id"`
	UserID       string    `json:"user_id" db:"user_id"`
	OrderID      string    `json:"order_id" db:"order_id"`
	RedeemedAt   time.Time `json:"redeemed_at" db:"redeemed_at"`
}
