package portsc

import (
	"net"
	"strconv"
	"time"
)

type PortInfo struct {
	IP   string
	Port int
}

func ScanIp(ip string, port int) bool {

	address := ip + ":" + strconv.Itoa(port)

	conn, err := net.DialTimeout("tcp", address, 60*time.Second)
	if err != nil {
		defer conn.Close()
		return true
	}
	return false

}

func (pi PortInfo) String() string {
	return pi.IP + ":" + strconv.Itoa(pi.Port)
}
