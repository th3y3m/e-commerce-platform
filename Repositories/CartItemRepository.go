package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetPaginatedCartItemList(searchValue, sortBy, cartId, productId string, pageIndex, pageSize int) (Util.PaginatedList[BusinessObjects.CartItem], error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return Util.PaginatedList[BusinessObjects.CartItem]{}, err
	}

	var cartItems []BusinessObjects.CartItem
	query := db.Model(&BusinessObjects.CartItem{})

	// Apply search filter
	if searchValue != "" {
		query = query.Where("LOWER(cart_id) LIKE ?", "%"+searchValue+"%")
	}

	// Apply sorting
	switch sortBy {
	case "cart_id_asc":
		query = query.Order("cart_id ASC")
	case "cart_id_desc":
		query = query.Order("cart_id DESC")
	case "product_id_asc":
		query = query.Order("product_id ASC")
	case "product_id_desc":
		query = query.Order("product_id DESC")
	case "quantity_asc":
		query = query.Order("quantity ASC")
	case "quantity_desc":
		query = query.Order("quantity DESC")
	case "price_asc":
		query = query.Order("price ASC")
	case "price_desc":
		query = query.Order("price DESC")
	default:
		query = query.Order("product_id ASC")
	}

	// Get the total count of records
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.CartItem]{}, err
	}

	// Paginate the results
	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&cartItems).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.CartItem]{}, err
	}

	return Util.NewPaginatedList(cartItems, total, pageIndex, pageSize), nil
}

// GetAllCartItems retrieves all freight cartItems from the database
func GetAllCartItems() ([]BusinessObjects.CartItem, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var cartItem []BusinessObjects.CartItem
	if err := db.Find(&cartItem).Error; err != nil {
		return nil, err
	}

	return cartItem, nil
}

// GetCartItemByID retrieves a freight cartItem by its ID
func GetCartItemByID(cartItemID string) ([]BusinessObjects.CartItem, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var cartItem []BusinessObjects.CartItem
	if err := db.Find(&cartItem, "cart_id = ?", cartItemID).Error; err != nil {
		return nil, err
	}

	return cartItem, nil
}

// CreateCartItem adds a new freight cartItem to the database
func CreateCartItem(cartItem BusinessObjects.CartItem) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Create(&cartItem).Error; err != nil {
		return err
	}

	return nil
}

// UpdateCartItem updates an existing freight cartItem
func UpdateCartItem(cartItem BusinessObjects.CartItem) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Save(&cartItem).Error; err != nil {
		return err
	}

	return nil
}

// DeleteCartItem removes a freight cartItem from the database by its ID
func DeleteCartItem(cartItemID string) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Delete(&BusinessObjects.CartItem{}, "cart_id = ?", cartItemID).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByToken(token string) (BusinessObjects.User, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return BusinessObjects.User{}, err
	}

	var user BusinessObjects.User
	if err := db.First(&user, "token = ?", token).Error; err != nil {
		return BusinessObjects.User{}, err
	}

	return user, nil
}
