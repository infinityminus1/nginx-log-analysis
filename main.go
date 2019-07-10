package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Create map to maintain list and count of IPs
var ips = make(map[string]int)

func main() {
	fmt.Println("Beginning Analysis")

	// Open nginx access log
	accessLog, err := os.Open("/var/log/nginx/access.log")
	if err != nil {
		log.Fatal(err)
	}
	defer accessLog.Close()

	// loop through line entries, add IPs to the map, increment the counter.
	accessLogReader := bufio.NewReader(accessLog)
	
	for {
		line, err := accessLogReader.ReadString('\n')
		if err == io.EOF {
			break
		}

		ip := strings.Split(line, "-")[0]

		if _, ok := ips[ip]; ok {
			ips[ip]++
		} else {
			ips[ip] = 1
		}
	}

	// print the IPs map
	fmt.Println(ips)
	for key, value := range ips {
		fmt.Printf("%s\t=>\t%d requests.\n", key, value)
	}
}
