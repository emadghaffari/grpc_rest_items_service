package items

import (
	"github.com/emadghaffari/grpc_rest_items_service/model/queries"
	"github.com/emadghaffari/res_errors/errors"
)

type itemInterface interface{
	Save() errors.ResError
	IndexExists(string) errors.ResError
	Get() errors.ResError
	Search(query queries.EsQuery) ([]Item,errors.ResError)
	Delete() errors.ResError
	Update() errors.ResError
}

// Item struct
type Item struct {
	ID                string      `json:"id"`
	Seller            int64       `json:"seller"`
	Title             string      `json:"title"`
	Description       Description `json:"description"`
	Pictures          []Picture   `json:"pictures"`
	Video             string      `json:"video"`
	Price             float32     `json:"price"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Status            string      `json:"status"`
}

// Description struct
type Description struct {
	PlainText string `json:"plain_text"`
	HTML      string `json:"html"`
}

// Picture struct
type Picture struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}
