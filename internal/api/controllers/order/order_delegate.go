package order

import (
	"net/http"
	"order-service/internal/api/problems"
	"order-service/internal/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type OrderControllerDelegate struct {
	orderService service.OrderService
}

//you can get a pointer using the & operator (e.g. &c)
//you can dereference a pointer to get the value using *
func NewOrderControllerDelegate(orderService service.OrderService) *OrderControllerDelegate {
	return &OrderControllerDelegate{orderService: orderService}
}

func (p OrderControllerDelegate) CreateOrder(c *gin.Context) {
	log.Info().Msg("Order created check coupon")
}

func (p OrderControllerDelegate) GetById(c *gin.Context, orderId string) {
	var q OrderDto
	_ = c.Bind(&q)
	order, err := p.orderService.Get(orderId)

	if order == nil {
		c.JSON(http.StatusNotFound, problems.NewProblem(404, "/orders", "No found", "Order not found"))
	} else if err != nil {
		log.Error().Err(err).Msg("Error retrieving order")
		//problem := *problems.NewProblem(http.StatusInternalServerError, fmt.Sprintf("/api/users"), "Response processing error", "Errore recupero lista utenti")
		c.JSON(http.StatusInternalServerError, nil)
		return
	} else {
		c.JSON(http.StatusOK, order)
	}
}
