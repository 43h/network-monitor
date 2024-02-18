package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

const confFile = "conf.txt"

func loadConf() {
	f, err := os.Open(confFile)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parseLine(scanner.Text())
	}
}

func parseLine(line string) bool {

	if strings.Contains(line, "#") {
		return true
	}
	var elem = IP{}
	elem.ipv4 = net.ParseIP(line).To4()
	elem.avgRtt = 0
	ips = append(ips, elem)
	return true
}

func writeFile() {
	if len(ips) == 0 {
		return
	}
	err := os.Remove(confFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	f, err := os.Create(confFile)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err = os.Open(confFile)
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range ips {
		f.WriteString(value.ipv4.String() + "," +
			strconv.FormatInt(value.avgRtt, 10))
	}
}
