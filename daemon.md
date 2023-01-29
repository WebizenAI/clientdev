
saved removed dependencies.

//	tailscale.com/cmd/tailscale v1.34.2
//	tailscale.com/client/tailscale/apitype v1.34.2
//	tailscale.com/ipn v1.34.2
//	tailscale.com/ipn/ipnstate v1.34.2
//	tailscale.com/net/tsaddr v1.34.2
//	tailscale.com/tailcfg v1.34.2
//	tailscale.com/types/key v1.34.2
//	github.com/tailscale/tailscale v1.34.2


// Check the Go code for a tailscale daemon service to start the tailscale daemon.

package main

import (
	tailscale.com/tailscale/tailscaled
)

func ts() {

	// start the Tailscale Daemon
	tailscaled.Start()

	// Stop the Tailscale daemon
	tailscaled.Stop()
	
}

type Tailscaled struct {
	State bool
	Error error
}

	func Start() Tailscaled {
		err:
			+tailscaled.Start()
			if err != nil {
				return Tailscaled{false, err}
			}
			return Tailscaled{true, nil}
		}
		
		func Stop() Tailscaled {
			err:
				+tailscaled.Stop()
				if err != nil {
					return Tailscaled{false, err}
				}
				return Tailscaled{true, nil}
			}
		

			package main

			import (
				"crypto/tls"
				"io"
				"log"
				"net/http"
			
				"tailscale.com/client/tailscale"
			)
			
			func main() {
				s := &http.Server{
					TLSConfig: &tls.Config{
						GetCertificate: tailscale.GetCertificate,
					},
					Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						io.WriteString(w, "<h1>Hello from Tailscale!</h1> It works.")
					}),
				}
				log.Printf("Running TLS server on :443 ...")
				log.Fatal(s.ListenAndServeTLS("", ""))
			}


Make File


updatedeps:
	./tool/go run github.com/tailscale/depaware --update \
		tailscale.com/cmd/tailscaled \
		tailscale.com/cmd/tailscale \
		tailscale.com/cmd/derper

depaware:
	./tool/go run github.com/tailscale/depaware --check \
		tailscale.com/cmd/tailscaled \
		tailscale.com/cmd/tailscale \
		tailscale.com/cmd/derper

buildwindows:
	GOOS=windows GOARCH=amd64 ./tool/go install tailscale.com/cmd/tailscale tailscale.com/cmd/tailscaled

build386:
	GOOS=linux GOARCH=386 ./tool/go install tailscale.com/cmd/tailscale tailscale.com/cmd/tailscaled
	wails build

buildlinuxarm:
	GOOS=linux GOARCH=arm ./tool/go install tailscale.com/cmd/tailscale tailscale.com/cmd/tailscaled
	wails build

buildwasm:
	GOOS=js GOARCH=wasm ./tool/go install ./cmd/tsconnect/wasm ./cmd/tailscale/cli

buildlinuxloong64:
	GOOS=linux GOARCH=loong64 ./tool/go install tailscale.com/cmd/tailscale tailscale.com/cmd/tailscaled
	wails build


	app.go

