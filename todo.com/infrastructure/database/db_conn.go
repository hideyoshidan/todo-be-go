package database

import "gorm.io/gorm"

type DBConnIF interface {
	connectDB() (db *gorm.DB, err error)
}

type dbConf struct {
	user     string
	password string
	addr     string
	database string
	location string
	opt      *gorm.Config
}

func NewDBConn(conf DBConnIF) (db *gorm.DB, err error) {
	return conf.connectDB()
}
