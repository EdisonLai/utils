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

func ConvertIpToUint64(ip net.IP) (ipHigh, ipLow uint64) {
	return binary.BigEndian.Uint64(ip[:8]), binary.BigEndian.Uint64(ip[8:16])
}

func ConvertUint64ToIp(ipHigh, ipLow uint64) (ip net.IP) {
	var high = make([]byte, 8)
	var low = make([]byte, 8)
	binary.BigEndian.PutUint64(high, ipHigh)
	binary.BigEndian.PutUint64(low, ipLow)
	ip = append(high, low...)
	return ip
}
