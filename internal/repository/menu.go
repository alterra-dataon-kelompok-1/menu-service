package repository

import (
	"context"
	"strings"

	"github.com/alterra-kelompok-1/menu-service/internal/dto"
	"github.com/alterra-kelompok-1/menu-service/internal/model"
	"gorm.io/gorm"
)

type Menu interface {
	Create(ctx context.Context, data model.Menu) error
	Find(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Menu, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID string) (model.Menu, error)
	Update(ctx context.Context, ID string, data map[string]interface{}) error
	Delete(ctx context.Context, ID string) error
}

type menu struct {
	Db *gorm.DB
}

func NewMenu(db *gorm.DB) *menu {
	return &menu{
		db,
	}
}

func (p *menu) Create(ctx context.Context, data model.Menu) error {
	return p.Db.WithContext(ctx).Model(&model.Menu{}).Create(&data).Error
}

func (p *menu) Find(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Menu, *dto.PaginationInfo, error) {
	var menus []model.Menu
	var count int64

	// query := p.Db.WithContext(ctx).Model(&model.Menu{})
	query := p.Db.Model(&model.Menu{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ?  ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)
	// fmt.Println(limit, offset)
	// fmt.Println(payload.Search)

	result := query.Limit(limit).Offset(offset).Find(&menus)
	err := result.Error
	// fmt.Println(result.RowsAffected)
	// fmt.Printf("%v", payload)

	return menus, dto.CheckInfoPagination(paginate, count), err
}

func (p *menu) FindByID(ctx context.Context, ID string) (model.Menu, error) {

	var data model.Menu
	err := p.Db.WithContext(ctx).Model(&data).Where("id = ?", ID).First(&data).Error

	return data, err
}

func (p *menu) Update(ctx context.Context, ID string, data map[string]interface{}) error {

	err := p.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.Menu{}).Updates(data).Error
	return err

	// var menu model.MenuMenu

	// err := p.Db.WithContext(ctx).Model(&data).Where("id = ?", ID).First(&data).Error

	// if data.Name != nil {
	// menu.Name = data.Name
	// }

	// menu.Stock = data.Stock
	// menu.Description = data.Description

	// err = p.Db.Save(&menu).Error

	// return nil
}

func (p *menu) Delete(ctx context.Context, ID string) error {

	err := p.Db.WithContext(ctx).Where("id = ?", ID).Delete(&model.Menu{}).Error
	return err
}
