package postgres

import (
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"os/exec"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	carsTable   = "cars"
	peopleTable = "people"
)

func NewDB(cfg Config) (*sqlx.DB, error) {
	if err := MakeMigrationUp(cfg); err != nil {
		return nil, err
	}
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	slog.Info("Success connect to DB")
	return db, err
}

func MakeMigrationUp(cfg Config) error {
	db := fmt.Sprintf("%s://%s:%s@%s:%s/postgres?sslmode=%s", cfg.Username, cfg.DBName, cfg.Password, cfg.Host, cfg.Port, cfg.SSLMode)
	slog.Debug(db)
	cmd := exec.Command("migrate", "-path", "./schema/", "-database", db, "up")

	output, err := cmd.CombinedOutput()
	if err != nil {
		slog.Info("Can't use migration")
		return err
	}

	slog.Info("migrations: ", output)
	return err
}
