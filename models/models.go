package models

import (
	"time"

	_ "gorm.io/gorm"
)

type Customer struct {
	Id         int       `json:"id" gorm:"primaryKey;column:id;type:integer"`
	Name       string    `json:"name" form:"name" gorm:"type:varchar(255);not null;column:name"`
	Email      string    `json:"email" form:"email" gorm:"type:varchar(255);not null;unique;column:email"`
	Created_at time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP();column:created_at"`
	Updated_at time.Time `json:"updated_at" gorm:"autoUpdateTime:milli;default:CURRENT_TIMESTAMP();column:updated_at"`
}

type Product struct {
	Id         int       `json:"id" gorm:"primaryKey;column:id;type:integer"`
	Name       string    `json:"name" form:"name" gorm:"type:varchar(255);not null;column:name"`
	Price      float64   `json:"price" form:"price" gorm:"type:decimal(10,2);not null;column:price"`
	Stock      int       `json:"stock" form:"stock" gorm:"not null;default:0;column:stock;type:int"`
	Created_at time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP();column:created_at"`
	Updated_at time.Time `json:"updated_at" gorm:"autoUpdateTime:milli;default:CURRENT_TIMESTAMP();column:updated_at"`
}

type Order struct {
	Id          int       `json:"id" gorm:"primaryKey;column:id;type:integer"`
	Customer_id int       `json:"customer_id" form:"customer_id" gorm:"column:customer_id"`
	Product_id  int       `json:"product_id" form:"product_id" gorm:"column:product_id;type:integer"`
	Quantity    int       `json:"quantity" form:"quantity" gorm:"not null;type:integer"`
	Total       float64   `json:"total" gorm:"not null"`
	Created_at  time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	Updated_at  time.Time `json:"updated_at" gorm:"autoUpdateTime:milli;default:CURRENT_TIMESTAMP()"`
	Customer    Customer  `json:"customer" gorm:"foreignKey:Customer_id"`
	Product     Product   `json:"product" gorm:"foreignKey:Product_id"`
}
