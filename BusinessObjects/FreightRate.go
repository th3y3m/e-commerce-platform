package BusinessObjects

// FreightRate represents the FreightRates table
type FreightRate struct {
	RateID        string  `gorm:"primaryKey;column:rate_id"`
	CourierID     string  `gorm:"column:courier_id"`
	DistanceMinKM int     `gorm:"column:distance_min_km"`
	DistanceMaxKM int     `gorm:"column:distance_max_km"`
	CostPerKM     float64 `gorm:"column:cost_per_km"`
	Status        bool    `gorm:"column:status"`
}
