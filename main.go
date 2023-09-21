package main

import (
	"flag"
	"log"
	"time"
	"vyos/util"
)

var (
	remoteIP string
	sourceIP string
)

func main() {
	locip := util.GetLocalIp()
	remote := flag.String("remote", "", "remote")
	tunip := flag.String("tunip", "198.0.0.1/30", "tunIp")
	tun := flag.String("tun", "tun0", "tun")
	sourceip := flag.String("sourceip", locip, "sourceIp")
	ethname := flag.String("ethname", "eth0", "ethname")

	flag.Parse()
	if *remote == "" {
		log.Print("remote must")
		return
	}
	remoteIP = util.PingDomain(*remote)
	sourceIP = util.PingDomain(*sourceip)
	timer := time.NewTicker(60 * time.Second)
	defer timer.Stop()

	util.SetupTunnel(*remote, *tunip, *tun, *sourceip, *ethname)

	go func() {
		for range timer.C {
			newRemoteIP := util.PingDomain(*remote)
			newSourceIP := util.PingDomain(*sourceip)
			if newRemoteIP != remoteIP || newSourceIP != sourceIP {
				log.Printf("Remote IP changed from %s to %s", remoteIP, newRemoteIP)
				log.Printf("Source IP changed from %s to %s", sourceIP, newSourceIP)
				util.SetupTunnel(*remote, *tunip, *tun, *sourceip, *ethname)
				remoteIP = newRemoteIP
				sourceIP = newSourceIP
			}
		}
	}()

	select {}
}
