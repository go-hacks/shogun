package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

func main() {
	var addr [257]net.UDPAddr
	var serv [257]*net.UDPConn
	for i := 0; i < 257; i++ {
		addr[i] = net.UDPAddr{
			Port: 9000+i,
			IP: net.ParseIP("127.0.0.1"),
		}
	}
	for i := 0; i < 257; i++ {
		serv[i], _ = net.ListenUDP("udp", &addr[i])
	}
	p := make([]byte, 1)
	var remoteaddr [257]*net.UDPAddr
	var err error
	fmt.Printf("Getting remote addresses...")
	for i := 0; i < 257; i++ {
		_, remoteaddr[i], err = serv[i].ReadFromUDP(p)
		if err !=  nil {
			fmt.Printf("Start read error  %v", err)
		}
	}
	fmt.Printf("Done.\n")
	fmt.Printf("Loading file into memory...")
	fileInfo, _ := os.Stat(os.Args[1])
	fileSize := int(fileInfo.Size())
	inFile, _ := os.OpenFile(os.Args[1],os.O_RDONLY,0666)
	readFile := bufio.NewReader(inFile)
	data := make([]byte, fileSize)
	_, _ = readFile.Read(data)
	fmt.Printf("Done.\n")
	fmt.Printf("Sending file to %v...", remoteaddr[256])
	for i := 0; i < len(data); i++ {
		v := int(data[i])
		writer(serv[v], remoteaddr[v])
		// Wait for signal to send next byte
		_, _, _ = serv[256].ReadFromUDP(p)
	}
	writer(serv[256], remoteaddr[256])
	fmt.Printf("Complete!\n")
}

func writer(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte(""), addr)
	if err != nil {
		fmt.Printf("Couldn't send byte %v", err)
	}
}
