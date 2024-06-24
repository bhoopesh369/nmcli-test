package main

import (
	"fmt"

	"github.com/bhoopesh369/nmcli-test/dhcp"
	"github.com/bhoopesh369/nmcli-test/nmcli"
)

func main() {
	fmt.Println("Hello, World!")
	mynmcli.Nmcli()

	fmt.Println("\n DHCP Info:")
	go dhcp.Info()

	var input string
	fmt.Scanln(&input)
}
