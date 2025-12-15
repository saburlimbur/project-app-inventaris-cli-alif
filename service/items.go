package service

import (
	"errors"
	"fmt"
	"inventory-cli/dto"
	"inventory-cli/model"
	"inventory-cli/repository"
	"inventory-cli/utils"
)

type ItemsService interface {
	CreateItems(itm *model.ItemsModel) error
	Lists() ([]*model.ItemsModel, error)
	GetItemByID(id int) (*model.ItemsModel, error)
	DeleteItem(id int) error
	SearchItems(keyword string) ([]*model.ItemsModel, error)
	UpdateItem(itm *model.ItemsModel) error
	GetItemsNeedReplacement() ([]*model.ItemsModel, error)

	GetInvestmentSummary() (*dto.InvestmentSummaryResponse, error)
	GetInvestmentDetail(id int) (*dto.InvestmentDetailResponse, error)
}

type itemsService struct {
	Repo repository.ItemsRepository
}

func NewServiceItems(repo repository.ItemsRepository) ItemsService {
	return &itemsService{
		Repo: repo,
	}
}

func (svc *itemsService) CreateItems(itm *model.ItemsModel) error {
	return svc.Repo.Create(itm)
}

func (svc *itemsService) Lists() ([]*model.ItemsModel, error) {
	return svc.Repo.FindAll()
}

func (svc *itemsService) GetItemByID(id int) (*model.ItemsModel, error) {
	return svc.Repo.FindById(id)
}

func (svc *itemsService) DeleteItem(id int) error {
	if id <= 0 {
		return utils.ErrInvalidItemID
	}

	exist, err := svc.Repo.FindById(id)
	if err != nil {
		return fmt.Errorf("service: failed to check item: %w", err)
	}

	if exist == nil {
		return utils.ErrItemNotFound
	}

	if err := svc.Repo.Delete(id); err != nil {
		return fmt.Errorf("service: failed to delete item: %w", err)
	}

	return nil
}

func (svc *itemsService) SearchItems(keyword string) ([]*model.ItemsModel, error) {
	return svc.Repo.SearchByName(keyword)
}

func (svc *itemsService) UpdateItem(itm *model.ItemsModel) error {
	return svc.Repo.Update(itm)
}

func (svc *itemsService) GetItemsNeedReplacement() ([]*model.ItemsModel, error) {
	return svc.Repo.FindNeedReplacement()
}

func (svc *itemsService) GetInvestmentSummary() (*dto.InvestmentSummaryResponse, error) {
	items, err := svc.Repo.FindAll()
	if err != nil {
		return nil, err
	}

	var totalInitial float64
	var totalCurrent float64

	for _, itm := range items {
		totalInitial += itm.Price

		years := utils.YearsUsed(itm.PurchaseDate)
		currentValue := utils.DecliningBalance(itm.Price, years)

		totalCurrent += currentValue
	}

	return &dto.InvestmentSummaryResponse{
		TotalInitialValue: totalInitial,
		TotalCurrentValue: totalCurrent,
	}, nil
}

func (svc *itemsService) GetInvestmentDetail(id int) (*dto.InvestmentDetailResponse, error) {
	itm, err := svc.Repo.FindById(id)
	if err != nil {
		return nil, err
	}
	if itm == nil {
		return nil, errors.New("item not found")
	}

	years := utils.YearsUsed(itm.PurchaseDate)
	currentValue := utils.DecliningBalance(itm.Price, years)

	return &dto.InvestmentDetailResponse{
		ItemID:       itm.ID,
		ItemName:     itm.Name,
		InitialValue: itm.Price,
		CurrentValue: currentValue,
		Depreciation: itm.Price - currentValue,
		YearsUsed:    years,
	}, nil
}
