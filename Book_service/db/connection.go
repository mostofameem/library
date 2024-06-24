package db

import (
	"book_service/config"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func connect(dbConfig config.DBConfig) *sqlx.DB {
	dbSource := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	if !dbConfig.EnableSSLMode {
		dbSource += " sslmode=disable"
	}

	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	dbCon.SetConnMaxIdleTime(
		time.Duration(dbConfig.MaxIdleTimeInMinute * int(time.Minute)),
	)

	return dbCon
}

func ConnectDB() {
	conf := config.GetConfig()

	Db = connect(conf.DB)
	slog.Info("Connected to database")

}

func CloseDB() {
	if err := Db.Close(); err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info("Disconnected from database")
}
