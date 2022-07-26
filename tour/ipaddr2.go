package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (addr *IPAddr) String() string {
	var strs [len(addr)]string
	for i, intgr := range addr {
		strs[i] = strconv.Itoa(int(intgr))
	}
	return strings.Join(strs[:], ".")
}

// https://stackoverflow.com/questions/48296826/stringer-method-requires-value
func main() {
	ip := IPAddr{127, 0, 0, 1}
	fmt.Printf("ip: %v\n", ip)                   // ip: [127 0 0 1] default String() of array type is invoked
	fmt.Printf("ip.String(): %v\n", ip.String()) // ip.String(): 127.0.0.1
}
