package testdb

import (
	"log"
	"os"

	"fruitshop/internal/models"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Database() (*gorm.DB, error) {
	testDbName := os.Getenv("TestDbName")
	db, err := gorm.Open("sqlite3", testDbName)
	if err != nil {
		return nil, err
	}
	db.Exec("PRAGMA foreign_keys = ON")
	db.LogMode(true)

	return db, nil
}

func RefreshCustomerTable(db *gorm.DB) error {
	err := db.DropTableIfExists(&models.Customer{}).Error
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Customer{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed customer table")
	return nil
}

func RefreshDiscountsTable(db *gorm.DB) error {
	err := db.DropTableIfExists(
		&models.SingleItemDiscount{},
		&models.DualItemDiscount{},
		&models.SingleItemCoupon{},
		&models.AppliedDualItemDiscount{},
		&models.AppliedSingleItemCoupon{},
		&models.AppliedSingleItemDiscount{}).Error
	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		&models.SingleItemDiscount{},
		&models.DualItemDiscount{},
		&models.SingleItemCoupon{},
		&models.AppliedDualItemDiscount{},
		&models.AppliedSingleItemCoupon{},
		&models.AppliedSingleItemDiscount{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed discounts table")
	return nil
}

func RefreshFruitTable(db *gorm.DB) error {
	err := db.DropTableIfExists(&models.Fruit{}).Error
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Fruit{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed Fruit table")
	return nil
}

func RefreshCartTable(db *gorm.DB) error {
	err := db.DropTableIfExists(&models.Cart{},
		&models.CartItem{},
		&models.Payment{},
		&models.AppliedDualItemDiscount{},
		&models.AppliedSingleItemDiscount{},
		&models.AppliedSingleItemCoupon{}).Error
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Cart{},
		&models.CartItem{},
		&models.Payment{},
		&models.AppliedDualItemDiscount{},
		&models.AppliedSingleItemDiscount{},
		&models.AppliedSingleItemCoupon{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed cart table")
	log.Printf("refreshCartTable routine OK !!!")
	return nil
}

func RefreshCartItemTable(db *gorm.DB) error {
	err := db.DropTableIfExists(
		&models.Cart{},
		&models.CartItem{},
		&models.Fruit{},
		&models.Payment{},
		&models.SingleItemDiscount{},
		&models.DualItemDiscount{},
		&models.SingleItemCoupon{},
		&models.AppliedDualItemDiscount{},
		&models.AppliedSingleItemDiscount{},
		&models.AppliedSingleItemCoupon{}).Error

	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		&models.Cart{},
		&models.CartItem{},
		&models.Fruit{},
		&models.Payment{},
		&models.SingleItemDiscount{},
		&models.DualItemDiscount{},
		&models.SingleItemCoupon{},
		&models.AppliedDualItemDiscount{},
		&models.AppliedSingleItemDiscount{},
		&models.AppliedSingleItemCoupon{},
	).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed CartItem table")
	log.Printf("refreshCartItemTable routine OK !!!")
	return nil
}

func SeedOneCart(db *gorm.DB) (models.Cart, error) {
	if err := RefreshCartTable(db); err != nil {
		return models.Cart{}, err
	}

	newCart := models.Cart{
		CustomerId:   1,
		Total:        5,
		TotalSavings: 2,
		Status:       "OPEN",
	}

	err := db.Model(&models.Cart{}).Create(&newCart).Error
	if err != nil {
		return models.Cart{}, err
	}

	log.Printf("seedOneCart routine OK !!!")
	return newCart, nil
}

func SeedSingleItemDiscount(db *gorm.DB) (models.AppliedSingleItemDiscount, error) {
	if err := RefreshDiscountsTable(db); err != nil {
		return models.AppliedSingleItemDiscount{}, err
	}

	newDiscount := models.AppliedSingleItemDiscount{
		CartID:  1,
		Savings: 2.0,
	}

	newAppleDiscount := models.SingleItemDiscount{
		FruitID: 1,
		Count:   7,
		Name:    "APPLE10",
		Model: gorm.Model{
			ID: 1,
		},
	}
	err := db.Model(&models.SingleItemDiscount{}).Create(&newAppleDiscount).Error
	if err != nil {
		return models.AppliedSingleItemDiscount{}, err
	}
	err = db.Model(&models.AppliedSingleItemDiscount{}).Create(&newDiscount).Error
	if err != nil {
		return models.AppliedSingleItemDiscount{}, err
	}

	log.Printf("seedSingleItemDiscount routine OK !!!")
	return newDiscount, nil
}

func SeedOneCartItem(db *gorm.DB) (models.CartItem, error) {
	if err := RefreshCartItemTable(db); err != nil {
		return models.CartItem{}, err
	}

	newCartItem := models.CartItem{
		CartID:              1,
		FruitID:             1,
		Name:                "Apple",
		Quantity:            10,
		ItemTotal:           10,
		ItemDiscountedTotal: 0.0,
	}

	err := db.Model(&models.CartItem{}).Create(&newCartItem).Error
	if err != nil {
		return models.CartItem{}, err
	}

	log.Printf("seedOneCartItem routine OK !!!")
	return newCartItem, nil
}

func SeedOneCustomer(db *gorm.DB) (models.Customer, error) {
	if err := RefreshCustomerTable(db); err != nil {
		return models.Customer{}, err
	}
	if err := RefreshCartTable(db); err != nil {
		return models.Customer{}, err
	}

	newcart := models.Cart{
		Total:  0.0,
		Status: "OPEN",
	}

	customer := models.Customer{
		FirstName: "Rakesh",
		LastName:  "Mothukuri",
		LoginID:   "a",
		Cart:      newcart,
	}

	err := db.Model(&models.Customer{}).Create(&customer).Error
	if err != nil {
		log.Fatalf("cannot seed customers table: %v", err)
	}

	log.Printf("seedOneCustomer routine OK !!!")
	return customer, nil
}

func SeedFruits(db *gorm.DB) ([]models.Fruit, error) {
	var err error
	if err != nil {
		return nil, err
	}
	fruits := []models.Fruit{
		models.Fruit{
			Name:  "Apple",
			Price: 1.0,
		},
		models.Fruit{
			Name:  "Pear",
			Price: 1.0,
		},
		models.Fruit{
			Name:  "Banana",
			Price: 1.0,
		},
		models.Fruit{
			Name:  "Orange",
			Price: 1.0,
		},
	}
	for i, _ := range fruits {
		err := db.Model(&models.Fruit{}).Create(&fruits[i]).Error
		if err != nil {
			return []models.Fruit{}, err
		}
	}
	return fruits, nil
}

// possibility to load environment varables from a .env file and pass then to Database() function
/*err := godotenv.Load(os.ExpandEnv("../../.env"))
if err != nil {
	log.Fatalf("Error getting env %v\n", err)
}*/

/*
	TestDbDriver := os.Getenv("TestDbDriver")

		if TestDbDriver == "mysql" {
			DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("TestDbUser"), os.Getenv("TestDbPassword"), os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbName"))
			server.DB, err = gorm.Open(TestDbDriver, DBURL)
			if err != nil {
				fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
				log.Fatal("This is the error:", err)
			} else {
				fmt.Printf("We are connected to the %s database\n", TestDbDriver)
			}
		}
		if TestDbDriver == "postgres" {
			DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbUser"), os.Getenv("TestDbName"), os.Getenv("TestDbPassword"))
			server.DB, err = gorm.Open(TestDbDriver, DBURL)
			if err != nil {
				fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
				log.Fatal("This is the error:", err)
			} else {
				fmt.Printf("We are connected to the %s database\n", TestDbDriver)
			}
		}
*/
