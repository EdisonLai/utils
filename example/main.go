package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/EdisonLai/utils"
	"github.com/EdisonLai/utils/example/conf"
	"github.com/EdisonLai/utils/example/etcd"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
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

	if err := conf.ReadConf(*confPath); err != nil {
		logrus.Errorf("read config error! exit: %s", err.Error())
		os.Exit(0)
	}

	ctx := context.Background()
	if err := etcd.InitEtcdClient(ctx); err != nil {
		logrus.Error(err)
		os.Exit(0)
	}
}
