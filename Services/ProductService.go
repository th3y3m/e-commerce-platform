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

func CreateProduct(SellerID, ProductName, Description, CategoryID string, Price float64, Quantity int) error {
	product := BusinessObjects.Product{
		ProductID:   "PROD" + Util.GenerateID(10),
		SellerID:    SellerID,
		ProductName: ProductName,
		Description: Description,
		CategoryID:  CategoryID,
		Price:       Price,
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

func UpdateProduct(productID, SellerID, ProductName, Description, CategoryID string, Price float64, Quantity int) error {

	product, err := Repositories.GetProductByID(productID)
	if err != nil {
		return err
	}

	product.SellerID = SellerID
	product.ProductName = ProductName
	product.Description = Description
	product.CategoryID = CategoryID
	product.Price = Price
	product.Quantity = Quantity
	product.UpdatedAt = time.Now()

	return Repositories.UpdateProduct(product)
}

func DeleteProduct(id string) error {
	return Repositories.DeleteProduct(id)
}
