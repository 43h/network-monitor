package main

import (
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
	pinger.Count = 6
	pinger.Interval = 500 * time.Millisecond
	pinger.Timeout = 3 * time.Second
	err = pinger.Run()
	if err == nil {
		stats := pinger.Statistics()
		if stats.AvgRtt != 0 {
			elem.avgRtt = stats.AvgRtt.Milliseconds()
		} else {
			elem.avgRtt = -1
		}
	} else {
		elem.avgRtt = -1
	}
}
