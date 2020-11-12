package controllertests

import (
	"encoding/json"
	"fruitshop/internal/models"
	"fruitshop/internal/testdb"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/assert.v1"
)

func TestGetDiscounts(t *testing.T) {

	err := testdb.RefreshDiscountsTable(db)
	if err != nil {
		log.Fatal(err)
	}
	_, err = testdb.SeedSingleItemDiscount(db)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("GET", "/discounts", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	req = mux.SetURLVars(req, map[string]string{"cart_id": "1"})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetAppliedDiscounts)
	handler.ServeHTTP(rr, req)

	var discounts []models.Discount
	err = json.Unmarshal([]byte(rr.Body.Bytes()), &discounts)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(discounts), 1)
}

func TestGetDiscountsNoDiscountAvailble(t *testing.T) {

	err := testdb.RefreshDiscountsTable(db)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/discounts", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	req = mux.SetURLVars(req, map[string]string{"cart_id": "1"})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetAppliedDiscounts)
	handler.ServeHTTP(rr, req)

	var discounts []models.Discount
	err = json.Unmarshal([]byte(rr.Body.Bytes()), &discounts)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(discounts), 0)
}
