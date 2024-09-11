package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetPaginatedShoppingCartList(sortBy, cartID, userID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.ShoppingCart], error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return Util.PaginatedList[BusinessObjects.ShoppingCart]{}, err
	}

	var carts []BusinessObjects.ShoppingCart
	query := db.Model(&BusinessObjects.ShoppingCart{})

	if cartID != "" {
		query = query.Where("cart_id = ?", cartID)
	}

	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}

	if status != nil {
		query = query.Where("status = ?", *status)
	}

	switch sortBy {
	case "cart_id_asc":
		query = query.Order("cart_id ASC")
	case "cart_id_desc":
		query = query.Order("cart_id DESC")
	case "user_id_asc":
		query = query.Order("user_id ASC")
	case "user_id_desc":
		query = query.Order("user_id DESC")
	case "status_asc":
		query = query.Order("status ASC")
	case "status_desc":
		query = query.Order("status DESC")
	case "created_at_asc":
		query = query.Order("created_at ASC")
	case "created_at_desc":
		query = query.Order("created_at DESC")
	default:
		query = query.Order("created_at DESC")
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.ShoppingCart]{}, err
	}

	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&carts).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.ShoppingCart]{}, err
	}

	return Util.PaginatedList[BusinessObjects.ShoppingCart]{Items: carts, TotalCount: total, PageIndex: pageIndex, PageSize: pageSize}, nil
}

// GetAllShoppingCarts retrieves all shopping carts from the database
func GetAllShoppingCarts() ([]BusinessObjects.ShoppingCart, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var carts []BusinessObjects.ShoppingCart
	if err := db.Find(&carts).Error; err != nil {
		return nil, err
	}

	return carts, nil
}

// GetShoppingCartByID retrieves a shopping cart by its ID
func GetShoppingCartByID(cartID string) (BusinessObjects.ShoppingCart, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return BusinessObjects.ShoppingCart{}, err
	}

	var cart BusinessObjects.ShoppingCart
	if err := db.First(&cart, "cart_id = ?", cartID).Error; err != nil {
		return BusinessObjects.ShoppingCart{}, err
	}

	return cart, nil
}

// CreateShoppingCart adds a new shopping cart to the database
func CreateShoppingCart(cart BusinessObjects.ShoppingCart) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Create(&cart).Error; err != nil {
		return err
	}

	return nil
}

// UpdateShoppingCart updates an existing shopping cart
func UpdateShoppingCart(cart BusinessObjects.ShoppingCart) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Save(&cart).Error; err != nil {
		return err
	}

	return nil
}

// DeleteShoppingCart removes a shopping cart from the database by its ID
func DeleteShoppingCart(cartID string) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Delete(&BusinessObjects.ShoppingCart{}, "cart_id = ?", cartID).Error; err != nil {
		return err
	}

	return nil
}
