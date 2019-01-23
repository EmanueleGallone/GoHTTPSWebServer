package main

import (
	"fmt"
	"./cert"
	"log"
	"net/http"
)

var listenAddr = "127.0.0.1"
var httpsPort = ":9090"

func main() {

	httpsAddr := listenAddr+httpsPort

	// Check if the cert files are available.
	// If they are not available, generate new ones.
	// Obviously, only for testing purpose.
	// for production, substitute the certificate with genuine ones
	if err := cert.Verify("cert.pem", "key.pem"); err != nil {
		addrPointer := &httpsAddr
		err = cert.Create(addrPointer)
		if err != nil {
			log.Fatal("Error: Couldn't create https certs.")
		}
	}

	http.HandleFunc("/", handler)
	//start HTTPS on Goroutine
	go http.ListenAndServeTLS(httpsAddr, "cert.pem", "key.pem", nil)

	// Start the HTTP server and redirect all incoming connections to HTTPS
	err3 := http.ListenAndServe(listenAddr+":80", http.HandlerFunc(redirectToHttps))
	if err3 != nil {
		fmt.Println("Error:",err3)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if _,err :=fmt.Fprintf(w, "Hi there! from HTTPS!"); err != nil {
		fmt.Println("Error:",err)
	}
}

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	// Redirect the incoming HTTP request. Note that "127.0.0.1:80" will only work if you are accessing the server from your local machine.
	fmt.Println("redirecting " + r.RemoteAddr + " to https")
	http.Redirect(w, r, "https://"+listenAddr+":9090"+r.RequestURI, http.StatusMovedPermanently)

}