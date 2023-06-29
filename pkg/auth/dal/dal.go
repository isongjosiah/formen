package dal

import (
	"context"
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"log"

	"statuarius/internal/config"
	"statuarius/internal/dal"
)

type AuthDAL struct {
	DB *bun.DB
}

func connectDB(cfg *config.AuthConfig) (*bun.DB, error) {
	pgDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.DatabaseURL)))
	db := bun.NewDB(pgDB, pgdialect.New())

	if cfg.DebugDatabase {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	return db, db.Ping()
}

func (ad AuthDAL) createTables() error {
	dbModels := []interface{}{}

	for _, m := range dbModels {
		_, err := ad.DB.NewCreateTable().
			IfNotExists().
			Model(m).Exec(context.TODO())
		if err != nil {
			return err
		}
	}

	return nil
}

func (ad AuthDAL) migrateSchema(cfg *config.AuthConfig) error {
	m, err := migrate.New("file://pkg/auth/dal/migrations", cfg.DatabaseURL)
	if err != nil {
		return errors.Wrap(err, "failed to start migration")
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return errors.Wrap(err, "failed to process up migration")
		}
	}
	return nil

}

func New(cfg *config.AuthConfig) dal.DataAccessLayerInterface {

	bunDB, connErr := connectDB(cfg)
	if connErr != nil {
		log.Fatal("[DB]: failed to connect DB")
	}

	aDal := &AuthDAL{
		DB: bunDB,
	}

	createErr := aDal.createTables()
	if createErr != nil {
		log.Fatal("[DB]: failed to create tables")
	}

	return aDal
}

func (ad AuthDAL) Create(model interface{}, tx *bun.Tx) error {
	return nil
}

func (ad AuthDAL) Fetch(model interface{}, conditions []string) (interface{}, error) {
	return nil, nil
}

func (ad AuthDAL) Update(model interface{}, tx *bun.Tx, conditions []string) error {
	return nil
}

func (ad AuthDAL) Delete(model interface{}, tx *bun.Tx, conditions []string) error {
	return nil
}

func (ad AuthDAL) RawQuery(model interface{}, query string, tx *bun.Tx) error {
	return nil
}
