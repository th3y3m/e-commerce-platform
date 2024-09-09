package BusinessObjects

// FreightRate represents the FreightRates table
type FreightRate struct {
	RateID        string  `json:"rate_id" db:"rate_id"`
	CourierID     string  `json:"courier_id" db:"courier_id"`
	DistanceMinKM int     `json:"distance_min_km" db:"distance_min_km"`
	DistanceMaxKM int     `json:"distance_max_km" db:"distance_max_km"`
	CostPerKM     float64 `json:"cost_per_km" db:"cost_per_km"`
}
