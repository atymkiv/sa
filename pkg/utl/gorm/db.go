package gorm

import (
	"fmt"
	sa "github.com/atymkiv/sa/model"
	"github.com/atymkiv/sa/pkg/utl/config"
	"github.com/cenkalti/backoff"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type DB struct {
	 gorm.DB
}

var dbInstance *DB

func GetDbInstance(cfg *config.Database) (*DB, error) {
	if dbInstance == nil {
		bdInstance, err := New(cfg)
		return bdInstance, err
	}
	return dbInstance, nil
}

func New(cfg *config.Database) (*DB, error) {
	connectString := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Name)
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 1 * time.Second

	var db gorm.DB
	operation := func() error {
		dbT, err := gorm.Open("mysql", connectString)
		if err != nil {
			return err
		}
		db = *dbT
		return nil
	}

	err := backoff.Retry(operation, b)
	if err != nil {
		log.Fatalf("error after retrying: %v", err)
	}
	//defer db.Close()
	db.AutoMigrate(sa.Shop{}, sa.Address{})
	dbIn := DB{db}
	return &dbIn, nil
}
