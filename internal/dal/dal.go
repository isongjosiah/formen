package dal

import (
	"github.com/uptrace/bun"
)

type DataAccessLayerInterface interface {
	Create(model interface{}, tx *bun.Tx) error
	Fetch(model interface{}, conditions []string) (interface{}, error)
	Update(model interface{}, tx *bun.Tx, conditions []string) error
	Delete(model interface{}, tx *bun.Tx, conditions []string) error
	RawQuery(response interface{}, query string, tx *bun.Tx) error
}
