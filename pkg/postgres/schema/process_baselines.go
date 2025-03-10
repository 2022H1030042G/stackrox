// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/search"
)

var (
	// CreateTableProcessBaselinesStmt holds the create statement for table `process_baselines`.
	CreateTableProcessBaselinesStmt = &postgres.CreateStmts{
		GormModel: (*ProcessBaselines)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ProcessBaselinesSchema is the go schema for table `process_baselines`.
	ProcessBaselinesSchema = func() *walker.Schema {
		schema := GetSchemaForTable("process_baselines")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.ProcessBaseline)(nil)), "process_baselines")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_PROCESS_BASELINES, "processbaseline", (*storage.ProcessBaseline)(nil)))
		RegisterTable(schema, CreateTableProcessBaselinesStmt)
		return schema
	}()
)

const (
	ProcessBaselinesTableName = "process_baselines"
)

// ProcessBaselines holds the Gorm model for Postgres table `process_baselines`.
type ProcessBaselines struct {
	Id              string `gorm:"column:id;type:varchar;primaryKey"`
	KeyDeploymentId string `gorm:"column:key_deploymentid;type:varchar"`
	KeyClusterId    string `gorm:"column:key_clusterid;type:varchar;index:sac_filter,type:btree"`
	KeyNamespace    string `gorm:"column:key_namespace;type:varchar;index:sac_filter,type:btree"`
	Serialized      []byte `gorm:"column:serialized;type:bytea"`
}
