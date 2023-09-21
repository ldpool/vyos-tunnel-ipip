package util

import (
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

func SetupTunnel(remote, ip, tun, sourceip, ethname string) {
	remoteIP := PingDomain(remote)
	locIP := PingDomain(sourceip)
	script := []string{
		"#!/bin/vbash",
		"source /opt/vyatta/etc/functions/script-template",
		"configure",
		"set interfaces tunnel " + tun + " encapsulation ipip",
		"set interfaces tunnel " + tun + " address " + ip,
		"set interfaces tunnel " + tun + " mtu 1460",
		"set interfaces tunnel " + tun + " source-address " + locIP,
		"set interfaces tunnel " + tun + " remote " + remoteIP,
		"set interfaces tunnel " + tun + " source-interface " + ethname,
		"commit",
		"save",
		"exit",
	}

	scriptPath := "v.sh"
	CreateScript(scriptPath, script)
	defer os.Remove(scriptPath)
	cmd := exec.Command("./" + scriptPath)
	cmd.Run()
}

func PingDomain(domain string) string {
	var ip string
	ips, err := net.LookupIP(domain)
	if err != nil {
		log.Print(err)
		return ""
	}

	for _, v := range ips {
		ip = v.String()
	}
	return ip
}

func GetLocalIp() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}

func CreateScript(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	os.Chmod(filename, 0755)
	return nil
}
