package sqlx

import (
	"context"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/Masterminds/squirrel"
)

func RetrieveAppliedMigrations(
	ctx context.Context,
	logger lager.Logger,
	conn *DB,
	tableName string,
) (map[int]AppliedMigration, error) {
	rows, err := squirrel.Select("version", "name", "applied_at").
		From(tableName).
		RunWith(conn).
		QueryContext(ctx)

	if err != nil {
		logger.Error(failedToQueryMigrations, err)
		return nil, err
	}

	defer rows.Close()
	var (
		version   int
		name      string
		appliedAt time.Time
	)

	versions := make(map[int]AppliedMigration)
	for rows.Next() {
		err = rows.Scan(&version, &name, &appliedAt)
		if err != nil {
			logger.Error(failedToParseAppliedMigration, err)

			return nil, err
		}
		versions[version] = AppliedMigration{
			Version:   version,
			Name:      name,
			AppliedAt: appliedAt,
		}
	}

	err = rows.Err()
	if err != nil {
		logger.Error(failedToQueryMigrations, err)
		return nil, err
	}

	return versions, nil
}
