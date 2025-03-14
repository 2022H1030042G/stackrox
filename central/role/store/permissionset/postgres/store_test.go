// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/suite"
)

type PermissionSetsStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
	store       Store
	testDB      *pgtest.TestPostgres
}

func TestPermissionSetsStore(t *testing.T) {
	suite.Run(t, new(PermissionSetsStoreSuite))
}

func (s *PermissionSetsStoreSuite) SetupSuite() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(env.PostgresDatastoreEnabled.EnvVar(), "true")

	if !env.PostgresDatastoreEnabled.BooleanSetting() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.Pool)
}

func (s *PermissionSetsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE permission_sets CASCADE")
	s.T().Log("permission_sets", tag)
	s.NoError(err)
}

func (s *PermissionSetsStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
	s.envIsolator.RestoreAll()
}

func (s *PermissionSetsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	permissionSet := &storage.PermissionSet{}
	s.NoError(testutils.FullInit(permissionSet, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundPermissionSet, exists, err := store.Get(ctx, permissionSet.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundPermissionSet)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, permissionSet))
	foundPermissionSet, exists, err = store.Get(ctx, permissionSet.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(permissionSet, foundPermissionSet)

	permissionSetCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(1, permissionSetCount)
	permissionSetCount, err = store.Count(withNoAccessCtx)
	s.NoError(err)
	s.Zero(permissionSetCount)

	permissionSetExists, err := store.Exists(ctx, permissionSet.GetId())
	s.NoError(err)
	s.True(permissionSetExists)
	s.NoError(store.Upsert(ctx, permissionSet))
	s.ErrorIs(store.Upsert(withNoAccessCtx, permissionSet), sac.ErrResourceAccessDenied)

	foundPermissionSet, exists, err = store.Get(ctx, permissionSet.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(permissionSet, foundPermissionSet)

	s.NoError(store.Delete(ctx, permissionSet.GetId()))
	foundPermissionSet, exists, err = store.Get(ctx, permissionSet.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundPermissionSet)
	s.ErrorIs(store.Delete(withNoAccessCtx, permissionSet.GetId()), sac.ErrResourceAccessDenied)

	var permissionSets []*storage.PermissionSet
	for i := 0; i < 200; i++ {
		permissionSet := &storage.PermissionSet{}
		s.NoError(testutils.FullInit(permissionSet, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		permissionSets = append(permissionSets, permissionSet)
	}

	s.NoError(store.UpsertMany(ctx, permissionSets))

	permissionSetCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(200, permissionSetCount)
}
