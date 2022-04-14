package mysql

import (
	"fmt"

	"github.com/EdisonLai/utils/mysqlutils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Service struct {
	Name string `gorm:"column:name"`
	LDC  string `gorm:"column:logical_dc"`
}

func (s *Service) BeforeCreate(tx *gorm.DB) (err error) {
	s.LDC = "lxf_test"
	logrus.Debug("BeforeCreate")
	return
}

func (s *Service) AfterCreate(tx *gorm.DB) (err error) {
	logrus.Debug("AfterCreate")
	return
}

func MysqlTest() {
	var s = Service{
		Name: "tes",
	}
	tx := mysqlutils.GetDB().Begin()
	tx.Create(&s)
	s.Name = "t"
	tx.Create(&s)
	fmt.Print("finish")
	tx.Commit()
}
