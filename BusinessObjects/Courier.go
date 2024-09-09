package BusinessObjects

type Courier struct {
	CourierID string `json:"courier_id" db:"courier_id"`
	Courier   string `json:"courier" db:"courier"`
}
