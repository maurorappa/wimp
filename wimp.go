package main

import (
	"fmt"
	"github.com/ipinfo/go-ipinfo/ipinfo"
	"log"
	"net"
	"net/http"
	"os"
	"syscall"
)

import (
	//#include <unistd.h>
	//#include <errno.h>
	"C"
)

func init() {
    if syscall.Getuid() == 0 {
        log.Println("Running as root, downgrading to user nobody")
        cerr, errno := C.setgid(C.__gid_t(65534))
        if cerr != 0 {
            log.Fatalln("Unable to set GID due to error:", errno)
        }
        cerr, errno = C.setuid(C.__uid_t(65534))
        if cerr != 0 {
            log.Fatalln("Unable to set UID due to error:", errno)
        }
    }
}

var client *ipinfo.Client

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	// we look for CloudFlare header
	ip := r.Header.Get("CF-Connecting-IP")
	fmt.Fprintln(w, ip)
}

func detailHandler(w http.ResponseWriter, r *http.Request) {
	// we look for CloudFlare header
	ip := r.Header.Get("CF-Connecting-IP")
	fmt.Fprintf(w, "Hi there, you come from %s\n", ip)
	info, err := client.GetInfo(net.ParseIP(ip))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, info.Hostname)
}

func main() {
	// the API token is needed for ipinfo service
	auth := os.Getenv("TOKEN")
	if len(auth) == 0 {
		log.Fatal("Please specify the env var TOKEN")
	}
	authTransport := ipinfo.AuthTransport{Token: auth}
	httpClient := authTransport.Client()
	client = ipinfo.NewClient(httpClient)
	fmt.Printf("Wimp server ready\n")
	http.HandleFunc("/s", simpleHandler)
	http.HandleFunc("/d", detailHandler)
	log.Fatal(http.ListenAndServe(":8181", nil))
}
