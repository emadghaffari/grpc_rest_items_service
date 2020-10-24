package items

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/emadghaffari/grpc_rest_items_service/databases/elasticsearch"
	"github.com/emadghaffari/grpc_rest_items_service/model/queries"
	"github.com/emadghaffari/res_errors/errors"
)

const (
	indexES = "items"
	docType = "_doc"
)



// Save meth
// save(store) new doc to elastic DB
func (i *Item) Save() errors.ResError{
	result, err := elasticsearch.Client.Index(indexES,docType,i)
	if err != nil {
		return err
	}
	i.ID = result.Id

	return nil
}

// IndexExists meth
// check index exists or not
func (i *Item) IndexExists(index string) errors.ResError{
	return elasticsearch.Client.IndexExists(index)
}
// Get meth
// get single value from elastic
func (i *Item) Get() errors.ResError{
	itemID := i.ID
	result,err := elasticsearch.Client.Get(indexES,docType,i.ID)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return errors.HandlerNotFoundError(fmt.Sprintf("item not found %s", i.ID))
		}
		return err
	}

	if !result.Found{
		return errors.HandlerNotFoundError(fmt.Sprintf("item not found %s", i.ID))
	}

	bytes,marshalError := result.Source.MarshalJSON()
	if marshalError != nil {
		return errors.HandlerInternalServerError(fmt.Sprintf("error in MarshalJSON from elasticsearch DB %s", i.ID), marshalError)
	}
	if err := json.Unmarshal(bytes,&i); err != nil {
		return errors.HandlerInternalServerError(fmt.Sprintf("error in unmarshal data from Source.MarshalJSON elasticsearch %s", i.ID), err)
	}
	i.ID = itemID
	return nil
}

// Search meth
// search and return a list of matched elements
func (i *Item) Search(query queries.EsQuery) ([]Item,errors.ResError){
	result,err := elasticsearch.Client.Search(indexES,query.Build())
	if err != nil{
		return nil,errors.HandlerInternalServerError(fmt.Sprintf("error in Search from DB %s", i.ID), err)
	}
	items := make([]Item,result.TotalHits())
	for index,response := range result.Hits.Hits {
		bytes,err := response.Source.MarshalJSON()
		if err != nil {
			return nil, errors.HandlerInternalServerError(fmt.Sprintf("error in MarshalJSON from elasticsearch DB %s", i.ID), err)
		}
		var item Item
		if err := json.Unmarshal(bytes,&item); err != nil {
			return nil, errors.HandlerInternalServerError(fmt.Sprintf("error in unmarshal data from Source.MarshalJSON elasticsearch %s", i.ID), err)
		}
		item.ID = response.Id
		items[index] = item
	}

	if len(items) == 0 {
		return nil,errors.HandlerInternalServerError(fmt.Sprintf("no items found by filter "), nil)
	}

	return items,nil
}

// Delete meth
// Delete from elastic DB
func (i *Item) Delete() errors.ResError{
	result,err := elasticsearch.Client.Delete(indexES,docType,i.ID)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return errors.HandlerNotFoundError(fmt.Sprintf("item not found %s", i.ID))
		}
		return err
	}
	if result.Shards.Successful < 0 {
		return errors.HandlerNotFoundError(fmt.Sprintf("cannot delete item successfuly %s", i.ID))
	}

	return nil
}

// Update meth
// update value in elastic DB
func (i *Item) Update() errors.ResError{
	return nil
}