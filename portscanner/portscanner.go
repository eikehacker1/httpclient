// portscanner.go
package portscanner

import (
	"net"
	"strconv"
)

func IsPortOpen(host, port string) bool {
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func ScanPorts(host string, ports []int) []int {
	openPorts := []int{}
	for _, port := range ports {
		portStr := strconv.Itoa(port)
		if IsPortOpen(host, portStr) {
			openPorts = append(openPorts, port)
		}
	}
	return openPorts
}
