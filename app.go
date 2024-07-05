package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"speed_production/api"
	"speed_production/storage"
)

func getRunningIPAddress() string {
	addrs, _ := net.InterfaceAddrs()

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				// os.Stdout.WriteString(ipnet.IP.String() + "\n")
				// fmt.Printf("\n%v", ipnet.IP.String())
				return ipnet.IP.String()
			}
		}
	}
	return "0.0.0.0"
}

func main() {
	fmt.Println("")

	currPath, _ := os.Getwd()

	address := flag.String("address", "0.0.0.0", "listening address")
	port := flag.String("port", "8101", "listening port")
	storage.Fpath = *flag.String("path", currPath, "local files path")

	r := api.NewRouter()

	fmt.Printf("\thttp://%v:%v\n", address, port)
	// http.ListenAndServeTLS(address+":"+port, "localhost.crt", "localhost.key", r)

	http.ListenAndServe(*address+":"+*port, r)
}
