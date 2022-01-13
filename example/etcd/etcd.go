package etcd

import (
	"context"
	"github.com/EdisonLai/utils/etcdutils"
	"github.com/EdisonLai/utils/example/conf"
	"github.com/coreos/etcd/clientv3"
	"github.com/sirupsen/logrus"
)

var etcdClient *clientv3.Client

func InitEtcdClient(ctx context.Context) (err error){
	etcdClient, err = etcdutils.InitEtcdClient(ctx, etcdutils.Config{
		Address:     conf.Conf.ETCD.Address,
		DialTimeout: conf.Conf.ETCD.DialTimeout,
		OpTimeout:   conf.Conf.ETCD.OpTimeout,
		TLSFilePath: conf.Conf.ETCD.TLSFilePath,
	})
	if err != nil {
		logrus.Error(err)
		return
	}
	return
}
