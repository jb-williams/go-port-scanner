package main

import (
	"flag"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var lPf = log.Printf

func main() {
	help := flag.Bool("h", false, "Show help")
	tcpScan := flag.Bool("t", false, "TCP Scan")
	udpScan := flag.Bool("u", false, "UDP Scan")
	host2Scan := flag.String("a", "", "REQUIRED: Host to Scan/Attack")

	doneChannel := make(chan bool)
	activeThreadCount := 0
	log.Println("Scanning Host: " + *host2Scan)

	switch {
	case *tcpScan && !*udpScan && *host2Scan != "":
		netWork := "tcp"
		for portNum := 1; portNum <= 65535; portNum++ {
			activeThreadCount++
			go scanningTCP(netWork, *host2Scan, portNum, doneChannel)
		}

		for {
			<-doneChannel
			activeThreadCount--
			if activeThreadCount == 0 {
				break
			}
		}

	case *udpScan && *host2Scan != "":
		netWork := "udp"
		for portNum := 1; portNum <= 65535; portNum++ {
			activeThreadCount++
			go scanningTCP(netWork, *host2Scan, portNum, doneChannel)
		}

		for {
			<-doneChannel
			activeThreadCount--
			if activeThreadCount == 0 {
				break
			}
		}

	case *help:
		flag.Usage()
		os.Exit(0)
	default:
		flag.Usage()
		os.Exit(0)
	}
}

func scanningTCP(netWork string, host2Scan string, port int, doneChannel chan bool) {
	timeoutLength := 5 * time.Second
	conn, err := net.DialTimeout(netWork, host2Scan+":"+strconv.Itoa(port), timeoutLength)
	//conn, err := net.DialTimeout("tcp", host + ":" + strconv.Itoa(port), timeoutLength)
	if err != nil {
		doneChannel <- false
		return // could not connect
	}
	conn.Close()
	lPf("[+] %d connnected", port)
	doneChannel <- true
}

func scanningUDP(netWork string, host2Scan string, port int, doneChannel chan bool) {
	timeoutLength := 5 * time.Second
	conn, err := net.DialTimeout(netWork, host2Scan+":"+strconv.Itoa(port), timeoutLength)
	//conn, err := net.DialTimeout("tcp", host + ":" + strconv.Itoa(port), timeoutLength)
	if err != nil {
		doneChannel <- false
		return // could not connect
	}
	conn.Close()
	lPf("[+] %d connnected", port)
	doneChannel <- true
}
