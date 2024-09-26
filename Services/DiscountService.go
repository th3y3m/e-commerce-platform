package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

type DiscountService struct {
	discountRepository Interface.IDiscountRepository
}

func NewDiscountService(discountRepository Interface.IDiscountRepository) Interface.IDiscountService {
	return &DiscountService{discountRepository: discountRepository}
}

func (d *DiscountService) GetPaginatedDiscountList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Discount], error) {
	return d.discountRepository.GetPaginatedDiscountList(searchValue, sortBy, pageIndex, pageSize, status)
}

func (d *DiscountService) GetAllDiscounts() ([]BusinessObjects.Discount, error) {
	return d.discountRepository.GetAllDiscounts()
}

func (d *DiscountService) GetDiscountByID(id string) (BusinessObjects.Discount, error) {
	return d.discountRepository.GetDiscountByID(id)
}

func (d *DiscountService) CreateDiscount(DiscountType string, DiscountValue float64, startDate, endDate time.Time) error {
	discount := BusinessObjects.Discount{
		DiscountID:    "DISC" + Util.GenerateID(10),
		DiscountType:  DiscountType,
		DiscountValue: DiscountValue,
		StartDate:     startDate,
		EndDate:       endDate,
	}

	err := d.discountRepository.CreateDiscount(discount)
	if err != nil {
		return err
	}

	return nil
}

func (d *DiscountService) UpdateDiscount(discountID, discountType string, discountValue float64, startDate, endDate time.Time) error {

	discount, err := d.discountRepository.GetDiscountByID(discountID)
	if err != nil {
		return err
	}

	discount.DiscountType = discountType
	discount.DiscountValue = discountValue
	discount.StartDate = startDate
	discount.EndDate = endDate

	return d.discountRepository.UpdateDiscount(discount)
}

func (d *DiscountService) DeleteDiscount(id string) error {
	return d.discountRepository.DeleteDiscount(id)
}
