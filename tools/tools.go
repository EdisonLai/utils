package tools

import (
	"encoding/binary"
	"net"
)

func ConvertIpToUint32(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func ConvertUint32ToIp(num uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, num)
	return ip
}
