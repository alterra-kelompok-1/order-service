package order

import "github.com/labstack/echo/v4"

func RegisterHandlers(e *echo.Echo, h Handler) {
	e.GET("/orders", h.Get)
	e.POST("/orders", h.Create)

	// e.GET("/orders/:id", h.GetByID)
	// e.DELETE("/orders/:id", h.Delete)
	e.PUT("/orders/:id", h.UpdateOrderByID)
}
