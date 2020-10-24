package services

import (
	"github.com/emadghaffari/grpc_rest_items_service/model/items"
	"github.com/emadghaffari/grpc_rest_items_service/model/queries"
	"github.com/emadghaffari/res_errors/errors"
)

var (
	// ItemService var from itemsServiceInterface interface
	ItemService itemsServiceInterface = &itemService{}
)

type itemsServiceInterface interface {
	Get(string) (*items.Item, errors.ResError)
	Create(items.Item) (*items.Item, errors.ResError)
	Search(queries.EsQuery) ([]items.Item, errors.ResError)
	Delete(string) (errors.ResError)

}

type itemService struct{}

func (s *itemService) Get(id string) (*items.Item, errors.ResError) {
	item := items.Item{ID: id}
	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) Create(item items.Item) (*items.Item, errors.ResError) {
	if err := item.Save(); err != nil {
		return nil, err
	}

	return &item, nil

}
func (s *itemService) Search(item queries.EsQuery) ([]items.Item, errors.ResError) {
	dao := items.Item{}
	result, err := dao.Search(item)
	if  err != nil {
		return nil, err
	}

	return result, nil

}

func (s *itemService) Delete(id string) (errors.ResError) {
	item := items.Item{ID: id}
	if err := item.Delete(); err != nil {
		return err
	}
	return nil
}
