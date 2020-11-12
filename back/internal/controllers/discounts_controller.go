package controllers

import (
	"net/http"

	"fruitshop/internal/models"
	"fruitshop/internal/responses"

	"github.com/gorilla/mux"
)

//GetAppliedDiscounts will fetch all the discounts applied on a given cart_id of the customer
func (s *TooCommon) GetAppliedDiscounts(w http.ResponseWriter, r *http.Request) {
	// Reading cart_id from request params
	vars := mux.Vars(r)
	cartid := vars["cart_id"]
	discount := models.Discount{}
	appliedDiscounts := discount.FindAllDiscounts(s.DB, cartid)
	responses.JSON(w, http.StatusOK, appliedDiscounts)
}
