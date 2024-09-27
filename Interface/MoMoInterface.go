package Interface

import (
	"net/url"
	"th3y3m/e-commerce-platform/BusinessObjects"
)

type IMoMoService interface {
	CreateMoMoUrl(amount float64, orderId string) (string, error)
	ValidateMoMoResponse(queryString url.Values) (*BusinessObjects.TransactionStatusModel, error)
}
