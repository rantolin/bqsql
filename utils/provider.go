/*
Copyright © 2024 Roberto Antolin <rantolin DOT geo AT gmail DOT com>
*/
package formats

import (
	"cloud.google.com/go/bigquery"
)

// RowProvider abstracts the iteration and schema retrieval from BigQuery results.
type RowProvider interface {
	Next(dst interface{}) error
	Schema() bigquery.Schema
}

// BigQueryRowProvider wraps bigquery.RowIterator to satisfy the RowProvider interface.
type BigQueryRowProvider struct {
	It *bigquery.RowIterator
}

func (b *BigQueryRowProvider) Next(dst interface{}) error {
	return b.It.Next(dst)
}

func (b *BigQueryRowProvider) Schema() bigquery.Schema {
	return b.It.Schema
}
