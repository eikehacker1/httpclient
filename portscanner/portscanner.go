// portscanner.go
package portscanner

import (
    "strconv"
    "strings"
)

func ParsePortList(portsStr string) []int {
    portList := []int{}
    portRanges := strings.Split(portsStr, ",")
    for _, pr := range portRanges {
        if strings.Contains(pr, "-") {
            // Handle port ranges, e.g., "80-85"
            rangeParts := strings.Split(pr, "-")
            if len(rangeParts) != 2 {
                continue
            }
            startPort, err := strconv.Atoi(rangeParts[0])
            if err != nil {
                continue
            }
            endPort, err := strconv.Atoi(rangeParts[1])
            if err != nil {
                continue
            }
            for p := startPort; p <= endPort; p++ {
                portList = append(portList, p)
            }
        } else {
            // Handle single ports, e.g., "80"
            port, err := strconv.Atoi(pr)
            if err != nil {
                continue
            }
            portList = append(portList, port)
        }
    }
    return portList
}
