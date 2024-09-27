package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CartItemRepository struct {
	log *logrus.Logger
}

func NewCartItemRepository(log *logrus.Logger) Interface.ICartItemRepository {
	return &CartItemRepository{log: log}
}

func (c *CartItemRepository) GetPaginatedCartItemList(searchValue, sortBy, cartId, productId string, pageIndex, pageSize int) (Util.PaginatedList[BusinessObjects.CartItem], error) {
	c.log.Infof("Fetching paginated cart items list with searchValue: %s, sortBy: %s, cartId: %s, productId: %s, pageIndex: %d, pageSize: %d", searchValue, sortBy, cartId, productId, pageIndex, pageSize)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
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
		c.log.Error("Failed to count cart items:", err)
		return Util.PaginatedList[BusinessObjects.CartItem]{}, err
	}

	// Paginate the results
	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&cartItems).Error; err != nil {
		c.log.Error("Failed to fetch paginated cart items:", err)
		return Util.PaginatedList[BusinessObjects.CartItem]{}, err
	}

	c.log.Infof("Successfully fetched paginated cart items list with total count: %d", total)
	return Util.NewPaginatedList(cartItems, total, pageIndex, pageSize), nil
}

// GetAllCartItems retrieves all freight cartItems from the database
func (c *CartItemRepository) GetAllCartItems() ([]BusinessObjects.CartItem, error) {
	c.log.Info("Fetching all cart items")
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return nil, err
	}

	var cartItem []BusinessObjects.CartItem
	if err := db.Find(&cartItem).Error; err != nil {
		c.log.Error("Failed to fetch all cart items:", err)
		return nil, err
	}

	c.log.Info("Successfully fetched all cart items")
	return cartItem, nil
}

// GetCartItemByID retrieves a freight cartItem by its ID
func (c *CartItemRepository) GetCartItemByCartID(cartItemID string) ([]BusinessObjects.CartItem, error) {
	c.log.Infof("Fetching cart item by cart ID: %s", cartItemID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return nil, err
	}

	var cartItem []BusinessObjects.CartItem
	if err := db.Find(&cartItem, "cart_id = ?", cartItemID).Error; err != nil {
		c.log.Error("Failed to fetch cart item by cart ID:", err)
		return nil, err
	}

	c.log.Infof("Successfully fetched cart item by cart ID: %s", cartItemID)
	return cartItem, nil
}

// CreateCartItem adds a new freight cartItem to the database
func (c *CartItemRepository) CreateCartItem(cartItem BusinessObjects.CartItem) error {
	c.log.Infof("Creating new cart item with cart ID: %s and product ID: %s", cartItem.CartID, cartItem.ProductID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Create(&cartItem).Error; err != nil {
		c.log.Error("Failed to create new cart item:", err)
		return err
	}

	c.log.Infof("Successfully created new cart item with cart ID: %s and product ID: %s", cartItem.CartID, cartItem.ProductID)
	return nil
}

// UpdateCartItem updates an existing freight cartItem
func (c *CartItemRepository) UpdateCartItem(cartItem BusinessObjects.CartItem) error {
	c.log.Infof("Updating cart item with cart ID: %s and product ID: %s", cartItem.CartID, cartItem.ProductID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Save(&cartItem).Error; err != nil {
		c.log.Error("Failed to update cart item:", err)
		return err
	}

	c.log.Infof("Successfully updated cart item with cart ID: %s and product ID: %s", cartItem.CartID, cartItem.ProductID)
	return nil
}

// DeleteCartItem removes a freight cartItem from the database by its ID
func (c *CartItemRepository) DeleteCartItem(cartID, productID string) error {
	c.log.Infof("Deleting cart item with cart ID: %s and product ID: %s", cartID, productID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Where("cart_id = ? AND product_id = ?", cartID, productID).Delete(&BusinessObjects.CartItem{}).Error; err != nil {
		c.log.Error("Failed to delete cart item:", err)
		return err
	}

	c.log.Infof("Successfully deleted cart item with cart ID: %s and product ID: %s", cartID, productID)
	return nil
}

func (c *CartItemRepository) UpdateOrCreateCartItem(cartItem BusinessObjects.CartItem) error {
	c.log.Infof("Updating or creating cart item with cart ID: %s and product ID: %s", cartItem.CartID, cartItem.ProductID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	var existingCartItem BusinessObjects.CartItem
	if err := db.Where("cart_id = ? AND product_id = ?", cartItem.CartID, cartItem.ProductID).First(&existingCartItem).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create new cart item if not found
			c.log.Infof("Cart item not found, creating new cart item with cart ID: %s and product ID: %s", cartItem.CartID, cartItem.ProductID)
			if err := db.Create(&cartItem).Error; err != nil {
				c.log.Error("Failed to create new cart item:", err)
				return err
			}
		} else {
			c.log.Error("Failed to fetch cart item for update or create:", err)
			return err
		}
	} else {
		// Update existing cart item
		c.log.Infof("Cart item found, updating cart item with cart ID: %s and product ID: %s", cartItem.CartID, cartItem.ProductID)
		existingCartItem.Quantity = cartItem.Quantity
		if err := db.Save(&existingCartItem).Error; err != nil {
			c.log.Error("Failed to update cart item:", err)
			return err
		}
	}

	c.log.Infof("Successfully updated or created cart item with cart ID: %s and product ID: %s", cartItem.CartID, cartItem.ProductID)
	return nil
}
