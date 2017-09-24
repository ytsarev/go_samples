package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const ipService string = "https://ifconfig.co/json"

type ipInfo struct {
	IP       string
	Country  string
	City     string
	Hostname string
}

func getIPinfo(url string) ipInfo {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error occurred: %s\n", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	info := ipInfo{}
	json.Unmarshal([]byte(body), &info)
	return info
}

func main() {
	info := getIPinfo(ipService)
	fmt.Printf("IP: %s\nCountry: %s\nCity: %s\nHostname: %s\n",
		info.IP, info.Country, info.City, info.Hostname)
}
