package etcd

import (
	"context"

	"github.com/EdisonLai/utils/etcdutils"
	"github.com/EdisonLai/utils/example/conf"
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
)

var etcdClient *clientv3.Client

func GetEtcdClient() *clientv3.Client {
	return etcdClient
}

func InitEtcdClient(ctx context.Context) (err error) {
	etcdClient, err = etcdutils.InitEtcdClient(ctx, conf.Conf.ETCD)
	if err != nil {
		logrus.Error(err)
		return
	}
	return
}
