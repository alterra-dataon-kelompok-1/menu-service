package seeder

import (
	"log"

	"github.com/alterra-kelompok-1/menu-service/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func menuTableSeeder(conn *gorm.DB) {
	menu := []model.Menu{
		{
			ID:             uuid.New(),
			MenuKategoriID: 2,
			Name:           "Mie goreng",
			Description:    "Mie yang terbuat dari gandum, dan dimasak dengan cara digoreng",
			ImageUrl:       "http://s3-ap-southeast-3.amazonaws.com/bucket/afa08a08a90s890",
			Price:          10000,
			InStock:        3,
		},
		{
			ID:             uuid.New(),
			MenuKategoriID: 1,
			Name:           "Nasi rebus",
			Description:    "Nasi yang terbuat dari gabah, dan dimasak dengan cara direbus",
			ImageUrl:       "http://s3-ap-southeast-3.amazonaws.com/bucket/nokrnoifgj09fuh90",
			Price:          25000,
			InStock:        9,
		},
		{
			ID:             uuid.New(),
			MenuKategoriID: 55,
			Name:           "Teriyaki",
			Description:    "Ikan teri yang dibalur dengan saus rahasia",
			ImageUrl:       "http://s3-ap-southeast-3.amazonaws.com/bucket/ikjfioejaia90fu0a9",
			Price:          40000,
			InStock:        4,
		},
	}

	if err := conn.Create(&menu).Error; err != nil {
		log.Printf("cannot seed data, %s", err)
	} else {
		log.Println("seed menu data success")
	}
}
