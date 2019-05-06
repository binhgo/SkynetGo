package Connection

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func getIPAndPort(addr net.Addr) (IP string, Port string) {
	s := strings.Split(addr.String(), ":")
	return s[0], s[1]
}
