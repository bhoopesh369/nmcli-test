package mynmcli

import (
	"context"
	"fmt"

	"github.com/leberKleber/go-nmcli"
	"github.com/leberKleber/go-nmcli/device"
)

func Nmcli() {
	var nmcli = go_nmcli.NewNMCli()
	// fmt.Println(nmcli.Device.Status(context.TODO()))

	// nmcli -f DHCP4 device show oob_net0
	var _ device.WiFiListOptions = device.WiFiListOptions{
		IfName: "wlp45s0",
	}

	// fmt.Println(nmcli.Device.WiFiList(context.TODO(), options))

	fmt.Println(nmcli.Device.WiFiList(context.TODO(), device.WiFiListOptions{}))
}
