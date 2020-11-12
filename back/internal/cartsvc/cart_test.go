package cartsvc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"fruitshop/internal/testdb"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/assert.v1"
)

func TestGetCartByCartID(t *testing.T) {
	db, err := testdb.Database()
	if err != nil {
		t.Fatal(err)
	}
	s := New(db)

	if err = testdb.RefreshCartTable(db); err != nil {
		t.Fatal(err)
	}
	cart, err := testdb.SeedOneCart(db)
	if err != nil {
		t.Fatal(err)
	}
	userSample := []struct {
		statusCode   int
		status       string
		total        float64
		totalsavings float64
	}{
		{
			statusCode:   200,
			status:       "OPEN",
			total:        cart.Total,
			totalsavings: cart.TotalSavings,
		},
	}
	for _, v := range userSample {

		req, err := http.NewRequest("GET", "/carts", nil)
		if err != nil {
			t.Errorf("This is the error: %v\n", err)
		}
		req = mux.SetURLVars(req, map[string]string{"cart_id": "1"})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.Get)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}

		fmt.Println(responseMap)
		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 200 {
			assert.Equal(t, cart.Total, responseMap["total"])
			assert.Equal(t, cart.TotalSavings, responseMap["totalsavings"])
		}
	}
}

func TestGetCartByCartIDNotAvailble(t *testing.T) {
	db, err := testdb.Database()
	if err != nil {
		t.Fatal(err)
	}
	s := New(db)

	if err := testdb.RefreshCartTable(db); err != nil {
		t.Fatal(err)
	}

	userSample := []struct {
		statusCode   int
		status       string
		total        float64
		totalsavings float64
	}{
		{
			statusCode: 400,
			status:     "OPEN",
		},
	}
	for _, v := range userSample {
		req, err := http.NewRequest("GET", "/carts", nil)
		if err != nil {
			t.Errorf("This is the error: %v\n", err)
		}
		req = mux.SetURLVars(req, map[string]string{"cart_id": "1"})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.Get)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.Bytes()), &responseMap)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}

		fmt.Println(responseMap)
		assert.Equal(t, rr.Code, v.statusCode)
	}
}

func TestCreateCartItem(t *testing.T) {
	db, err := testdb.Database()
	if err != nil {
		t.Fatal(err)
	}
	s := New(db)

	if err := testdb.RefreshCartItemTable(db); err != nil {
		t.Fatal(err)
	}

	if _, err = testdb.SeedFruits(db); err != nil {
		t.Fatal(err)
	}
	samples := []struct {
		inputJSON    string
		statusCode   int
		CartID       uint
		FruitID      uint
		itemtotal    float64
		errorMessage string
	}{
		{
			inputJSON: `{
				"cartid": 1,
				"name": "Apple",
				"quantity": 2
			}`,
			statusCode:   201,
			CartID:       1,
			FruitID:      1,
			itemtotal:    2.0,
			errorMessage: "",
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("POST", "/cartitem", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.CreateItem)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.Bytes()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 201 {
			assert.Equal(t, responseMap["itemtotal"], v.itemtotal)

		}
	}
}

func TestUpdateCartItem(t *testing.T) {
	db, err := testdb.Database()
	if err != nil {
		t.Fatal(err)
	}
	s := New(db)

	if err = testdb.RefreshCartItemTable(db); err != nil {
		t.Fatal(err)
	}
	if _, err = testdb.SeedFruits(db); err != nil {
		t.Fatal(err)
	}
	if _, err = testdb.SeedOneCart(db); err != nil {
		t.Fatal(err)
	}
	if _, err = testdb.SeedOneCartItem(db); err != nil {
		t.Fatal(err)
	}
	samples := []struct {
		inputJSON    string
		statusCode   int
		CartID       uint
		FruitID      uint
		quantity     uint
		errorMessage string
	}{
		{
			inputJSON: `{
				"cartid": 1,
				"name": "Apple",
				"quantity": 2
			}`,
			statusCode:   500,
			CartID:       1,
			FruitID:      1,
			quantity:     4,
			errorMessage: "",
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("PUT", "/cartitem", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.UpdateItem)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.Bytes()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 201 {
			assert.NotEqual(t, responseMap["quantity"], v.quantity)

		}
	}
}

func TestDeleteCartItem(t *testing.T) {
	db, err := testdb.Database()
	if err != nil {
		t.Fatal(err)
	}
	s := New(db)

	if err = testdb.RefreshCartItemTable(db); err != nil {
		t.Fatal(err)
	}
	if _, err = testdb.SeedFruits(db); err != nil {
		t.Fatal(err)
	}
	if _, err = testdb.SeedOneCartItem(db); err != nil {
		t.Fatal(err)
	}
	if _, err = testdb.SeedOneCart(db); err != nil {
		t.Fatal(err)
	}
	samples := []struct {
		inputJSON    string
		statusCode   int
		CartID       uint
		FruitID      uint
		quantity     uint
		errorMessage string
	}{
		{
			inputJSON: `{
				"cartid": 1,
				"name": "Apple",
				"quantity": 0
			}`,
			statusCode:   200,
			CartID:       1,
			FruitID:      1,
			quantity:     4,
			errorMessage: "",
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("DELETE", "/cartitem", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.DeleteItem)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.Bytes()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 200 {
			assert.NotEqual(t, responseMap["quantity"], v.quantity)

		}
	}
}
