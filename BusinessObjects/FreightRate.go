package BusinessObjects

// FreightRate represents the FreightRates table
type FreightRate struct {
	RateID         string  `json:"rate_id" db:"rate_id"`
	Courier        string  `json:"courier" db:"courier"`
	ShippingMethod string  `json:"shipping_method" db:"shipping_method"`
	DistanceMinKM  int     `json:"distance_min_km" db:"distance_min_km"`
	DistanceMaxKM  int     `json:"distance_max_km" db:"distance_max_km"`
	CostPerKM      float64 `json:"cost_per_km" db:"cost_per_km"`
}
