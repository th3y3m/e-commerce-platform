package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

func GetPaginatedVoucherList(sortBy string, voucherID string, pageIndex int, pageSize int, status *bool, startDate time.Time, endDate time.Time) (Util.PaginatedList[BusinessObjects.Voucher], error) {
	return Repositories.GetPaginatedVoucherList(sortBy, voucherID, pageIndex, pageSize, status, startDate, endDate)
}

func GetAllVouchers() ([]BusinessObjects.Voucher, error) {
	return Repositories.GetAllVouchers()
}

func GetVoucherByID(id string) (BusinessObjects.Voucher, error) {
	return Repositories.GetVoucherByID(id)
}

func CreateVoucher(vou BusinessObjects.NewVoucher) error {
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

	err := Repositories.CreateVoucher(voucher)
	if err != nil {
		return err
	}

	return nil
}

func UpdateVoucher(vou BusinessObjects.Voucher) error {
	return Repositories.UpdateVoucher(vou)
}

func DeleteVoucher(id string) error {
	return Repositories.DeleteVoucher(id)
}
