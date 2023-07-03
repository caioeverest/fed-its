package database

import (
	"fmt"
	"strings"
	"time"

	"github.com/caioeverest/fed-its/internal/config"
	"github.com/caioeverest/fed-its/internal/logger"
	"github.com/samber/lo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
	log *logger.Logger
	cfg *config.Config
}

// New builds a database that will be used by the application.
// It will be available to all of the application's dependencies.
func New(cfg *config.Config, log *logger.Logger) (*Database, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn(cfg)}), &gorm.Config{
		Logger: gormLogger.New(log, gormLogger.Config{
			SlowThreshold:        time.Second,
			Colorful:             true,
			ParameterizedQueries: true,
		}),
	})

	if err != nil {
		log.Errorf("Fail to connect to database - %+v", err)
		return nil, err
	}

	return &Database{db, log, cfg}, nil
}

func dsn(cfg *config.Config) string {
	dsn := []string{
		fmt.Sprintf("user=%s", cfg.Database.User),
		fmt.Sprintf("dbname=%s", cfg.Database.Name),
		fmt.Sprintf("host=%s", cfg.Database.Host),
		fmt.Sprintf("port=%d", cfg.Database.Port),
		fmt.Sprintf("TimeZone=%s", cfg.Database.TimeZone),
	}
	if lo.IsNotEmpty(cfg.Database.Password) {
		dsn = append(dsn, fmt.Sprintf("password=%s", cfg.Database.Password))
	}
	if lo.IsNotEmpty(cfg.Database.SSLMode) {
		dsn = append(dsn, fmt.Sprintf("sslmode=%s", cfg.Database.SSLMode))
	}

	return strings.Join(dsn, " ")
}
