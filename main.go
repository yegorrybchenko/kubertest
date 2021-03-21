package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ifaces, _ := net.Interfaces()

	// fmt.Fprintf(w, "Hi there, I love kubernetes! My ip is %s", ip)
	// handle err
	var addresses []net.IP

	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err®
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			addresses = append(addresses, ip)
		}
	}
	result, err := json.Marshal(addresses)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(result)
}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
