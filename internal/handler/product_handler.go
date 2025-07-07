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
	sess    session.Session
}

func NewProductHandler(prodSvc service.ProductService, sess *session.Session) *ProductHandler {
	return &ProductHandler{prodSvc: prodSvc, sess: *sess}
}


// GET /products  – list page

func (ph *ProductHandler) ListProducts(c echo.Context) error {
	if !ph.sess.Has(c.Request()) {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	products, getErr := ph.prodSvc.GetProducts()
	if getErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "server error")
	}

	return c.Render(http.StatusOK, "products.tpl", map[string]any{
		"Products": products,
	})
}

// GET /products/:id  – detail page
func (ph *ProductHandler) ListProductDetails(c echo.Context) error {
	if !ph.sess.Has(c.Request()) {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	id, convErr := strconv.Atoi(c.Param("id"))
	if convErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid product id")
	}

	product, getErr := ph.prodSvc.GetProductByID(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "server error")
	}

	return c.Render(http.StatusOK, "singleProduct.tpl", map[string]any{
		"Product": product,
	})
}
