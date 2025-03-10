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

type GroupsStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
	store       Store
	testDB      *pgtest.TestPostgres
}

func TestGroupsStore(t *testing.T) {
	suite.Run(t, new(GroupsStoreSuite))
}

func (s *GroupsStoreSuite) SetupSuite() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(env.PostgresDatastoreEnabled.EnvVar(), "true")

	if !env.PostgresDatastoreEnabled.BooleanSetting() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.Pool)
}

func (s *GroupsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE groups CASCADE")
	s.T().Log("groups", tag)
	s.NoError(err)
}

func (s *GroupsStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
	s.envIsolator.RestoreAll()
}

func (s *GroupsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	group := &storage.Group{}
	s.NoError(testutils.FullInit(group, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundGroup, exists, err := store.Get(ctx, group.GetProps().GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundGroup)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, group))
	foundGroup, exists, err = store.Get(ctx, group.GetProps().GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(group, foundGroup)

	groupCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(1, groupCount)
	groupCount, err = store.Count(withNoAccessCtx)
	s.NoError(err)
	s.Zero(groupCount)

	groupExists, err := store.Exists(ctx, group.GetProps().GetId())
	s.NoError(err)
	s.True(groupExists)
	s.NoError(store.Upsert(ctx, group))
	s.ErrorIs(store.Upsert(withNoAccessCtx, group), sac.ErrResourceAccessDenied)

	foundGroup, exists, err = store.Get(ctx, group.GetProps().GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(group, foundGroup)

	s.NoError(store.Delete(ctx, group.GetProps().GetId()))
	foundGroup, exists, err = store.Get(ctx, group.GetProps().GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundGroup)
	s.ErrorIs(store.Delete(withNoAccessCtx, group.GetProps().GetId()), sac.ErrResourceAccessDenied)

	var groups []*storage.Group
	for i := 0; i < 200; i++ {
		group := &storage.Group{}
		s.NoError(testutils.FullInit(group, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		groups = append(groups, group)
	}

	s.NoError(store.UpsertMany(ctx, groups))
	allGroup, err := store.GetAll(ctx)
	s.NoError(err)
	s.ElementsMatch(groups, allGroup)

	groupCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(200, groupCount)
}
