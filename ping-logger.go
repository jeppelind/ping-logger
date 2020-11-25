package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/go-ping/ping"
	"gopkg.in/ini.v1"
)

var version string

func main() {
	fmt.Println("Starting Ping-Logger", version)
	duration, interval, logpath, hostIP := getConfigValues()

	fmt.Println("Logging to", logpath)

	logfile, err := os.OpenFile(logpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	log.SetOutput(logfile)
	log.Printf("Ping-Logger %s. Running app against %s", version, hostIP)

	pingWithInterval(hostIP, time.Second*time.Duration(interval))

	time.Sleep(time.Second * time.Duration(duration))

	log.Println("Exit app")
	fmt.Println("Closing")
}

func getConfigValues() (int, int, string, string) {
	const maxDuration = 86400
	const minInterval = 5
	const defaultDuration = 21600
	const defaultInterval = 600
	const defaulLogPath = "./pinglogger.txt"

	config, err := ini.Load("config.ini")
	if err != nil {
		panic(err)
	}
	var duration = config.Section("timing").Key("duration").RangeInt(defaultDuration, 0, maxDuration)
	var interval = config.Section("timing").Key("interval").RangeInt(defaultInterval, minInterval, math.MaxInt32)
	var logpath = config.Section("paths").Key("log").MustString(defaulLogPath)
	var hostIP = config.Section("host").Key("ip").String()
	return duration, interval, logpath, hostIP
}

func pingWithInterval(host string, delay time.Duration) {
	ticker := time.NewTicker(delay)
	go func() {
		for range ticker.C {
			packetLoss := pingHost(host)
			if packetLoss > 0 {
				log.Printf("Could not reach host %s", host)
			}
		}
	}()
}

func pingHost(host string) (float64) {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		panic(err)
	}

	pinger.Timeout = time.Second * 3
	pinger.Count = 1
	pinger.SetPrivileged(true)

	err = pinger.Run()
	if err != nil {
		panic(err)
	}

	stats := pinger.Statistics()
	return stats.PacketLoss
}
