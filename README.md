# GoHTTPSWebServer
Simple BoilerPlate to create an HTTTPS webserver
using private certificates and key (for testing purpose only)
generated from the generate_cert.go written by The Go Authors https://golang.org/src/crypto/tls/generate_cert.go

using "go run main.go" will start an HTTP server on port 80 redirecting
to an HTTPS server on port 9090
