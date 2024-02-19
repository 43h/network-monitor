package main

import (
	"fmt"
	"github.com/go-ping/ping"
	"net"
	"time"
)

type IP struct {
	ipv4   net.IP
	avgRtt int64
}

var ips = []IP{}

func (elem *IP) pingIP() {
	pinger, err := ping.NewPinger(elem.ipv4.String())
	if err != nil {
		return
	}
	pinger.SetPrivileged(true)
	pinger.Count = 3
	pinger.Interval = 500 * time.Millisecond
	pinger.Timeout = 3 * time.Second
	err = pinger.Run()
	if err == nil {
		stats := pinger.Statistics()
		fmt.Println(elem.ipv4.To4().String(), "up,", stats.AvgRtt.Milliseconds())
		if stats.AvgRtt != 0 {
			elem.avgRtt = stats.AvgRtt.Milliseconds()
		} else {
			elem.avgRtt = -1
		}
	} else {
		fmt.Println(elem.ipv4.To4().String(), "down, err")
		elem.avgRtt = -1
	}
}
