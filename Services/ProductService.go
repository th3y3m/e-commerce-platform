package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

func GetPaginatedProductList(searchValue, sortBy, productID, sellerID, categoryID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Product], error) {
	return Repositories.GetPaginatedProductList(searchValue, sortBy, productID, sellerID, categoryID, pageIndex, pageSize, status)
}

func GetAllProducts() ([]BusinessObjects.Product, error) {
	return Repositories.GetAllProducts()
}

func GetProductByID(id string) (BusinessObjects.Product, error) {
	return Repositories.GetProductByID(id)
}

func CreateProduct(SellerID, ProductName, Description, CategoryID, ImageURL string, Price float64, Quantity int) error {
	product := BusinessObjects.Product{
		ProductID:   "PROD" + Util.GenerateID(10),
		SellerID:    SellerID,
		ProductName: ProductName,
		Description: Description,
		CategoryID:  CategoryID,
		Price:       Price,
		ImageURL:    ImageURL,
		Quantity:    Quantity,
		Status:      true,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}

	err := Repositories.CreateProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProduct(productID, SellerID, ProductName, Description, CategoryID, ImageURL string, Price float64, Quantity int) error {

	product, err := Repositories.GetProductByID(productID)
	if err != nil {
		return err
	}

	product.SellerID = SellerID
	product.ProductName = ProductName
	product.Description = Description
	product.CategoryID = CategoryID
	product.Price = Price
	product.ImageURL = ImageURL
	product.Quantity = Quantity
	product.UpdatedAt = time.Now()

	return Repositories.UpdateProduct(product)
}

func DeleteProduct(id string) error {
	return Repositories.DeleteProduct(id)
}

func GetProductPriceAfterDiscount(productID string) (float64, error) {
	product, err := Repositories.GetProductByID(productID)
	if err != nil {
		return 0, err
	}

	discounts, err := GetProductDiscountByID(productID)
	if err != nil {
		return 0, err
	}

	for _, discount := range discounts {
		discountEvent, err := GetDiscountByID(discount.DiscountID)
		if err != nil {
			return 0, err
		}

		if discountEvent.DiscountType == "Percentage" {
			product.Price -= product.Price * discountEvent.DiscountValue
		} else {
			product.Price -= discountEvent.DiscountValue

		}
	}

	return product.Price, nil
}
