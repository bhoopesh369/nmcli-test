package dhcp

import (
	"fmt"
	"log"
	"net"
	"time"

	_ "github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/insomniacslk/dhcp/dhcpv4/client4"
)
type DHCPInfo struct {
    IPAddress         net.IP
    SubnetMask        net.IPMask
    LeaseTime         time.Duration
    ServerIdentifier  net.IP
    DomainNameServers []net.IP
    SZTPRedirectURLs  string
}
type SZTPOptionCode struct{}

func (c SZTPOptionCode) Code() uint8 {
    return 143
}

func (c SZTPOptionCode) String() string {
    return "SZTPRedirectURLs"
}

var OptionSZTPRedirectURLs = SZTPOptionCode{}

func getDHCPInfo(interfaceName string) (*DHCPInfo, error) {
    client := client4.NewClient()
    iface, err := net.InterfaceByName(interfaceName)
    if err != nil {
        return nil, fmt.Errorf("interface error: %v", err)
    }

    conversation, err := client.Exchange(iface.Name)
    if err != nil {
        return nil, fmt.Errorf("DHCP conversation error: %v", err)
    }

    if len(conversation) == 0 {
        return nil, fmt.Errorf("no DHCP messages received")
    }

    ack := conversation[len(conversation)-1]
    
    info := &DHCPInfo{
        IPAddress:         ack.YourIPAddr,
        SubnetMask:        ack.SubnetMask(),
        LeaseTime:         ack.IPAddressLeaseTime(time.Hour),
        ServerIdentifier:  ack.ServerIdentifier(),
        DomainNameServers: ack.DNS(),
    }

	sztpOption := ack.Options.Get(OptionSZTPRedirectURLs)
    if sztpOption != nil {
        info.SZTPRedirectURLs = string(sztpOption)
    }

    return info, nil
}

func Info() {
    info, err := getDHCPInfo("wlp45s0")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("IP Address: %s\n", info.IPAddress)
    fmt.Printf("Subnet Mask: %s\n", net.IP(info.SubnetMask))
    fmt.Printf("Lease Time: %s\n", info.LeaseTime)
    fmt.Printf("DHCP Server: %s\n", info.ServerIdentifier)
    fmt.Printf("DNS Servers: %v\n", info.DomainNameServers)
    
    if info.SZTPRedirectURLs != "" {
        fmt.Printf("SZTP Redirect URLs: %s\n", info.SZTPRedirectURLs)
    } else {
        fmt.Println("No SZTP Redirect URLs found")
    }
}
