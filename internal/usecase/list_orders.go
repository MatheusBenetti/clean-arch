package usecase

import (
	"github.com/MatheusBenetti/clean-architecture/internal/entity"
	"github.com/MatheusBenetti/clean-architecture/pkg/events"
)

type ListOrdersOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrdersListed    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrdersListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		OrdersListed:    OrdersListed,
		EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrdersUseCase) Execute() ([]ListOrdersOutputDTO, error) {
	orders, err := c.OrderRepository.ListOrders()
	if err != nil {
		return nil, err
	}
	var ordersDTO []ListOrdersOutputDTO
	for _, order := range orders {
		ordersDTO = append(ordersDTO, ListOrdersOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	c.OrdersListed.SetPayload(ordersDTO)
	c.EventDispatcher.Dispatch(c.OrdersListed)

	return ordersDTO, nil
}
