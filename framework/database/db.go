package database

import (
	"github.com/jinzhu/gorm"
	"github.com/m3k3r1/video-encoder/domain"
	"log"
)

type Database struct {
	Db          *gorm.DB
	Dsn         string
	DsnTest     string
	DbType      string
	DbTypeTest  string
	Debug       bool
	AutoMigrate bool
	Env         string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Env = "test"
	dbInstance.DbType = "sqlite3"
	dbInstance.DsnTest = ":memory:"
	dbInstance.AutoMigrate = true
	dbInstance.Debug = true

	connection, err := dbInstance.Connect()
	if err != nil {
		log.Fatalf("Test DB error: %v", err)
	}

	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	if d.Env != "test" {
		d.Db, err = gorm.Open(d.DbType, d.Dsn)
	} else {
		d.Db, err = gorm.Open(d.DbTypeTest, d.DsnTest)
	}

	if err != nil {
		return nil, err
	}

	if d.Debug {
		d.Db.LogMode(true)
	}

	if d.AutoMigrate {
		d.Db.AutoMigrate(&domain.Video{}, &domain.Job{})
	}

	return d.Db, nil
}
