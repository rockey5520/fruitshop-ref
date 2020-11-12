package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fruitshop/internal/models"
	"fruitshop/internal/responses"
	"fruitshop/internal/utils/formaterror"

	"github.com/gorilla/mux"
)

// CreateCustomer will create a new customer entry into the database along with first cart to add items later during purchase
func (s *TooCommon) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	// Reading the request body from http request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	// Creating customer , cart structs and mapping request body to customer and a creating new card with customer ID
	customer := models.Customer{}
	newcart := models.Cart{
		Total:  0.0,
		Status: "OPEN",
	}
	customer.Cart = newcart
	err = json.Unmarshal(body, &customer)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// Customer validation for the payload sent
	err = customer.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// Saves the customer details in the database
	createdCustomer, err := customer.SaveCustomer(s.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, createdCustomer.ID))
	responses.JSON(w, http.StatusCreated, createdCustomer)
}

//GetCustomer will fetch the information about the customer based on the loginid provided which is unique to each customer
func (s *TooCommon) GetCustomer(w http.ResponseWriter, r *http.Request) {
	// Reading loginid from request params
	vars := mux.Vars(r)
	loginid := vars["loginid"]

	customer := models.Customer{}
	customerFetched, err := customer.FindCustomerByLoginID(s.DB, loginid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, customerFetched)
}
