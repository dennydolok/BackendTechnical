package databases

import (
	"dennydolok/experimentals-1/config"
	"dennydolok/experimentals-1/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(conf config.Config) *gorm.DB {
	// make database connection string
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)
	//open connection
	DB, err := gorm.Open(mysql.Open(connection))

	//error checking
	if err != nil {
		fmt.Println("Error! cannot connect to database : ", err)
	}

	//Migrate Database models
	DB.AutoMigrate(&models.Customer{})
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.Product{})
	return DB
}
