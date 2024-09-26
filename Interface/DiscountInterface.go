package Interface

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

type IDiscountRepository interface {
	GetPaginatedDiscountList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Discount], error)
	GetAllDiscounts() ([]BusinessObjects.Discount, error)
	GetDiscountByID(discountID string) (BusinessObjects.Discount, error)
	CreateDiscount(discount BusinessObjects.Discount) error
	UpdateDiscount(discount BusinessObjects.Discount) error
	DeleteDiscount(discountID string) error
}

type IDiscountService interface {
	GetPaginatedDiscountList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Discount], error)
	GetAllDiscounts() ([]BusinessObjects.Discount, error)
	GetDiscountByID(discountID string) (BusinessObjects.Discount, error)
	CreateDiscount(DiscountType string, DiscountValue float64, startDate, endDate time.Time) error
	UpdateDiscount(discountID, discountType string, discountValue float64, startDate, endDate time.Time) error
	DeleteDiscount(discountID string) error
}
