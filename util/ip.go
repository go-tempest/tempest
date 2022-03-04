package util

import (
    "errors"
    "net"
)

// GetLocalIP Obtain the valid IP address of the first valid NIC from the list of native NICs
func GetLocalIP() (net.IP, error) {

    ifaces, err := net.Interfaces()
    if err != nil {
        return nil, err
    }

    for _, iface := range ifaces {

        // interface down
        if iface.Flags&net.FlagUp == 0 {
            continue
        }

        // loopback interface
        if iface.Flags&net.FlagLoopback != 0 {
            continue
        }

        addrs, err := iface.Addrs()
        if err != nil {
            return nil, err
        }

        for _, addr := range addrs {
            ip := getIPfromAddr(addr)
            if ip == nil {
                continue
            }
            return ip, nil
        }
    }

    return nil, errors.New("no valid ip address found")
}

func getIPfromAddr(addr net.Addr) net.IP {

    var ip net.IP

    switch v := addr.(type) {
    case *net.IPNet:
        ip = v.IP
    case *net.IPAddr:
        ip = v.IP
    }

    if ip == nil || ip.IsLoopback() {
        return nil
    }

    // not an ipv4 address
    ip = ip.To4()
    if ip == nil {
        return nil
    }

    return ip
}
