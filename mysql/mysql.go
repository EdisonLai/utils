package mysql

import (
	"fmt"
	"time"

	gormlogruslogger "github.com/aklinkert/go-gorm-logrus-logger"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func InitMySQLConnection(addr string) (err error) {
	mysqlConfig := mysql.Config{
		DSN:                    addr,
		DontSupportRenameIndex: true,
	}

	db, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: gormlogruslogger.NewGormLogrusLogger(
			logrus.WithFields(logrus.Fields{
				"service": "mysql",
			}),
			2*time.Second),
	})
	if err != nil {
		err = fmt.Errorf("%v: %s", err, "Building mysql connection failed!")
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		err = fmt.Errorf("Get mysql connect pool %v failed!\n%v", addr, err)
		logrus.Error(err)
		return err
	}
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(50)
	return nil
}
