package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID			uint			`gorm:"primaryKey" json:"id"`
	Name      	string         	`gorm:"unique;not null" json:"name"`
	Users     	[]User         	`gorm:"foreignKey:RoleID" json:"-"`
	CreatedAt 	time.Time		`json:"created_at"`
	UpdatedAt 	time.Time		`json:"updated_at"`
	DeletedAt 	gorm.DeletedAt 	`gorm:"index" json:"-"`
}

type User struct {
	ID        	uint           	`gorm:"primaryKey" json:"id"`
	Name      	string         	`gorm:"not null" json:"name"`
	Email     	string         	`gorm:"unique;not null" json:"email"`
	Password  	string         	`gorm:"not null" json:"password"`
	RoleID    	uint           	`gorm:"not null" json:"role_id"`
	Role      	Role           	`json:"role"`
	CreatedAt 	time.Time		`json:"created_at"`
	UpdatedAt 	time.Time		`json:"updated_at"`
	DeletedAt 	gorm.DeletedAt 	`gorm:"index" json:"-"`
}

type Product struct {
	ID        	uint           	`gorm:"primaryKey" json:"id"`
	Name      	string         	`gorm:"not null" json:"name"`
	Price     	float64        	`gorm:"not null" json:"price"`
	Description	string			`gorm:"not null" json:"description"`
	Stock     	int            	`gorm:"not null" json:"stock"`
	CreatedAt 	time.Time		`json:"created_at"`
	UpdatedAt 	time.Time		`json:"updated_at"`
	DeletedAt 	gorm.DeletedAt 	`gorm:"index" json:"-"`
}