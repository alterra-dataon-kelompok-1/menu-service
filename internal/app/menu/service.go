package menu

import (
	"context"

	"github.com/alterra-kelompok-1/menu-service/internal/dto"
	"github.com/alterra-kelompok-1/menu-service/internal/factory"
	"github.com/alterra-kelompok-1/menu-service/internal/model"
	"github.com/alterra-kelompok-1/menu-service/internal/repository"
	"github.com/alterra-kelompok-1/menu-service/pkg/constant"
	res "github.com/alterra-kelompok-1/menu-service/pkg/util/response"
)

type service struct {
	MenuRepository repository.Menu
}

type Service interface {
	Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.Menu], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.Menu, error)
	Create(ctx context.Context, payload *dto.CreateMenuRequest) (string, error)
	Update(ctx context.Context, ID string, payload *dto.UpdateMenuRequest) (string, error)
	Delete(ctx context.Context, ID string) (*model.Menu, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		MenuRepository: f.MenuRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.Menu], error) {

	Menus, info, err := s.MenuRepository.Find(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := new(dto.SearchGetResponse[model.Menu])
	result.Datas = Menus
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.Menu, error) {

	data, err := s.MenuRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}

func (s *service) Create(ctx context.Context, payload *dto.CreateMenuRequest) (string, error) {

	var menu = model.Menu{
		Name:        payload.Name,
		InStock:     payload.Stock,
		Description: payload.Description,
	}

	err := s.MenuRepository.Create(ctx, menu)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Update(ctx context.Context, ID string, payload *dto.UpdateMenuRequest) (string, error) {

	var data = make(map[string]interface{})

	if payload.Name != nil {
		data["name"] = payload.Name
	}
	if payload.Stock != nil {
		data["stock"] = payload.Stock
	}
	if payload.Description != nil {
		data["description"] = payload.Description
	}

	err := s.MenuRepository.Update(ctx, ID, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Delete(ctx context.Context, ID string) (*model.Menu, error) {

	data, err := s.MenuRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	err = s.MenuRepository.Delete(ctx, ID)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil

}
