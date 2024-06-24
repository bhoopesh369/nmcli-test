package dhcp

import (
    "fmt"
    "log"
    "net"
    "time"

    _ "github.com/insomniacslk/dhcp/dhcpv4"
    "github.com/insomniacslk/dhcp/dhcpv4/client4"
)

type DHCPInfoOld struct {
    BroadcastAddress     net.IP
    LeaseTime            time.Duration
    ServerIdentifier     net.IP
    DomainNameServers    []net.IP
    Expiry               time.Time
    HostName             string
    IPAddress            net.IP
    NextServer           net.IP
    Routers              []net.IP
    SubnetMask           net.IPMask
}

func getDHCPInfoOld(interfaceName string) (*DHCPInfoOld, error) {
    client := client4.NewClient()
    iface, err := net.InterfaceByName(interfaceName)
    if err != nil {
        return nil, fmt.Errorf("interface error: %v", err)
    }

	stringIface := iface.Name

    conversation, err := client.Exchange(stringIface)
    if err != nil {
        return nil, fmt.Errorf("DHCP conversation error: %v", err)
    }

    if len(conversation) == 0 {
        return nil, fmt.Errorf("no DHCP messages received")
    }
    ack := conversation[len(conversation)-1]

    info := &DHCPInfoOld{
        BroadcastAddress:  ack.BroadcastAddress(),
        LeaseTime:         ack.IPAddressLeaseTime(time.Hour),
        ServerIdentifier:  ack.ServerIdentifier(),
        DomainNameServers: ack.DNS(),
        Expiry:            time.Now().Add(ack.IPAddressLeaseTime(time.Hour)),
        HostName:          string(ack.HostName()),
        IPAddress:         ack.YourIPAddr,
        NextServer:        ack.ServerIPAddr,
        Routers:           ack.Router(),
        SubnetMask:        ack.SubnetMask(),
    }

    return info, nil
}

func InfoOld() {
    info, err := getDHCPInfoOld("wlp45s0")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("IP Address: %s\n", info.IPAddress)
    fmt.Printf("Subnet Mask: %s\n", net.IP(info.SubnetMask))
    fmt.Printf("Lease Time: %s\n", info.LeaseTime)
    fmt.Printf("DHCP Server: %s\n", info.ServerIdentifier)
    fmt.Printf("DNS Servers: %v\n", info.DomainNameServers)

	fmt.Printf("Broadcast Address: %s\n", info.BroadcastAddress)
	fmt.Printf("Expiry: %s\n", info.Expiry)
	fmt.Printf("Host Name: %s\n", info.HostName)
	fmt.Printf("Next Server: %s\n", info.NextServer)
	fmt.Printf("Routers: %v\n", info.Routers)
	fmt.Printf("Subnet Mask: %s\n", info.SubnetMask)
}
