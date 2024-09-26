package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

type ProductService struct {
	productRepository      Interface.IProductRepository
	discountService        Interface.IDiscountService
	productDiscountService Interface.IProductDiscountService
}

func NewProductService(productRepository Interface.IProductRepository, discountService Interface.IDiscountService, ProductDiscountService Interface.IProductDiscountService) Interface.IProductService {
	return &ProductService{
		productRepository:      productRepository,
		discountService:        discountService,
		productDiscountService: ProductDiscountService,
	}
}

func (p *ProductService) GetPaginatedProductList(searchValue, sortBy, productID, sellerID, categoryID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Product], error) {
	return p.productRepository.GetPaginatedProductList(searchValue, sortBy, productID, sellerID, categoryID, pageIndex, pageSize, status)
}

func (p *ProductService) GetAllProducts() ([]BusinessObjects.Product, error) {
	return p.productRepository.GetAllProducts()
}

func (p *ProductService) GetProductByID(id string) (BusinessObjects.Product, error) {
	return p.productRepository.GetProductByID(id)
}

func (p *ProductService) CreateProduct(SellerID, ProductName, Description, CategoryID, ImageURL string, Price float64, Quantity int) error {
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

	err := p.productRepository.CreateProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductService) UpdateProduct(productID, SellerID, ProductName, Description, CategoryID, ImageURL string, Price float64, Quantity int) error {

	product, err := p.productRepository.GetProductByID(productID)
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

	return p.productRepository.UpdateProduct(product)
}

func (p *ProductService) DeleteProduct(id string) error {
	return p.productRepository.DeleteProduct(id)
}

func (p *ProductService) GetProductPriceAfterDiscount(productID string) (float64, error) {
	product, err := p.productRepository.GetProductByID(productID)
	if err != nil {
		return 0, err
	}

	discounts, err := p.productDiscountService.GetProductDiscountByID(productID)
	if err != nil {
		return 0, err
	}

	for _, discount := range discounts {
		discountEvent, err := p.discountService.GetDiscountByID(discount.DiscountID)
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
