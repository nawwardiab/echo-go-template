package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v3"

	"echo-server/internal/model"
	"echo-server/internal/service"
	"echo-server/internal/session"
)

type CartHandler struct {
	prodSvc service.ProductService
}

func NewCartHandler(prodSvc service.ProductService) *CartHandler {
	return &CartHandler{prodSvc: prodSvc}
}

// GET /cart
func (ch *CartHandler) ViewCart(c echo.Context) error {
	if session.GetValue(c, "user_id") == nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	} else {
		cart, loadErr := loadCart(c)
		if loadErr != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "cannot load cart")
		} else {
			var items []model.CartItem
			for pid, qty := range cart {
				prod, _ := ch.prodSvc.GetProductByID(pid)
				items = append(items, model.CartItem{
					ProductID: pid,
					Quantity:  qty,
					Product:   *prod,
				})
			}
			return c.Render(http.StatusOK, "cart.tpl", map[string]any{"CartItems": items})
		}
	}
}

// POST /cart/add
func (ch *CartHandler) AddToCart(c echo.Context) error {
	if session.GetValue(c, "user_id") == nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	} else {
		pid, _ := strconv.Atoi(c.FormValue("product_id"))
		qty, _ := strconv.Atoi(c.FormValue("quantity"))

		cart, loadErr := loadCart(c)
		if loadErr != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "could not load cart")
		} else {
			updated := service.AddToCart(cart, pid, qty)
			saveErr := saveCart(c, updated);
			
			if saveErr != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "cannot save cart")
			} else {
				return c.Redirect(http.StatusSeeOther, "/cart")
			}
		}
	}
}

// POST /cart/remove
func (ch *CartHandler) RemoveFromCart(c echo.Context) error {
	// checks if user logged in
	if session.GetValue(c, "user_id") == nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	} else {
		// retrieves product_id value and converts to int
		pid, convErr := strconv.Atoi(c.FormValue("product_id"))
		if convErr != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid product id")
		} else {
			// loads cart map
			cart, loadErr := loadCart(c)
			if loadErr != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "could not load cart")
			} else {				
				// updates cart map
				updated := service.RemoveFromCart(cart, pid)
				saveErr := saveCart(c, updated)
				if saveErr != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, "could not save cart")
				} else {
					return c.Redirect(http.StatusSeeOther, "/cart")
				}
			}
		}
	}
}


// HELPERS
// loadCart returns the cart map stored in the session or an empty one.
func loadCart(c echo.Context) (service.CartMap, error) {
	raw := session.GetValue(c, "cart")
	str, ok := raw.(string)
	if raw == nil {
		return make(service.CartMap), nil
	} else if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "malformed cart data")
	} else {	
		var cart service.CartMap
		unmarshalErr := yaml.Unmarshal([]byte(str), &cart);
		if  unmarshalErr != nil {
			return nil, unmarshalErr
		}
		return cart, nil
	}	
}

// saveCart marshals the cart map and stores it back in the session.
func saveCart(c echo.Context, cart service.CartMap) error {
	enc, marshalErr := yaml.Marshal(cart)
	if marshalErr != nil {
		return marshalErr
	} else {
		return session.Set(c, "cart", string(enc))
	}
}
