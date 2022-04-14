package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"runtime"

	"github.com/EdisonLai/utils"
	"github.com/EdisonLai/utils/example/conf"
	"github.com/EdisonLai/utils/example/etcd"
	"github.com/EdisonLai/utils/example/mysql"
	"github.com/EdisonLai/utils/mysqlutils"
	"github.com/EdisonLai/utils/tools"
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
)

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

	logrus.SetLevel(logrus.DebugLevel)
	var err error
	if confPath != nil && *confPath != "" {
		if err = conf.ReadConf(*confPath); err != nil {
			logrus.Errorf("read config error! exit: %s", err.Error())
			os.Exit(0)
		}
	}

	if len(conf.Conf.ETCD.Address) != 0 {
		ctx := context.Background()
		if err := etcd.InitEtcdClient(ctx); err != nil {
			logrus.Error(err)
			os.Exit(0)
		}

		etcd.GetEtcdClient().Get(ctx, "", clientv3.WithPrefix())
	}

	if conf.Conf.Mysql.Address != "" {
		//mysql
		mysqlutils.InitMySQLConnection(conf.Conf.Mysql.Address, conf.Conf.Mysql.LogMode)

		mysql.MysqlTest()
	}

	var high, low = tools.ConvertIpToUint64(net.ParseIP("1.1.1.1"))
	fmt.Printf("%d, %d\n", high, low)
	var ip = tools.ConvertUint64ToIp(0, 4294901761)
	fmt.Printf("%s", ip.String())
}
