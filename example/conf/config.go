package conf

import (
	"github.com/EdisonLai/utils/configutils"
	"github.com/sirupsen/logrus"
)

var Conf Config

type EtcdConfig struct {
	Address     []string `toml:"address"`
	DialTimeout int      `toml:"DialTimeout"`
	OpTimeout   int      `toml:"OperationTimeout"`
	TLSFilePath string   `toml:"TLSFilePath"`
}

type Config struct {
	ETCD        EtcdConfig        `toml:"etcd"`
}


func checkConfValidation(c interface{}) (err error){
	conf := c.(Config)
	logrus.Debug("%+v", conf)
	return nil
}

func ReadConf(path string) (err error) {
	return configutils.ReadConfig(path, checkConfValidation, &Conf)
}