package httpsrv

import (
	"fmt"
	"fruitshop/internal/cartsvc"
	"fruitshop/internal/controllers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//initializeRoutes function will initialize http routes for the application using gorilla mux
func applyRoutes(db *gorm.DB, m *mux.Router) {
	s := controllers.TooCommon{db} // should be replaced by individual services/packages
	cartSvc := cartsvc.New(db)

	// Customer routes
	m.HandleFunc("/server/customers", setMiddlewareJSON(s.CreateCustomer)).Methods("POST")
	m.HandleFunc("/server/customers/{loginid}", setMiddlewareJSON(s.GetCustomer)).Methods("GET")

	// Fruit routes
	m.HandleFunc("/server/fruits", setMiddlewareJSON(s.GetFruits)).Methods("GET")

	// CartItem routes
	m.HandleFunc("/server/cartitem", setMiddlewareJSON(cartSvc.CreateItem)).Methods("POST")
	m.HandleFunc("/server/cartitem", setMiddlewareJSON(cartSvc.UpdateItem)).Methods("PUT")
	m.HandleFunc("/server/cartitem/{cart_id}/{fruitname}", setMiddlewareJSON(cartSvc.DeleteItem)).Methods("DELETE")
	m.HandleFunc("/server/cartitems/{cart_id}", setMiddlewareJSON(cartSvc.GetItems)).Methods("GET")

	// Cart route
	m.HandleFunc("/server/cart/{cart_id}", setMiddlewareJSON(cartSvc.Get)).Methods("GET")

	// Discounts routes
	m.HandleFunc("/server/discounts/{cart_id}", setMiddlewareJSON(s.GetAppliedDiscounts)).Methods("GET")

	// Coupon route
	m.HandleFunc("/server/orangecoupon/{cart_id}/{fruit_id}", setMiddlewareJSON(s.ApplyTimeSensitiveCoupon)).Methods("GET")

	// Pay route
	m.HandleFunc("/server/pay", setMiddlewareJSON(s.Pay)).Methods("POST")

	// Serves angular application on / endpoint
	m.PathPrefix("/").Handler(http.FileServer(http.Dir("frontend/dist/fruitshop-ui"))) // for docker
	//m.PathPrefix("/").Handler(http.FileServer(http.Dir("../frontend/dist/fruitshop-ui"))) // for local

	fmt.Println()
	fmt.Println("These below are the initialized routes for the application : ")
	fmt.Println()
	err := m.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	http.Handle("/", m)
}

func setMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}
