package configutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"reflect"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

func ReadConfig(f string, checkConfValidation func(interface{}) error, GConfig interface{}) (err error) {
	if reflect.ValueOf(GConfig).Type().Kind() != reflect.Ptr {
		err = fmt.Errorf("GConfig should be Ptr")
		logrus.Error(err)
		return err
	}
	if f == "" {
		panic("invalid config file name")
	}
	file, err := os.OpenFile(f, os.O_RDWR, 0666)
	if err != nil {
		// for test
		file, err = os.OpenFile(path.Join("..", f), os.O_RDWR, 0666)
		if err != nil {
			panic(err.Error())
		}
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err.Error())
	}

	_, err = toml.Decode(string(content), GConfig)
	if err != nil {
		panic(err.Error())
	}

	if checkConfValidation != nil {
		if err = checkConfValidation(GConfig); err != nil {
			logrus.Errorf("read config error: %s", err.Error())
			return err
		}
	}

	return nil
}
