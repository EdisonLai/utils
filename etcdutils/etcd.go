package etcdutils

import (
	"context"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
)

type Config struct {
	Address     []string `toml:"address"`
	DialTimeout int      `toml:"DialTimeout"`
	OpTimeout   int      `toml:"OperationTimeout"`
	TLSFilePath string   `toml:"TLSFilePath"`
}

func InitEtcdClient(ctx context.Context, conf Config) (etcdClient *clientv3.Client, err error) {
	tlsConfig, err := getTlsConfig(ctx, conf.TLSFilePath)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if etcdClient, err = clientv3.New(clientv3.Config{
		Context:     context.Background(),
		Endpoints:   conf.Address,
		DialTimeout: time.Duration(conf.DialTimeout) * time.Second,
		TLS:         tlsConfig,
	}); err != nil {
		logrus.Errorf("can not init etcdutils: %v, timeout: %v", conf.Address, conf.DialTimeout)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.OpTimeout)*time.Second)
	defer cancel()
	if _, err = etcdClient.Put(ctx, "init", strconv.Itoa(int(time.Now().UnixNano()))); err != nil {
		logrus.Errorf("can not put etcdutils when init")
		etcdClient.Close()
		return nil, err
	}
	if _, err = etcdClient.Delete(ctx, "init"); err != nil {
		logrus.Errorf("can not Delete etcdutils when init")
		etcdClient.Close()
		return nil, err
	}

	return etcdClient, nil
}
