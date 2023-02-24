package main

import (
	"flag"
	"io/ioutil"
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
	singlePort := flag.Int("p", 0, "Specify a Single Port")
	startPort := flag.Int("start", 1, "Specify Starting Port")
	endPort := flag.Int("end", 65535, "Specify Ending Port")
	outFile := flag.String("o", "", "Output to a file")
	flag.Parse()

	doneChannel := make(chan bool)
	activeThreadCount := 0
	log.Println("Scanning Host: " + *host2Scan)

	switch {
	case *tcpScan != false && *udpScan != true && *host2Scan != "":
		netWork := "tcp"
		if *singlePort != 0 {
			startPort = singlePort
			endPort = singlePort
		}
		for portNum := *startPort; portNum <= *endPort; portNum++ {
			activeThreadCount++
			go scanningTCP(netWork, *host2Scan, portNum, doneChannel, *outFile)
		}

		for {
			<-doneChannel
			activeThreadCount--
			if activeThreadCount == 0 {
				break
			}
		}

	case *udpScan != false && *tcpScan != true && *host2Scan != "":
		netWork := "udp"
		if *singlePort != 0 {
			startPort = singlePort
			endPort = singlePort
		}
		for portNum := *startPort; portNum <= *endPort; portNum++ {
			activeThreadCount++
			go scanningUDP(netWork, *host2Scan, portNum, doneChannel, *outFile)
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

func scanningTCP(netWork string, tcpScan string, port int, doneChannel chan bool, outFile string) {
	timeoutLength := 5 * time.Second
	conn, err := net.DialTimeout(netWork, tcpScan+":"+strconv.Itoa(port), timeoutLength)
	if err != nil {
		doneChannel <- false
		return // could not connect
	}
	conn.Close()
	checkConn := "[+] " + strconv.Itoa(port) + " connected"
	// TRY TO DO NEW BUFFER TO WRITE TO FILE
	if outFile != "" {
		outfile, err := os.Create(outFile)
		if err != nil {
			lPf("ERROR: Failed to create %s", outfile)
		}
		defer outfile.Close()
		err = ioutil.WriteFile(outFile, []byte(checkConn), 0644)
		if err != nil {
			lPf("ERROR: Failed to write to %s", outfile)
		}
	}
	lPf("[+] %d connnected", port)
	lPf(string(checkConn[:]))
	doneChannel <- true
}

func scanningUDP(netWork string, host2Scan string, port int, doneChannel chan bool, outFile string) {
	timeoutLength := 5 * time.Second
	conn, err := net.DialTimeout(netWork, host2Scan+":"+strconv.Itoa(port), timeoutLength)
	if err != nil {
		doneChannel <- false
		return // could not connect
	}
	conn.Close()
	checkConn := "[+] " + strconv.Itoa(port) + " connected"
	if outFile != "" {
		outfile, err := os.Create(outFile)
		if err != nil {
			lPf("ERROR: Failed to create %s", outfile)
		}
		defer outfile.Close()
		err = ioutil.WriteFile(outFile, []byte(checkConn), 0644)
		if err != nil {
			lPf("ERROR: Failed to write to %s", outfile)
		}
	}
	lPf(string(checkConn[:]))
	doneChannel <- true
}
