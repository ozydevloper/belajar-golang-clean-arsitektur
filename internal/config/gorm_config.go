package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(cfg ConfigDB) *gorm.DB {

	pgConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSL)

	db, err := gorm.Open(postgres.Open(pgConn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      true,
				Colorful:                  false,
			}),
	})
	if err != nil {
		log.Panicf("failed to connect database: %+v", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Panicf("failed to connect database: %+v", err)
	}

	connection.SetMaxOpenConns(cfg.OpenConn)
	connection.SetMaxIdleConns(cfg.IdleConn)
	connection.SetConnMaxIdleTime(cfg.IdleLifetime)
	connection.SetConnMaxLifetime(cfg.Lifetime)

	return db
}

type LogrusWriter struct {
	Log *logrus.Logger
}

func (l *LogrusWriter) Printf(message string, args ...any) {
	l.Log.Tracef(message, args...)
}
