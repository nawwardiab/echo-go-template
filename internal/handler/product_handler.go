package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"echo-server/internal/service"
	"echo-server/internal/session"
)

type ProductHandler struct {
	prodSvc service.ProductService
}

func NewProductHandler(prodSvc service.ProductService, ) *ProductHandler {
	return &ProductHandler{prodSvc: prodSvc}
}


// GET /products  – list page

func (ph *ProductHandler) ListProducts(c echo.Context) error {
	if session.GetValue(c, "user_id") == nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	} else {
		products, getErr := ph.prodSvc.GetProducts()
		if getErr != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "server error")
		} else {			
			return c.Render(http.StatusOK, "products.tpl", map[string]any{
				"Products": products,
			})
		}
	}
}

// GET /products/:id  – detail page
func (ph *ProductHandler) ListProductDetails(c echo.Context) error {
	if session.GetValue(c, "user_id") == nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	} else {
		id, convErr := strconv.Atoi(c.Param("id"))
		product, getErr := ph.prodSvc.GetProductByID(id)
		
		if convErr != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid product id")
		} else if getErr != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "server error")
		} else {
			return c.Render(http.StatusOK, "singleProduct.tpl", map[string]any{
				"Product": product,
			})
		}
	}
}
