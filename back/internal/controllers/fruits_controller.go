package controllers

import (
	"net/http"

	"fruitshop/internal/models"
	"fruitshop/internal/responses"
)

// GetFruits fetches all fruits from the fruit table which is meta table where data loaded during application start-up
func (s *TooCommon) GetFruits(w http.ResponseWriter, r *http.Request) {
	fruit := models.Fruit{}
	fruits, err := fruit.FindAllFruits(s.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, fruits)
}
