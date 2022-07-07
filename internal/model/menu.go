package model

import "github.com/google/uuid"

type Menu struct {
	ID             uuid.UUID `json:"id" gorm:"primarykey;autoIncrement"`
	MenuKategoriID int       `json:"menu_kategori_id" gorm:"not null"`
	Name           string    `json:"name" gorm:"size:200;unique;not null"`
	Description    string    `json:"description" gorm:"not null"`
	ImageUrl       string    `json:"image_url" gorm:"not null"`
	Price          float64   `json:"price" gorm:"not null"`
	InStock        int64     `json:"in_stock" gorm:"not null"`
}

// func (m *Menu) BeforeCreate(tx *gorm.DB) (err error) {
// 	m.CreatedAt = time.Now()
// 	return
// }

// BeforeUpdate is a method for struct User
// gorm call this method before they execute query
// func (p *Menu) BeforeUpdate(tx *gorm.DB) (err error) {
// 	p.UpdatedAt = time.Now()
// 	return
// }
