package BusinessObjects

// Courier represents the Couriers table
type Courier struct {
	CourierID string `gorm:"primaryKey;column:courier_id"`
	Courier   string `gorm:"column:courier"`
	Status    bool   `gorm:"column:status"`

	Orders       []Order       `gorm:"foreignKey:CourierID"`
	FreightRates []FreightRate `gorm:"foreignKey:CourierID"`
}
