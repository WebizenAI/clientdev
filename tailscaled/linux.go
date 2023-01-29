// I have a library that requires the tailscale Client deamon.
// This package is used to install the tailscale client deamon on ubuntu.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kardianos/service"
	"tailscale.com/client/tailscale/cli"
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

	// ServiceUser is the user of the tailscale client deamon service.
	ServiceUser = ""

	// ServicePassword is the password of the tailscale client deamon service.
	ServicePassword = ""

	// ServiceInteractive is the interactive of the tailscale client deamon service.
	ServiceInteractive = false

	// ServiceStartType is the start type of the tailscale client deamon service.
	ServiceStartType = service.StartAutomatic

	// ServiceLogPath is the log path of the tailscale client deamon service.
	ServiceLogPath = ""

	// ServiceLogMode is the log mode of the tailscale client deamon service.
	ServiceLogMode = ""

	// ServiceLogMaxSize is the log max size of the tailscale client deamon service.
	ServiceLogMaxSize = 0

	// ServiceLogMaxBackups is the log max backups of the tailscale client deamon service.
	ServiceLogMaxBackups = 0

	// ServiceLogMaxAge is the log max age of the tailscale client deamon service.
	ServiceLogMaxAge = 0

	// ServiceLogCompress is the log compress of the tailscale client deamon service.
	ServiceLogCompress = false

	// ServiceLogLocalTime is the log local time of the tailscale client deamon service.
	ServiceLogLocalTime = false

	// ServiceLogUTC is the log utc of the tailscale client deamon service.
	ServiceLogUTC = false

	// ServiceLogShortFile is the log short file of the tailscale client deamon service.
	ServiceLogShortFile = false

	// ServiceLogStderr is the log stderr of the tailscale client deamon service.
	ServiceLogStderr = false

	// ServiceLogVerbose is the log verbose of the tailscale client deamon service.
	ServiceLogVerbose = false

	// ServiceLogDebug is the log debug of the tailscale client deamon service.
	ServiceLogDebug = false

	// ServiceLogTrace is the log trace of the tailscale client deamon service.
	ServiceLogTrace = false

	// ServiceLogJSON is the log json of the tailscale client deamon service.
	ServiceLogJSON = false

	// ServiceLogNoColor is the log no color of the tailscale client deamon service.
	ServiceLogNoColor = false

	// ServiceLogNoPrefix is the log no prefix of the tailscale client deamon service.
	ServiceLogNoPrefix = false

	// ServiceLogNoCaller is the log no caller of the tailscale client deamon service.
	ServiceLogNoCaller = false
)

// Service is the tailscale client deamon service.
type Service struct {
	// Service is the tailscale client deamon service.
	Service service.Service

	// Logger is the tailscale client deamon service logger.
	Logger service.Logger

	// Exit is the tailscale client deamon service exit.
	Exit chan struct{}
}

// Start is the tailscale client deamon service start.
func (s *Service) Start(s service.Service) error {
	s.Logger.Info("Starting tailscale client deamon service")
	go s.run()
	return nil
}

// Stop is the tailscale client deamon service stop.
func (s *Service) Stop(s service.Service) error {
	s.Logger.Info("Stopping tailscale client deamon service")
	close(s.Exit)
	return nil
}

// run is the tailscale client deamon service run.
func (s *Service) run() {
	s.Logger.Info("Running tailscale client deamon service")
}

// main is the tailscale client deamon main.
func main() {
	// Parse the flags.
	flag.Parse()

	// Get the arguments.
	args := flag.Args()

	// If the arguments are empty.
	if len(args) == 0 {
		// Print the usage.
		cli.Usage()
		os.Exit(1)
	}

	// Get the command.
	command := args[0]

	// If the command is version.
	if command == "version" {
		// Print the version.
		fmt.Printf("version: %s) (git commit: %s) (build date: %s) (go version: %s) (go arch: %s) (go os: %s)", Version, GitCommit, BuildDate, GoVersion, Arch, OS)
		os.Exit(0)
	}

	// If the command is install.
	if command == "install" {
		// Install the service.
		install()
		os.Exit(0)
	}

	// If the command is uninstall.
	if command == "uninstall" {
		// Uninstall the service.
		uninstall()
		os.Exit(0)
	}

	// If the command is start.
	if command == "start" {
		// Start the service.
		start()
		os.Exit(0)
	}

	// If the command is stop.
	if command == "stop" {
		// Stop the service.
		stop()
		os.Exit(0)
	}

	// If the command is restart.
	if command == "restart" {
		// Restart the service.
		restart()
		os.Exit(0)
	}

	// If the command is status.
	if command == "status" {
		// Status the service.
		status()
		os.Exit(0)
	}

	// If the command is run.
	if command == "run" {
		// Run the service.
		run()
		os.Exit(0)
	}

	// Print the usage.
	cli.Usage()
	os.Exit(1)
}
