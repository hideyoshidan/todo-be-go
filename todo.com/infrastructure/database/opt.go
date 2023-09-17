package database

import (
	"log"
	"time"

	"gorm.io/gorm"
	ct "todo.com/config"

	"github.com/go-sql-driver/mysql"
	driver "gorm.io/driver/mysql"
)

type DBConfWithOpt struct {
	*dbConf
}

func NewWithOpt(user, password, addr, database, location string, opt *gorm.Config) *DBConfWithOpt {
	return &DBConfWithOpt{
		&dbConf{
			user:     user,
			password: password,
			addr:     addr,
			database: database,
			location: location,
			opt:      opt,
		},
	}
}

func (conf *DBConfWithOpt) connectDB() (db *gorm.DB, err error) {
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

	db, err = gorm.Open(driver.Open(c.FormatDSN()), conf.dbConf.opt)
	if err != nil {
		log.Fatalln("DB could not be connected.", err)
		return nil, err
	}

	return db, nil
}
