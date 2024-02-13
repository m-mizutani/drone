package bq

import (
	"context"

	"github.com/m-mizutani/drone/pkg/domain/interfaces"
)

type Mock struct {
	InsertedData []any
}

var _ interfaces.BigQuery = &Mock{}

func NewMock() *Mock {
	return &Mock{}
}

func (x *Mock) CreateOrUpdateSchema(ctx context.Context, tableName string, target any) error {
	return nil
}

func (x *Mock) Insert(ctx context.Context, tableName string, data any) error {
	x.InsertedData = append(x.InsertedData, data)
	return nil
}
