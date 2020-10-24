package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/emadghaffari/res_errors/errors"
	"github.com/emadghaffari/res_errors/logger"
	"github.com/olivere/elastic"
)


var(
	// Client var
	Client esClientInterface = &esClient{}
)

type esClientInterface interface{
	SetClient(*elastic.Client)

	Index(string,string,interface{}) (*elastic.IndexResponse,errors.ResError)
	IndexExists(string) (errors.ResError)
	Delete(string,string,string) (*elastic.DeleteResponse, errors.ResError)
	Get(string,string,string) (*elastic.GetResult, errors.ResError)
	Search(string,elastic.Query) (*elastic.SearchResult,errors.ResError)
	Update(string,string,string, *elastic.Script) (*elastic.UpdateResponse, errors.ResError)
}
type esClient struct{
	client *elastic.Client
}

// SetClient method
// for set new client for elk
func (es *esClient) SetClient(client *elastic.Client){
	es.client = client
}

// Init func
// Init the elasticsearch for first time we want to use
func Init()  {
	logger := logger.GetLogger()
	client,err := elastic.NewClient(
		elastic.SetURL("http://elasticsearch:9200"),
		elastic.SetBasicAuth("elastic","changeme"),
		elastic.SetHealthcheckInterval(25 * time.Second),
		elastic.SetErrorLog(logger),
		elastic.SetInfoLog(logger),
	)
	if err != nil {
		panic(err)
	}
	Client.SetClient(client)
}

// Index meth
// index,docType and doc for Index
func (es *esClient) Index(index string,docType string,doc interface{}) (*elastic.IndexResponse,errors.ResError){
	ctx := context.Background()
	elk,err := es.client.
	Index().
	Index(index).
	BodyJson(doc).
	Type(docType).
	Do(ctx)
	if err != nil {
		logger.Error("error in index esClient", err)
		return nil, errors.HandlerInternalServerError("internal ELK error in Index", err)
	}
	return elk, nil
}

// IndexExists meth
// check index is exists or not! with index name
func (es *esClient) IndexExists(index string) errors.ResError {
	// Check if the index called "twitter" exists
	exists, err := es.client.IndexExists(index).Do(context.Background())
	if err != nil {
		logger.Error(fmt.Sprintf("invalid index: %s", index), err)
		return errors.HandlerInternalServerError(fmt.Sprintf("invalid index: %s", index), err)
	}
	if !exists {
		logger.Error(fmt.Sprintf("Index does not exist yet: %s", index), err)
		return errors.HandlerInternalServerError(fmt.Sprintf("Index does not exist yet: %s", index), err)
	}
	return nil
}

// Delete meth
// index,docType and id for Delete a row from elk
func (es *esClient) Delete(index string,docType string,id string) (*elastic.DeleteResponse, errors.ResError){
	elk,err := es.client.Delete().
	Id(id).
	Index(index).
	Type(docType).
	Do(context.Background())
	if err != nil {
		logger.Error(fmt.Sprintf("error in Delete from elastic %s", id), err)
		return nil, errors.HandlerInternalServerError("error in Delete from elastic", err)
	}
	return elk,nil
}

// Get meth
// index,doctype and id for get
func (es *esClient) Get(index string,docType string,id string) (*elastic.GetResult, errors.ResError){
	elk,err := es.client.Get().
	Id(id).
	Index(index).
	Type(docType).
	Do(context.Background())
	if err != nil {
		logger.Error(fmt.Sprintf("error in Get data from elastic %s", id), err)
		return nil, errors.HandlerInternalServerError("error in get from elastic", err)
	}
	return elk,nil
}

// Search meth
// index and query for search
func (es *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult,errors.ResError){
	elk,err := es.client.Search(index).
	Query(query).
	RestTotalHitsAsInt(true).
	Do(context.Background())
	if err != nil {
		logger.Error(fmt.Sprintf("error in Search data from elastic %v", query), err)
		return nil, errors.HandlerInternalServerError("error in Search from elastic", err)
	}
	return elk,nil
}

// Update meth
// index,Type,id and script query for Update
func (es *esClient) Update(index string,docType string,id string, script *elastic.Script) (*elastic.UpdateResponse, errors.ResError){
	elk,err := es.client.Update().
	Index(index).
	Type(docType).
	Id(id).
	Script(script).
	Do(context.Background())
	if err != nil {
		logger.Error(fmt.Sprintf("error in Get data from elastic %s", id), err)
		return nil, errors.HandlerInternalServerError("error in get from elastic", err)
	}
	return elk,nil
}