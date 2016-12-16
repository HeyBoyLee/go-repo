package main

import (
	"fmt"
	"os"
	"strings"
	"net"
	"io/ioutil"
)

func main() {
	fmt.Println(os.Getwd())
	fmt.Println(strings.TrimSpace("  strings  "));

	ifaces, _ := net.Interfaces()
	fmt.Println(ifaces)
	for _, iface := range ifaces {
		fmt.Println(iface.Addrs())
	}

	bs, _ := ioutil.ReadFile("/proc/stat")
	fmt.Println(bs);
}
