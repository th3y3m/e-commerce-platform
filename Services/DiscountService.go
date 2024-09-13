package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

func GetPaginatedDiscountList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Discount], error) {
	return Repositories.GetPaginatedDiscountList(searchValue, sortBy, pageIndex, pageSize, status)
}

func GetAllDiscounts() ([]BusinessObjects.Discount, error) {
	return Repositories.GetAllDiscounts()
}

func GetDiscountByID(id string) (BusinessObjects.Discount, error) {
	return Repositories.GetDiscountByID(id)
}

func CreateDiscount(DiscountType string, DiscountValue float64, startDate, endDate time.Time) error {
	discount := BusinessObjects.Discount{
		DiscountID:    "DISC" + Util.GenerateID(10),
		DiscountType:  DiscountType,
		DiscountValue: DiscountValue,
		StartDate:     startDate,
		EndDate:       endDate,
	}

	err := Repositories.CreateDiscount(discount)
	if err != nil {
		return err
	}

	return nil
}

func UpdateDiscount(discountID, discountType string, discountValue float64, startDate, endDate time.Time) error {

	discount, err := Repositories.GetDiscountByID(discountID)
	if err != nil {
		return err
	}

	discount.DiscountType = discountType
	discount.DiscountValue = discountValue
	discount.StartDate = startDate
	discount.EndDate = endDate

	return Repositories.UpdateDiscount(discount)
}

func DeleteDiscount(id string) error {
	return Repositories.DeleteDiscount(id)
}
