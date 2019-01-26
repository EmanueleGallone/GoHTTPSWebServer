package main

import (
	"fmt"
	"./cert"
	"log"
	"net/http"
	"text/template"
)

const listenAddr = "127.0.0.1"
const httpsPort = ":9090"
var httpsAddr = listenAddr+httpsPort

func main() {

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
	go HTTPS()

	// Start the HTTP server and redirect all incoming connections to HTTPS
	err3 := http.ListenAndServe(listenAddr+":80", http.HandlerFunc(redirectToHttps))
	if err3 != nil {
		log.Println("Error:",err3)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	//if _,err :=fmt.Fprintf(w, "Hi there! from HTTPS!"); err != nil {
	//	log.Println("Error:",err)
	//}

	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, "Hello World!")
}

func HTTPS() {
	log.Println("Starting Https Server on " + httpsAddr)
	if err := http.ListenAndServeTLS(httpsAddr, "cert.pem", "key.pem", nil); err != nil {
		log.Fatal(err)
		return
	}


}

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	// Redirect the incoming HTTP request.
	fmt.Println("redirecting " + r.RemoteAddr + " to https")
	http.Redirect(w, r, "https://"+listenAddr+":9090"+r.RequestURI, http.StatusMovedPermanently)

}