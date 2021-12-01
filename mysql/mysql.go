package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func InitMySQLConnection(addr string) (err error) {
	db, err = gorm.Open("mysql", addr)
	if err != nil {
		err = fmt.Errorf("%v: %s", err, "Building mysql connection failed!")
		return err
	}
	db.DB().SetMaxIdleConns(1)
	db.DB().SetMaxOpenConns(50)
	return nil
}
