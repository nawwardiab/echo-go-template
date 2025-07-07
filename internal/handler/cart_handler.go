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
	sess    session.Session
	prodSvc service.ProductService
}

func NewCartHandler(sess *session.Session, prodSvc service.ProductService) *CartHandler {
	return &CartHandler{sess: *sess, prodSvc: prodSvc}
}

// GET /cart
func (ch *CartHandler) ViewCart(c echo.Context) error {
	if !ch.sess.Has(c.Request()) {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	cart, loadErr := loadCart(&ch.sess, c.Request())
	if loadErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot load cart")
	}

	var items []model.CartItem
	for pid, qty := range cart {
		prod, _ := ch.prodSvc.GetProductByID(pid) // ignore not-found for brevity
		items = append(items, model.CartItem{
			ProductID: pid,
			Quantity:  qty,
			Product:   *prod,
		})
	}

	return c.Render(http.StatusOK, "cart.tpl", map[string]any{"CartItems": items})
}

// POST /cart/add
func (ch *CartHandler) AddToCart(c echo.Context) error {
	if !ch.sess.Has(c.Request()) {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	pid, _ := strconv.Atoi(c.FormValue("product_id"))
	qty, _ := strconv.Atoi(c.FormValue("quantity"))

	cart, loadErr := loadCart(&ch.sess, c.Request())
	if loadErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not load cart")
	}

	updated := service.AddToCart(cart, pid, qty)

	saveErr := saveCart(&ch.sess, c.Response().Writer, c.Request(), updated);
	if saveErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cannot save cart")
	}

	return c.Redirect(http.StatusSeeOther, "/cart")
}

// POST /cart/remove
func (ch *CartHandler) RemoveFromCart(c echo.Context) error {
	if !ch.sess.Has(c.Request()) {
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	
	pid, convErr := strconv.Atoi(c.FormValue("product_id"))
	if convErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid product id")
	}

	cart, loadErr := loadCart(&ch.sess, c.Request())
	if loadErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not load cart")
	}

	updated := service.RemoveFromCart(cart, pid)

	if err := saveCart(&ch.sess, c.Response().Writer, c.Request(), updated); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not save cart")
	}

	return c.Redirect(http.StatusSeeOther, "/cart")
}


// HELPERS
// loadCart returns the cart map stored in the session or an empty one.
func loadCart(s *session.Session, r *http.Request) (service.CartMap, error) {
	raw, getErr := s.Get(r, "cart")
	if getErr != nil {
		return nil, getErr
	}
	if raw == "" {
		return make(service.CartMap), nil
	}

	var cart service.CartMap
	unmarshalErr := yaml.Unmarshal([]byte(raw), &cart);
	if  unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return cart, nil
}

// saveCart marshals the cart map and stores it back in the session.
func saveCart(s *session.Session, w http.ResponseWriter, r *http.Request, cart service.CartMap) error {
	enc, marshalErr := yaml.Marshal(cart)
	if marshalErr != nil {
		return marshalErr
	}
	return s.Set(w, r, "cart", string(enc))
}
