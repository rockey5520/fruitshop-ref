package controllertests

import (
	"fruitshop/internal/controllers"
	"fruitshop/internal/testdb"

	"github.com/jinzhu/gorm"
)

var (
	server = &controllers.TooCommon{}
	db     *gorm.DB
)

func init() {
	// this is not recommended and is a temp workaround
	var err error
	db, err = testdb.Database()
	if err != nil {
		panic(err)
	}
	server.DB = db
}
