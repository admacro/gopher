package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (addr IPAddr) String() string {
	var strs [len(addr)]string
	for i, intgr := range addr {
		strs[i] = strconv.Itoa(int(intgr))
	}
	return strings.Join(strs[:], ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
