package ports

import "github.com/vteja-code/microservices/order/internal/application/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
