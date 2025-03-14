// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableNetworkGraphConfigsStmt holds the create statement for table `network_graph_configs`.
	CreateTableNetworkGraphConfigsStmt = &postgres.CreateStmts{
		GormModel: (*NetworkGraphConfigs)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// NetworkGraphConfigsSchema is the go schema for table `network_graph_configs`.
	NetworkGraphConfigsSchema = func() *walker.Schema {
		schema := GetSchemaForTable("network_graph_configs")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.NetworkGraphConfig)(nil)), "network_graph_configs")
		RegisterTable(schema, CreateTableNetworkGraphConfigsStmt)
		return schema
	}()
)

const (
	NetworkGraphConfigsTableName = "network_graph_configs"
)

// NetworkGraphConfigs holds the Gorm model for Postgres table `network_graph_configs`.
type NetworkGraphConfigs struct {
	Id         string `gorm:"column:id;type:varchar;primaryKey"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
