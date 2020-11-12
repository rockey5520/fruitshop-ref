package controllertests

import (
	"encoding/json"
	"fruitshop/internal/models"
	"fruitshop/internal/testdb"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetFruits(t *testing.T) {

	err := testdb.RefreshFruitTable(db)
	if err != nil {
		log.Fatal(err)
	}
	_, err = testdb.SeedFruits(db)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("GET", "/fruits", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetFruits)
	handler.ServeHTTP(rr, req)

	var fruits []models.Fruit
	err = json.Unmarshal([]byte(rr.Body.Bytes()), &fruits)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(fruits), 4)
}
