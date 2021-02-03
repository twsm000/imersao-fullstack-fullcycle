package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/twsm000/imersao-fullstack-fullcycle/domain/model"
	_ "github.com/lib/pq"
	_ "gorm.io/driver/sqlite"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.env")
	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

// ConnectDB ...
func ConnectDB(env string) *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	if env == "test" {
		dsn = os.Getenv("dsnTest")
		db, err = gorm.Open(os.Getenv("dbTypeTest"), dsn)
	} else {
		dsn = os.Getenv("dsn")
		db, err = gorm.Open(os.Getenv("dbType"), dsn)
	}

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv("debug") == "true" {
		db.LogMode(true)
	}

	if os.Getenv("AutoMigrateDb") == "true" {
		db.AutoMigrate(
			&model.Account{},
			&model.Bank{},
			&model.Base{},
			&model.PixKey{},
			&model.Transaction{},
		)
	}

	return db
}
