package command

import (
	"github.com/jaynarol/BdoDownAlert/source/val"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func IsAlive() (val.Client, error) {
	output, err := exec.Command(val.PingCmd[0], val.PingCmd[1:]...).CombinedOutput()
	netstat := string(output)
	if err != nil {
		log.Printf("NETSTAT ERROR: %s\r\n", netstat)
		return val.Client{}, err
	}
	client := findLocalPort(netstat)
	return client, nil
}

func findLocalPort(netstat string) val.Client {
	for _, bdoPort := range val.BdoPorts {
		if connectionIndex := strings.Index(netstat, bdoPort); connectionIndex > -1 {
			if port := extactLocalPort(netstat, connectionIndex); port != "" {
				return val.Client{Found: true, Port: port}
			}
		}
	}
	return val.Client{}
}

func extactLocalPort(netstat string, connectionIndex int) string {
	lastSemicolonIndex := strings.LastIndex(netstat[:connectionIndex], ":")
	lineConnection := netstat[lastSemicolonIndex:connectionIndex]
	reg := regexp.MustCompile("^:([0-9]+).*$")
	port := reg.FindStringSubmatch(lineConnection)
	return port[1]
}
