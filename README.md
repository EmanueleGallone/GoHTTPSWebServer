# GoHTTPSWebServer
Simple BoilerPlate to create a HTTPS webserver
creating private certificates and key (for testing purpose only)
using the generate_cert.go written by The Go Authors https://golang.org/src/crypto/tls/generate_cert.go

"go run main.go" will start a HTTP server on port 80 redirecting
to a HTTPS server on port 9090
