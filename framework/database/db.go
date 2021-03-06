package database

import (
	"github.com/jinzhu/gorm"
	"log"
	"video-encoder/domain"
)

type Database struct {
	Db 			  *gorm.DB
	Dsn 		  string
	DsnTest 	  string
	DbType 		  string
	DbTypeTest 	  string
	Debug 		  bool
	AutoMigrateDb bool
	Env 		  string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstance := &Database{
		Env:           "Test",
		DsnTest:       ":memory",
		DbTypeTest:    "sqlite3",
		Debug:         true,
		AutoMigrateDb: true,
	}

	connection, err := dbInstance.Connect()

	if err != nil {
		log.Fatalf("Test db error: %v", err)
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

	if d.AutoMigrateDb {
		d.Db.AutoMigrate(&domain.Video{}, &domain.Job{})
	}

	return d.Db, nil
}