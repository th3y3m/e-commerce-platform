package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ShoppingCartRepository struct {
	log *logrus.Logger
}

func NewShoppingCartRepository(log *logrus.Logger) Interface.IShoppingCartRepository {
	return &ShoppingCartRepository{log: log}
}

func (s *ShoppingCartRepository) GetPaginatedShoppingCartList(sortBy, cartID, userID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.ShoppingCart], error) {
	s.log.Infof("Fetching paginated shopping cart list with sortBy: %s, cartID: %s, userID: %s, pageIndex: %d, pageSize: %d, status: %v", sortBy, cartID, userID, pageIndex, pageSize, status)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		s.log.Error("Failed to connect to PostgreSQL:", err)
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
		s.log.Error("Failed to count shopping carts:", err)
		return Util.PaginatedList[BusinessObjects.ShoppingCart]{}, err
	}

	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&carts).Error; err != nil {
		s.log.Error("Failed to fetch paginated shopping carts:", err)
		return Util.PaginatedList[BusinessObjects.ShoppingCart]{}, err
	}

	s.log.Infof("Successfully fetched paginated shopping cart list with total count: %d", total)
	return Util.NewPaginatedList(carts, total, pageIndex, pageSize), nil
}

func (s *ShoppingCartRepository) GetUserShoppingCart(userID string) (BusinessObjects.ShoppingCart, error) {
	s.log.Infof("Fetching shopping cart for user ID: %s", userID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		s.log.Error("Failed to connect to PostgreSQL:", err)
		return BusinessObjects.ShoppingCart{}, err
	}

	var cart BusinessObjects.ShoppingCart
	if err := db.Order("created_at DESC").First(&cart, "user_id = ? AND status = ?", userID, true).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.log.Infof("No active shopping cart found for user ID: %s, creating a new one", userID)
			// Create a new shopping cart if not found
			newCart := BusinessObjects.ShoppingCart{
				UserID: userID,
				Status: true,
				// Add other necessary fields here
			}
			if err := db.Create(&newCart).Error; err != nil {
				s.log.Error("Failed to create new shopping cart:", err)
				return BusinessObjects.ShoppingCart{}, err
			}
			s.log.Infof("Successfully created new shopping cart for user ID: %s", userID)
			return newCart, nil
		}
		s.log.Error("Failed to fetch shopping cart for user ID:", err)
		return BusinessObjects.ShoppingCart{}, err
	}

	s.log.Infof("Successfully fetched shopping cart for user ID: %s", userID)
	return cart, nil
}

// GetAllShoppingCarts retrieves all shopping carts from the database
func (s *ShoppingCartRepository) GetAllShoppingCarts() ([]BusinessObjects.ShoppingCart, error) {
	s.log.Info("Fetching all shopping carts")
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		s.log.Error("Failed to connect to PostgreSQL:", err)
		return nil, err
	}

	var carts []BusinessObjects.ShoppingCart
	if err := db.Find(&carts).Error; err != nil {
		s.log.Error("Failed to fetch all shopping carts:", err)
		return nil, err
	}

	s.log.Info("Successfully fetched all shopping carts")
	return carts, nil
}

// GetShoppingCartByID retrieves a shopping cart by its ID
func (s *ShoppingCartRepository) GetShoppingCartByID(cartID string) (BusinessObjects.ShoppingCart, error) {
	s.log.Infof("Fetching shopping cart by ID: %s", cartID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		s.log.Error("Failed to connect to PostgreSQL:", err)
		return BusinessObjects.ShoppingCart{}, err
	}

	var cart BusinessObjects.ShoppingCart
	if err := db.First(&cart, "cart_id = ?", cartID).Error; err != nil {
		s.log.Error("Failed to fetch shopping cart by ID:", err)
		return BusinessObjects.ShoppingCart{}, err
	}

	s.log.Infof("Successfully fetched shopping cart by ID: %s", cartID)
	return cart, nil
}

// CreateShoppingCart adds a new shopping cart to the database
func (s *ShoppingCartRepository) CreateShoppingCart(cart BusinessObjects.ShoppingCart) (BusinessObjects.ShoppingCart, error) {
	s.log.Infof("Creating new shopping cart with ID: %s", cart.CartID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		s.log.Error("Failed to connect to PostgreSQL:", err)
		return BusinessObjects.ShoppingCart{}, err
	}

	if err := db.Create(&cart).Error; err != nil {
		s.log.Error("Failed to create new shopping cart:", err)
		return BusinessObjects.ShoppingCart{}, err
	}

	s.log.Infof("Successfully created new shopping cart with ID: %s", cart.CartID)
	return cart, nil
}

// UpdateShoppingCart updates an existing shopping cart
func (s *ShoppingCartRepository) UpdateShoppingCart(cart BusinessObjects.ShoppingCart) error {
	s.log.Infof("Updating shopping cart with ID: %s", cart.CartID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		s.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Save(&cart).Error; err != nil {
		s.log.Error("Failed to update shopping cart:", err)
		return err
	}

	s.log.Infof("Successfully updated shopping cart with ID: %s", cart.CartID)
	return nil
}

// DeleteShoppingCart removes a shopping cart from the database by its ID
func (s *ShoppingCartRepository) DeleteShoppingCart(cartID string) error {
	s.log.Infof("Deleting shopping cart with ID: %s", cartID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		s.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Delete(&BusinessObjects.ShoppingCart{}, "cart_id = ?", cartID).Error; err != nil {
		s.log.Error("Failed to delete shopping cart:", err)
		return err
	}

	s.log.Infof("Successfully deleted shopping cart with ID: %s", cartID)
	return nil
}

func (s *ShoppingCartRepository) UpdateShoppingCartStatus(cartID string, status bool) error {
	s.log.Infof("Updating shopping cart status for cart ID: %s to %v", cartID, status)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		s.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Model(&BusinessObjects.ShoppingCart{}).Where("cart_id = ?", cartID).Update("status", status).Error; err != nil {
		s.log.Error("Failed to update shopping cart status:", err)
		return err
	}

	s.log.Infof("Successfully updated shopping cart status for cart ID: %s to %v", cartID, status)
	return nil
}
