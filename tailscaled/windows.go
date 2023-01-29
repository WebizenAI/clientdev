// I have a library that requires the tailscale Client deamon. 
// This package is used to install the tailscale client deamon on windows.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	"github.com/kardianos/service"
	"github.com/tailscale/tailscale"
	"github.com/tailscale/tailscale/win/tailscaled"
	"tailscale.com/client/tailscale/cli"
	"tailscale.com/ipn"
	"tailscale.com/ipn/ipnserver"
	"tailscale.com/logtail/backoff"
	"tailscale.com/version"
)

var (
	// These are set by the linker.
	// See https://golang.org/cmd/link/ for details.
	// They are used by the version subcommand.
	// The version subcommand is used by the Windows service.
	// The Windows service is used to install the tailscale client deamon.
	// The tailscale client deamon is used to connect to the tailscale network.
	// The tailscale network is used to connect to the tailscale network.

	// Version is the version of the tailscale client deamon.
	Version = "unknown"

	// GitCommit is the git commit of the tailscale client deamon.
	GitCommit = "unknown"

	// BuildDate is the build date of the tailscale client deamon.
	BuildDate = "unknown"

	// GoVersion is the go version of the tailscale client deamon.
	GoVersion = "unknown"

	// OS is the operating system of the tailscale client deamon.
	OS = "unknown"

	// Arch is the architecture of the tailscale client deamon.
	Arch = "unknown"

	// ServiceName is the name of the tailscale client deamon service.
	ServiceName = "Tailscale Client Deamon"

	// ServiceDescription is the description of the tailscale client deamon service.
	ServiceDescription = "Tailscale Client Deamon"

	// ServiceDependencies is the dependencies of the tailscale client deamon service.
	ServiceDependencies = []string{"EventLog"}

	// ServiceStartType is the start type of the tailscale client deamon service.
	ServiceStartType = service.StartAutomatic

	// ServiceAccount is the account of the tailscale client deamon service.
	ServiceAccount = service.AccountLocalSystem

	// ServiceAccountPassword is the account password of the tailscale client deamon service.
	ServiceAccountPassword = ""
)

// Program is the tailscale client deamon service.
type Program struct {
	// Exit is the exit channel.
	Exit chan struct{}
}

// Start is the start function of the tailscale client deamon service.
func (p *Program) Start(s service.Service) error {
	go p.run()
	return nil
}

// Stop is the stop function of the tailscale client deamon service.
func (p *Program) Stop(s service.Service) error {
	close(p.Exit)
	return nil
}

// run is the run function of the tailscale client deamon service.
func (p *Program) run() {
	// Start the tailscale client deamon.
	tailscaled.Main()
}

// main is the main function of the tailscale client deamon service.
func main() {
	// Create the program.
	program := &Program{
		Exit: make(chan struct{}),
	}

	// Create the service.
	svcConfig := &service.Config{
		Name:        ServiceName,
		DisplayName: ServiceName,
		Description: ServiceDescription,
		Arguments:   []string{"run"},
	}

	// Create the service.
	s, err := service.New(program, svcConfig)
	if err != nil {
		fmt.Printf("Failed to create service: %v

", err)
		os.Exit(1)
	}

	// Check if the service is running.
	if s.IsRunning() {
		fmt.Printf("Service is already running"), err)
		os.Exit(1)
	}

	// Check if the service is installed.
	if s.Installed() {
		fmt.Printf("Service is already installed"), err)
		os.Exit(1)
	}

	// Install the service.
	err = s.Install()
	if err != nil {
		fmt.Printf("Failed to install service: %v", err)
		os.Exit(1)
	}

	// Start the service.
	err = s.Start()
	if err != nil {
		fmt.Printf("Failed to start service: %v", err)
		os.Exit(1)
	}

	// Stop the service.
	err = s.Stop()
	if err != nil {
		fmt.Printf("Failed to stop service: %v", err)
		os.Exit(1)
	}

	// Uninstall the service.
	err = s.Uninstall()
	if err != nil {
		fmt.Printf("Failed to uninstall service: %v", err)
		os.Exit(1)
	}
}
