package db

import (
	"database/sql"

	"github.com/helen/lumen5_miniurl/config"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type DB struct {
	*sql.DB
}

func (db *DB) Connect(cfg config.Config) error {
	dB, err := sql.Open("postgres", cfg.GetPostgresURL())
	if err != nil {
		log.Error().Err(err).Msgf("db.Connect - failed to open connection to DB %s", cfg.GetPostgresURL())
		return err
	}

	if err = dB.Ping(); err != nil {
		log.Error().Err(err).Msg("db.Connect - failed to ping db")
		return err
	}

	log.Info().Msg("Successfully connected to db")
	db.DB = dB
	return nil
}

// Todo: move to migrations
func (db *DB) SeedDB() error {
	log.Debug().Msg("Starting to seed db")
	sqlStatement := `
		DROP TABLE IF EXISTS urls;
		DROP TABLE IF EXISTS users;
		
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			password CHAR(60) NOT NULL
		);
		
		CREATE TABLE IF NOT EXISTS urls (
			id BIGSERIAL PRIMARY KEY,
			redirect_url TEXT NOT NULL,
			user_id UUID,

			CONSTRAINT fk_urls_user_id FOREIGN KEY (user_id)
				REFERENCES users (id) MATCH SIMPLE
				ON UPDATE CASCADE ON DELETE NO ACTION
		);

		CREATE INDEX IF NOT EXISTS idx_urls_user_id ON urls(user_id);
	`

	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Error().Err(err).Msg("db.SeedDB - failed to execute query")
		return err
	}

	log.Debug().Msg("Successfully seeded db")
	return nil
}
