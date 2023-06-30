package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

type prefix struct {
	IPPrefix string `json:"ip_prefix"`
	Region   string `json:"region"`
}

type IPRange struct {
	Prefixes []prefix `json:"prefixes"`
}

func main() {
	ipPtr := flag.String("ip", "", "IP address to lookup")
	flag.Parse()

	if *ipPtr == "" {
		log.Fatalln("Please provide an IP address using the --ip flag.")
	}

	resp, err := http.Get("https://ip-ranges.amazonaws.com/ip-ranges.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data IPRange
	json.Unmarshal(body, &data)

	ip := net.ParseIP(*ipPtr) // use the provided IP

	for _, prefix := range data.Prefixes {
		_, pnet, _ := net.ParseCIDR(prefix.IPPrefix)
		if pnet.Contains(ip) {
			log.Printf("The IP address %s is in the %s region.\n", ip, prefix.Region)
			return
		}
	}

	log.Printf("The IP address %s does not belong to any AWS region.\n", ip)
}


