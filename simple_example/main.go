package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/EdisonLai/utils"
	"github.com/EdisonLai/utils/configutils"
	"github.com/EdisonLai/utils/etcdutils"
	"github.com/coreos/etcd/clientv3"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
)

var Conf Config

type EtcdConfig struct {
	Address     []string `toml:"address"`
	DialTimeout int      `toml:"DialTimeout"`
	OpTimeout   int      `toml:"OperationTimeout"`
	TLSFilePath string   `toml:"TLSFilePath"`
}

type Config struct {
	ETCD EtcdConfig `toml:"etcd"`
}

func checkConfValidation(c interface{}) (err error) {
	conf := c.(Config)
	logrus.Debug("%+v", conf)
	return nil
}

func main() {
	runtime.GOMAXPROCS(1)
	var confPath = flag.String("config", "", "config file path")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		switch {
		case args[0] == "version" || args[0] == "v":
			fmt.Println("version ", utils.Version)
			fmt.Println("git Commit Hash:", utils.GitVersion)
			fmt.Println("build Time:", utils.BuildTime)
			return
		default:
			fmt.Printf("xxx version (v) --- Show the Version, Git Commit Hash and Build Time\n")
		}
	}

	if err := configutils.ReadConfig(*confPath, checkConfValidation, &Conf); err != nil {
		logrus.Errorf("read config error! exit: %s", err.Error())
		os.Exit(0)
	}

	var err error
	ctx := context.Background()
	var etcdClient *clientv3.Client
	if etcdClient, err = etcdutils.InitEtcdClient(ctx, etcdutils.Config{
		Address:     Conf.ETCD.Address,
		DialTimeout: Conf.ETCD.DialTimeout,
		OpTimeout:   Conf.ETCD.OpTimeout,
		TLSFilePath: Conf.ETCD.TLSFilePath,
	}); err != nil {
		logrus.Error(err)
		os.Exit(0)
	}

	etcdClient.Get(ctx, "", clientv3.WithPrefix())
}
