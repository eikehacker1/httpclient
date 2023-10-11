// portscanner.go
package portscanner

import (
    "fmt"
    "net"
    "strconv"
    "strings"
)

// ScanPorts verifica a abertura de portas e verifica se há serviços HTTP.
func ScanPorts(host string, ports []int) {
    for _, port := range ports {
        target := fmt.Sprintf("%s:%d", host, port)
        conn, err := net.Dial("tcp", target)
        if err != nil {
            fmt.Printf("%s [Port %d]: Closed\n", host, port)
            continue
        }
        defer conn.Close()

        fmt.Printf("%s [Port %d]: Open\n", host, port)

        // Você pode adicionar aqui a lógica para verificar se há um serviço HTTP em execução na porta.
    }
}

// ParsePortList analisa uma lista de portas no formato "80,443,8080".
func ParsePortList(portsStr string) []int {
    portStrings := strings.Split(portsStr, ",")
    ports := []int{}
    for _, portStr := range portStrings {
        port, err := strconv.Atoi(portStr)
        if err != nil {
            fmt.Printf("Error parsing port %s: %v\n", portStr, err)
            continue
        }
        ports = append(ports, port)
    }
    return ports
}
