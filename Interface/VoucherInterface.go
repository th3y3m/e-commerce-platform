package Interface

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

type IVoucherRepository interface {
	GetPaginatedVoucherList(sortBy, voucherID string, pageIndex, pageSize int, status *bool, startDate, endDate time.Time) (Util.PaginatedList[BusinessObjects.Voucher], error)
	GetAllVouchers() ([]BusinessObjects.Voucher, error)
	GetVoucherByID(voucherID string) (BusinessObjects.Voucher, error)
	CreateVoucher(voucher BusinessObjects.Voucher) error
	UpdateVoucher(voucher BusinessObjects.Voucher) error
	DeleteVoucher(voucherID string) error
}
type IVoucherService interface {
	GetPaginatedVoucherList(sortBy, voucherID string, pageIndex, pageSize int, status *bool, startDate, endDate time.Time) (Util.PaginatedList[BusinessObjects.Voucher], error)
	GetAllVouchers() ([]BusinessObjects.Voucher, error)
	GetVoucherByID(voucherID string) (BusinessObjects.Voucher, error)
	CreateVoucher(vou BusinessObjects.NewVoucher) error
	UpdateVoucher(voucher BusinessObjects.Voucher) error
	DeleteVoucher(voucherID string) error
}
