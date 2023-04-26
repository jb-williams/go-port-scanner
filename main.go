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
	singlePort := flag.Int("p", 0, "Specify a Single Port")
	startPort := flag.Int("start", 1, "Specify Starting Port")
	endPort := flag.Int("end", 65535, "Specify Ending Port")
	// outFile := flag.String("o", "", "Output to a file")
	flag.Parse()

	doneChannel := make(chan bool)
	activeThreadCount := 0
	log.Println("Scanning Host: " + *host2Scan)
	if *singlePort != 0 {
		startPort = singlePort
		endPort = singlePort
	}

	switch {
	case *tcpScan != false && *udpScan != true && *host2Scan != "":
		netWork := "tcp"
		// for portNum := 1; portNum <= 65535; portNum++ {

		// outfile, err := os.Create(*outFile)
		// if err != nil {
		// 	lPf("ERROR: Failed to create %s", outfile)
		// }
		// return *outfile
		// defer outfile.Close()
		for portNum := *startPort; portNum <= *endPort; portNum++ {
			activeThreadCount++
			go scanningTCP(netWork, *host2Scan, portNum, doneChannel)
			// go scanningTCP(netWork, *host2Scan, portNum, doneChannel, *outfile)
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
		// for portNum := 1; portNum <= 65535; portNum++ {
		for portNum := *startPort; portNum <= *endPort; portNum++ {
			activeThreadCount++
			// go scanningUDP(netWork, *host2Scan, portNum, doneChannel, *outFile)
			go scanningUDP(netWork, *host2Scan, portNum, doneChannel)
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

func scanningTCP(netWork string, tcpScan string, port int, doneChannel chan bool) {
	timeoutLength := 5 * time.Second
	conn, err := net.DialTimeout(netWork, tcpScan+":"+strconv.Itoa(port), timeoutLength)
	if err != nil {
		doneChannel <- false
		return // could not connect
	}
	conn.Close()
	checkConn := "[+] " + strconv.Itoa(port) + " connected"
	// TRY TO DO NEW BUFFER TO WRITE TO FILE
<<<<<<< HEAD
	// 	err = ioutil.WriteFile(outFile, []byte(checkConn), 0644)
	// 	if err != nil {
	// 		lPf("ERROR: Failed to write to %s", outfile)
	// 	}
	// }
	// lPf("[+] %d connnected", port)
=======
	//if outFile != "" {
	//	outfile, err := os.Create(outFile)
	//	if err != nil {
	//		lPf("ERROR: Failed to create %s", outfile)
	//	}
	//	defer outfile.Close()
	//	err = ioutil.WriteFile(outFile, []byte(checkConn), 0644)
	//	if err != nil {
	//		lPf("ERROR: Failed to write to %s", outfile)
	//	}
	//}
	//lPf("[+] %d connnected", port)
>>>>>>> 52ec809 (messing with fileouput)
	lPf(string(checkConn[:]))
	doneChannel <- true
}

<<<<<<< HEAD
func scanningUDP(netWork string, host2Scan string, port int, doneChannel chan bool) {
=======
func scanningUDP(netWork string, host2Scan string, port int, doneChannel chan bool, outFile string) {
>>>>>>> 52ec809 (messing with fileouput)
	timeoutLength := 5 * time.Second
	conn, err := net.DialTimeout(netWork, host2Scan+":"+strconv.Itoa(port), timeoutLength)
	if err != nil {
		doneChannel <- false
		return // could not connect
	}
	conn.Close()
	checkConn := "[+] " + strconv.Itoa(port) + " connected"
<<<<<<< HEAD
	// if outFile != "" {
	// 	outfile, err := os.Create(outFile)
	// 	if err != nil {
	// 		lPf("ERROR: Failed to create %s", outfile)
	// 	}
	// 	defer outfile.Close()
	// 	err = ioutil.WriteFile(outFile, []byte(checkConn), 0644)
	// 	if err != nil {
	// 		lPf("ERROR: Failed to write to %s", outfile)
	// 	}
	// }
=======
	//if outFile != "" {
	//	outfile, err := os.Create(outFile)
	//	if err != nil {
	//		lPf("ERROR: Failed to create %s", outfile)
	//	}
	//	defer outfile.Close()
	//	err = ioutil.WriteFile(outFile, []byte(checkConn), 0644)
	//	if err != nil {
	//		lPf("ERROR: Failed to write to %s", outfile)
	//	}
	//}
>>>>>>> 52ec809 (messing with fileouput)
	lPf(string(checkConn[:]))
	doneChannel <- true
}
