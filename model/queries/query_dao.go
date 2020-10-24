package queries

import (
	"fmt"

	"github.com/olivere/elastic"
)

// Build meth
func (ch *EsQuery) Build() elastic.Query {
	boolQuery := elastic.NewBoolQuery()
	matchQueries := make([]elastic.Query,0)
	notMatchQueries := make([]elastic.Query,0)
	
	for _, equal := range ch.Equals {
		matchQueries = append(matchQueries, elastic.NewMatchQuery(equal.Field,equal.Value))
	}

	for _, notEqual := range ch.NotEquals {
		notMatchQueries = append(notMatchQueries, elastic.NewMatchQuery(fmt.Sprintf("%s.keyword",notEqual.Field),notEqual.Value))
	}

	boolQuery.Must(matchQueries...)
	boolQuery.MustNot(notMatchQueries...)

	return boolQuery
}
