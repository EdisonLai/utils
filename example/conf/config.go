package conf

import (
	"github.com/EdisonLai/utils/configutils"
	"github.com/EdisonLai/utils/etcdutils"
	"github.com/EdisonLai/utils/mysqlutils"
	"github.com/sirupsen/logrus"
)

var Conf Config

type Config struct {
	ETCD  etcdutils.Config  `toml:"etcd"`
	Mysql mysqlutils.Config `toml:"mysql"`
}

func checkConfValidation(c interface{}) (err error) {
	conf := c.(*Config)
	logrus.Debug("%+v", conf)
	return nil
}

func ReadConf(path string) (err error) {
	return configutils.ReadConfig(path, checkConfValidation, &Conf)
}
