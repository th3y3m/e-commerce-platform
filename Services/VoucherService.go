package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

type VoucherService struct {
	voucherRepository Interface.IVoucherRepository
}

func NewVoucherService(voucherRepository Interface.IVoucherRepository) Interface.IVoucherService {
	return &VoucherService{voucherRepository}
}

func (v *VoucherService) GetPaginatedVoucherList(sortBy string, voucherID string, pageIndex int, pageSize int, status *bool, startDate time.Time, endDate time.Time) (Util.PaginatedList[BusinessObjects.Voucher], error) {
	return v.voucherRepository.GetPaginatedVoucherList(sortBy, voucherID, pageIndex, pageSize, status, startDate, endDate)
}

func (v *VoucherService) GetAllVouchers() ([]BusinessObjects.Voucher, error) {
	return v.voucherRepository.GetAllVouchers()
}

func (v *VoucherService) GetVoucherByID(id string) (BusinessObjects.Voucher, error) {
	return v.voucherRepository.GetVoucherByID(id)
}

func (v *VoucherService) CreateVoucher(vou BusinessObjects.NewVoucher) error {
	voucher := BusinessObjects.Voucher{
		VoucherID:          "VOU" + Util.GenerateID(10),
		DiscountType:       vou.DiscountType,
		DiscountValue:      vou.DiscountValue,
		StartDate:          vou.StartDate,
		EndDate:            vou.EndDate,
		Status:             true,
		MinimumOrderAmount: vou.MinimumOrderAmount,
		MaxDiscountAmount:  vou.MaxDiscountAmount,
		UsageLimit:         vou.UsageLimit,
		UsageCount:         0,
	}

	err := v.voucherRepository.CreateVoucher(voucher)
	if err != nil {
		return err
	}

	return nil
}

func (v *VoucherService) UpdateVoucher(vou BusinessObjects.Voucher) error {
	return v.voucherRepository.UpdateVoucher(vou)
}

func (v *VoucherService) DeleteVoucher(id string) error {
	return v.voucherRepository.DeleteVoucher(id)
}
