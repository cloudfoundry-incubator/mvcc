package db

import (
	"context"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/perm/pkg/api/repos"
)

func (s *DataService) HasPermission(
	ctx context.Context,
	logger lager.Logger,
	query repos.HasPermissionQuery,
) (bool, error) {
	return hasPermission(ctx, logger.Session("data-service"), s.conn, query)
}

func (s *DataService) ListResourcePatterns(
	ctx context.Context,
	logger lager.Logger,
	query repos.ListResourcePatternsQuery,
) ([]string, error) {
	return listResourcePatterns(ctx, logger.Session("data-service"), s.conn, query)
}
