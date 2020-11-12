package calc

import (
	"fmt"
	"fruitshop/internal/models"

	"github.com/jinzhu/gorm"
)

// RecalcualtePayments recalcuates the cart value and its saving based on the cart items
func RecalcualtePayments(db *gorm.DB, cartID uint) {
	// Fetch all items in a given cart
	var cartItems []models.CartItem
	if err := db.Where("cart_id = ?", cartID).Find(&cartItems).Error; err != nil {
		fmt.Println("Error ", err)
	}

	// calcuate the total cost of the cartitems and total discounts applied
	var totalCost float64
	var totalDiscountedCost float64
	for _, item := range cartItems {
		totalCost += item.ItemTotal
		totalDiscountedCost += item.ItemDiscountedTotal
	}
	var cart models.Cart
	if err := db.Where("ID = ?", cartID).Find(&cart).Error; err != nil {
		fmt.Println("Error ", err)
	}

	// Update Cart table with total cost and total savings
	db.Model(&cart).Update("total", totalCost).Update("total_savings", totalDiscountedCost)
}
