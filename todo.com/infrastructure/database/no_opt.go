package database

import (
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	driver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	ct "todo.com/config"
)

type DBConfWithNoOpt struct {
	*dbConf
}

func NewWithNoOpt(user, password, addr, database, location string) *DBConfWithNoOpt {
	return &DBConfWithNoOpt{
		&dbConf{
			user:     user,
			password: password,
			addr:     addr,
			database: database,
			location: location,
			opt:      nil,
		},
	}
}

func (conf *DBConfWithNoOpt) connectDB() (db *gorm.DB, err error) {
	tzone, err := time.LoadLocation(conf.dbConf.location)
	if err != nil {
		log.Fatalln("DB could not be connected.", err)
		return nil, err
	}

	c := mysql.Config{
		DBName:    conf.dbConf.database,
		User:      conf.dbConf.user,
		Passwd:    conf.dbConf.password,
		Addr:      conf.dbConf.addr,
		Net:       ct.DB_NET,
		ParseTime: true,
		Collation: ct.DB_CLLATION,
		Loc:       tzone,
	}

	db, err = gorm.Open(driver.Open(c.FormatDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalln("DB could not be connected.", err)
		return nil, err
	}

	log.Printf("DB Connection established.")

	return db, nil
}
