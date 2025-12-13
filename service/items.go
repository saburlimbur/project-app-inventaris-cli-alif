package service

import (
	"inventory-cli/model"
	"inventory-cli/repository"
)

type ItemsService interface {
	CreateItems(itm *model.ItemsModel) error
	Lists() ([]*model.ItemsModel, error)
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
