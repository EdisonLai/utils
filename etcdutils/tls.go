package etcdutils

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/pkg/transport"
)

//etcdctl --endpoints=:2379 --cacert="ca.pem" --cert="etcdutils-client.pem" --key="etcdutils-client-key.pem" get "" --prefix
func getTlsConfig(ctx context.Context, configPath string) (config *tls.Config, err error) {
	if configPath == "" {
		logrus.Debugf("start etcdutils without tls certification")
		return nil, nil
	}

	tlsInfo := transport.TLSInfo{
		CertFile:      configPath + "etcdutils-client.pem",
		KeyFile:       configPath + "etcdutils-client-key.pem",
		TrustedCAFile: configPath + "ca.pem",
	}
	tlsConfig, err := tlsInfo.ClientConfig()
	if err != nil {
		err = fmt.Errorf("tls info err")
		logrus.Error(err)
		return nil, err
	}
	return tlsConfig, nil
}
