package ports

import (
	"github.com/vteja-code/microservices/order/internal/application/domain"
)

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}
