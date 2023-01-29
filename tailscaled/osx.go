// I have a library that requires the tailscale Client deamon.
// This package is used to install the tailscale client deamon on OSX.

package main

import (
	"fmt"
	"io"
	"log"
	"log/syslog"
	"os"
	"runtime"
	"time"

	"github.com/kardianos/service"
	"tailscale.com/ipn/ipnserver"
)

var (
	// These are set by the linker.
	// See https://golang.org/cmd/link/ for details.
	// They are used by the version subcommand.
	// The version subcommand is used by the OSX service.
	// The OSX service is used to install the tailscale client deamon.
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

	// ServiceLog is the log of the tailscale client deamon service.
	ServiceLog = ""

	// ServiceLogFile is the log file of the tailscale client deamon service.
	ServiceLogFile = ""

	// ServiceLogPrefix is the log prefix of the tailscale client deamon service.
	ServiceLogPrefix = ""

	// ServiceLogFlags is the log flags of the tailscale client deamon service.
	ServiceLogFlags = 0

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

	// ServiceLogLongFile is the log long file of the tailscale client deamon service.
	ServiceLogLongFile = false

	// ServiceLogMicroseconds is the log microseconds of the tailscale client deamon service.
	ServiceLogMicroseconds = false

	// ServiceLogMilliseconds is the log milliseconds of the tailscale client deamon service.
	ServiceLogMilliseconds = false

	// ServiceLogUTCDate is the log utc date of the tailscale client deamon service.
	ServiceLogUTCDate = false

	// ServiceLogFullTimestamp is the log full timestamp of the tailscale client deamon service.
	ServiceLogFullTimestamp = false

	// ServiceLogCallerPrettyfier is the log caller prettyfier of the tailscale client deamon service.
	ServiceLogCallerPrettyfier = nil

	// ServiceLogLevel is the log level of the tailscale client deamon service.
	ServiceLogLevel = 0

	// ServiceLogFunc is the log func of the tailscale client deamon service.
	ServiceLogFunc = nil

	// ServiceLogFuncCallDepth is the log func call depth of the tailscale client deamon service.
	ServiceLogFuncCallDepth = 0

	// ServiceLogFuncCallSkip is the log func call skip of the tailscale client deamon service.
	ServiceLogFuncCallSkip = 0

	// ServiceLogFuncCallSkipFrame is the log func call skip frame of the tailscale client deamon service.
	ServiceLogFuncCallSkipFrame = 0

	// ServiceLogFuncCallSkipPC is the log func call skip pc of the tailscale client deamon service.
	ServiceLogFuncCallSkipPC = 0

	// ServiceLogFuncCallSkipFile is the log func call skip file of the tailscale client deamon service.
	ServiceLogFuncCallSkipFile = 0

	// ServiceLogFuncCallSkipLine is the log func call skip line of the tailscale client deamon service.
	ServiceLogFuncCallSkipLine = 0

	// ServiceLogFuncCallSkipFunc is the log func call skip func of the tailscale client deamon service.
	ServiceLogFuncCallSkipFunc = 0

	// ServiceLogFuncCallSkipName is the log func call skip name of the tailscale client deamon service.
	ServiceLogFuncCallSkipName = 0

	// ServiceLogFuncCallSkipEntry is the log func call skip entry of the tailscale client deamon service.
	ServiceLogFuncCallSkipEntry = 0

	// ServiceLogFuncCallSkipAll is the log func call skip all of the tailscale client deamon service.
	ServiceLogFuncCallSkipAll = 0

	// ServiceLogFuncCallSkipNone is the log func call skip none of the tailscale client deamon service.
	ServiceLogFuncCallSkipNone = 0

	// ServiceLogFuncCallSkipDefault is the log func call skip default of the tailscale client deamon service.
	ServiceLogFuncCallSkipDefault = 0
)

// Service is the tailscale client deamon service.
type Service struct {
	// Config is the configuration of the tailscale client deamon service.
	Config *ipnserver.Config `json:"config"` // Config is the configuration of the tailscale client deamon service.

	// Log is the log of the tailscale client deamon service.
	Log *log.Logger `json:"log"` // Log is the log of the tailscale client deamon service.

	// LogFile is the log file of the tailscale client deamon service.
	LogFile *os.File `json:"log_file"` // LogFile is the log file of the tailscale client deamon service.

	// LogFile is the log file of the tailscale client deamon service.
	LogFile *os.File `json:"log_file"` // LogFile is the log file of the tailscale client deamon service.

}

// NewService returns a new tailscale client deamon service.
func NewService(config *ipnserver.Config) (*Service, error) {
	return &Service{
		Config: config,
	}, nil
}

// Run runs the tailscale client deamon service.
func (s *Service) Run() error {
	// Create the log file.
	logFile, err := os.OpenFile(ServiceLogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	// Create the log.
	log := log.New(logFile, ServiceLogPrefix, ServiceLogFlags)

	// Create the service.
	service, err := service.New(ServiceName, ServiceDescription, ServiceDependencies, ServiceUser, ServicePassword, ServiceInteractive, ServiceStartType, ServiceLog, ServiceLogFile, ServiceLogPrefix, ServiceLogFlags, ServiceLogMaxSize, ServiceLogMaxBackups, ServiceLogMaxAge, ServiceLogCompress, ServiceLogLocalTime, ServiceLogUTC, ServiceLogShortFile, ServiceLogLongFile, ServiceLogMicroseconds, ServiceLogMilliseconds, ServiceLogUTCDate, ServiceLogFullTimestamp, ServiceLogCallerPrettyfier, ServiceLogLevel, ServiceLogFunc, ServiceLogFuncCallDepth, ServiceLogFuncCallSkip, ServiceLogFuncCallSkipFrame, ServiceLogFuncCallSkipPC, ServiceLogFuncCallSkipFile, ServiceLogFuncCallSkipLine, ServiceLogFuncCallSkipFunc, ServiceLogFuncCallSkipName, ServiceLogFuncCallSkipEntry, ServiceLogFuncCallSkipAll, ServiceLogFuncCallSkipNone, ServiceLogFuncCallSkipDefault)
	if err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}

	// Run the service.
	err = service.Run(s)
	if err != nil {
		return fmt.Errorf("failed to run service: %w", err)
	}

	return nil
}

// Start starts the tailscale client deamon service.
func (s *Service) Start() error {
	// Create the log file.
	logFile, err := os.OpenFile(ServiceLogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	// Create the log.
	log := log.New(logFile, ServiceLogPrefix, ServiceLogFlags)

	// Create the server.
	server, err := ipnserver.New(s.Config)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}

	// Start the server.
	err = server.Start()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

// Stop stops the tailscale client deamon service.
func (s *Service) Stop() error {
	return nil
}

// Install installs the tailscale client deamon service.
func Install() error {
	// Create the service.
	service, err := service.New(ServiceName, ServiceDescription, ServiceDependencies, ServiceUser, ServicePassword, ServiceInteractive, ServiceStartType, ServiceLog, ServiceLogFile, ServiceLogPrefix, ServiceLogFlags, ServiceLogMaxSize, ServiceLogMaxBackups, ServiceLogMaxAge, ServiceLogCompress, ServiceLogLocalTime, ServiceLogUTC, ServiceLogShortFile, ServiceLogLongFile, ServiceLogMicroseconds, ServiceLogMilliseconds, ServiceLogUTCDate, ServiceLogFullTimestamp, ServiceLogCallerPrettyfier, ServiceLogLevel, ServiceLogFunc, ServiceLogFuncCallDepth, ServiceLogFuncCallSkip, ServiceLogFuncCallSkipFrame, ServiceLogFuncCallSkipPC, ServiceLogFuncCallSkipFile, ServiceLogFuncCallSkipLine, ServiceLogFuncCallSkipFunc, ServiceLogFuncCallSkipName, ServiceLogFuncCallSkipEntry, ServiceLogFuncCallSkipAll, ServiceLogFuncCallSkipNone, ServiceLogFuncCallSkipDefault)
	if err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}

	// Install the service.
	err = service.Install()
	if err != nil {
		return fmt.Errorf("failed to install service: %w", err)
	}

	return nil
}

// Uninstall uninstalls the tailscale client deamon service.
func Uninstall() error {
	// Create the service.
	service, err := service.New(ServiceName, ServiceDescription, ServiceDependencies, ServiceUser, ServicePassword, ServiceInteractive, ServiceStartType, ServiceLog, ServiceLogFile, ServiceLogPrefix, ServiceLogFlags, ServiceLogMaxSize, ServiceLogMaxBackups, ServiceLogMaxAge, ServiceLogCompress, ServiceLogLocalTime, ServiceLogUTC, ServiceLogShortFile, ServiceLogLongFile, ServiceLogMicroseconds, ServiceLogMilliseconds, ServiceLogUTCDate, ServiceLogFullTimestamp, ServiceLogCallerPrettyfier, ServiceLogLevel, ServiceLogFunc, ServiceLogFuncCallDepth, ServiceLogFuncCallSkip, ServiceLogFuncCallSkipFrame, ServiceLogFuncCallSkipPC, ServiceLogFuncCallSkipFile, ServiceLogFuncCallSkipLine, ServiceLogFuncCallSkipFunc, ServiceLogFuncCallSkipName, ServiceLogFuncCallSkipEntry, ServiceLogFuncCallSkipAll, ServiceLogFuncCallSkipNone, ServiceLogFuncCallSkipDefault)
	if err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}

	// Uninstall the service.
	err = service.Uninstall()
	if err != nil {
		return fmt.Errorf("failed to uninstall service: %w", err)
	}

	return nil
}

// Start starts the tailscale client deamon service.
func Start() error {
	// Create the service.
	service, err := service.New(ServiceName, ServiceDescription, ServiceDependencies, ServiceUser, ServicePassword, ServiceInteractive, ServiceStartType, ServiceLog, ServiceLogFile, ServiceLogPrefix, ServiceLogFlags, ServiceLogMaxSize, ServiceLogMaxBackups, ServiceLogMaxAge, ServiceLogCompress, ServiceLogLocalTime, ServiceLogUTC, ServiceLogShortFile, ServiceLogLongFile, ServiceLogMicroseconds, ServiceLogMilliseconds, ServiceLogUTCDate, ServiceLogFullTimestamp, ServiceLogCallerPrettyfier, ServiceLogLevel, ServiceLogFunc, ServiceLogFuncCallDepth, ServiceLogFuncCallSkip, ServiceLogFuncCallSkipFrame, ServiceLogFuncCallSkipPC, ServiceLogFuncCallSkipFile, ServiceLogFuncCallSkipLine, ServiceLogFuncCallSkipFunc, ServiceLogFuncCallSkipName, ServiceLogFuncCallSkipEntry, ServiceLogFuncCallSkipAll, ServiceLogFuncCallSkipNone, ServiceLogFuncCallSkipDefault)
	if err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}

	// Start the service.
	err = service.Start()
	if err != nil {
		return fmt.Errorf("failed to start service: %w", err)
	}

	return nil
}

// Stop stops the tailscale client deamon service.
func Stop() error {
	// Create the service.
	service, err := service.New(ServiceName, ServiceDescription, ServiceDependencies, ServiceUser, ServicePassword, ServiceInteractive, ServiceStartType, ServiceLog, ServiceLogFile, ServiceLogPrefix, ServiceLogFlags, ServiceLogMaxSize, ServiceLogMaxBackups, ServiceLogMaxAge, ServiceLogCompress, ServiceLogLocalTime, ServiceLogUTC, ServiceLogShortFile, ServiceLogLongFile, ServiceLogMicroseconds, ServiceLogMilliseconds, ServiceLogUTCDate, ServiceLogFullTimestamp, ServiceLogCallerPrettyfier, ServiceLogLevel, ServiceLogFunc, ServiceLogFuncCallDepth, ServiceLogFuncCallSkip, ServiceLogFuncCallSkipFrame, ServiceLogFuncCallSkipPC, ServiceLogFuncCallSkipFile, ServiceLogFuncCallSkipLine, ServiceLogFuncCallSkipFunc, ServiceLogFuncCallSkipName, ServiceLogFuncCallSkipEntry, ServiceLogFuncCallSkipAll, ServiceLogFuncCallSkipNone, ServiceLogFuncCallSkipDefault)
	if err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}

	// Stop the service.
	err = service.Stop()
	if err != nil {
		return fmt.Errorf("failed to stop service: %w", err)
	}

	return nil
}

// Status returns the status of the tailscale client deamon service.
func Status() (string, error) {
	// Create the service.
	service, err := service.New(ServiceName, ServiceDescription, ServiceDependencies, ServiceUser, ServicePassword, ServiceInteractive, ServiceStartType, ServiceLog, ServiceLogFile, ServiceLogPrefix, ServiceLogFlags, ServiceLogMaxSize, ServiceLogMaxBackups, ServiceLogMaxAge, ServiceLogCompress, ServiceLogLocalTime, ServiceLogUTC, ServiceLogShortFile, ServiceLogLongFile, ServiceLogMicroseconds, ServiceLogMilliseconds, ServiceLogUTCDate, ServiceLogFullTimestamp, ServiceLogCallerPrettyfier, ServiceLogLevel, ServiceLogFunc, ServiceLogFuncCallDepth, ServiceLogFuncCallSkip, ServiceLogFuncCallSkipFrame, ServiceLogFuncCallSkipPC, ServiceLogFuncCallSkipFile, ServiceLogFuncCallSkipLine, ServiceLogFuncCallSkipFunc, ServiceLogFuncCallSkipName, ServiceLogFuncCallSkipEntry, ServiceLogFuncCallSkipAll, ServiceLogFuncCallSkipNone, ServiceLogFuncCallSkipDefault)
	if err != nil {
		return "", fmt.Errorf("failed to create service: %w", err)
	}

	// Get the status of the service.
	status, err := service.Status()
	if err != nil {
		return "", fmt.Errorf("failed to get status of service: %w", err)
	}

	return status, nil
}

// Run runs the tailscale client deamon service.
func Run() error {
	// Create the service.
	service, err := service.New(ServiceName, ServiceDescription, ServiceDependencies, ServiceUser, ServicePassword, ServiceInteractive, ServiceStartType, ServiceLog, ServiceLogFile, ServiceLogPrefix, ServiceLogFlags, ServiceLogMaxSize, ServiceLogMaxBackups, ServiceLogMaxAge, ServiceLogCompress, ServiceLogLocalTime, ServiceLogUTC, ServiceLogShortFile, ServiceLogLongFile, ServiceLogMicroseconds, ServiceLogMilliseconds, ServiceLogUTCDate, ServiceLogFullTimestamp, ServiceLogCallerPrettyfier, ServiceLogLevel, ServiceLogFunc, ServiceLogFuncCallDepth, ServiceLogFuncCallSkip, ServiceLogFuncCallSkipFrame, ServiceLogFuncCallSkipPC, ServiceLogFuncCallSkipFile, ServiceLogFuncCallSkipLine, ServiceLogFuncCallSkipFunc, ServiceLogFuncCallSkipName, ServiceLogFuncCallSkipEntry, ServiceLogFuncCallSkipAll, ServiceLogFuncCallSkipNone, ServiceLogFuncCallSkipDefault)
	if err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}

	// Run the service.
	err = service.Run()
	if err != nil {
		return fmt.Errorf("failed to run service: %w", err)
	}

	return nil
}

// ServiceName is the name of the tailscale client deamon service.
const ServiceName = "tailscale"

// ServiceDescription is the description of the tailscale client deamon service.
const ServiceDependencies = "tailscale"

// ServiceUser is the user of the tailscale client deamon service.
const ServiceUser = "tailscale"

// ServicePassword is the password of the tailscale client deamon service.
const ServicePassword = "tailscale"

// ServiceInteractive is the interactive of the tailscale client deamon service.
const ServiceInteractive = "tailscale"

// ServiceStartType is the start type of the tailscale client deamon service.
const ServiceStartType = "tailscale"

// ServiceLog is the log of the tailscale client deamon service.
const ServiceLog = "tailscale"

// ServiceLogFile is the log file of the tailscale client deamon service.
const ServiceLogFile = "tailscale"

// ServiceLogPrefix is the log prefix of the tailscale client deamon service.
const ServiceLogPrefix = "tailscale"

// ServiceLogFlags is the log flags of the tailscale client deamon service.
const ServiceLogFlags = "tailscale"

// ServiceLogMaxSize is the log max size of the tailscale client deamon service.
const ServiceLogMaxSize = "tailscale"

// ServiceLogMaxBackups is the log max backups of the tailscale client deamon service.
const ServiceLogMaxBackups = "tailscale"

// ServiceLogMaxAge is the log max age of the tailscale client deamon service.
const ServiceLogMaxAge = "tailscale"

// ServiceLogCompress is the log compress of the tailscale client deamon service.
const ServiceLogCompress = "tailscale"

// ServiceLogLocalTime is the log local time of the tailscale client deamon service.
const ServiceLogLocalTime = "tailscale"

// ServiceLogUTC is the log utc of the tailscale client deamon service.
const ServiceLogUTC = "tailscale"

// ServiceLogShortFile is the log short file of the tailscale client deamon service.
const ServiceLogShortFile = "tailscale"

// ServiceLogLongFile is the log long file of the tailscale client deamon service.
const ServiceLogLongFile = "tailscale"

// ServiceLogMicroseconds is the log microseconds of the tailscale client deamon service.
const ServiceLogMicroseconds = "tailscale"

// ServiceLogMilliseconds is the log milliseconds of the tailscale client deamon service.
const ServiceLogMilliseconds = "tailscale"

// ServiceLogUTCDate is the log utc date of the tailscale client deamon service.
const ServiceLogUTCDate = "tailscale"

// ServiceLogFullTimestamp is the log full timestamp of the tailscale client deamon service.
const ServiceLogFullTimestamp = "tailscale"

// ServiceLogCallerPrettyfier is the log caller prettyfier of the tailscale client deamon service.
const ServiceLogCallerPrettyfier = "tailscale"

// ServiceLogLevel is the log level of the tailscale client deamon service.
const ServiceLogLevel = "tailscale"

// ServiceLogFunc is the log func of the tailscale client deamon service.
const ServiceLogFunc = "tailscale"

// ServiceLogFuncCallDepth is the log func call depth of the tailscale client deamon service.
const ServiceLogFuncCallDepth = "tailscale"

// ServiceLogFuncCallSkip is the log func call skip of the tailscale client deamon service.
const ServiceLogFuncCallSkip = "tailscale"

// ServiceLogFuncCallSkipFrame is the log func call skip frame of the tailscale client deamon service.
const ServiceLogFuncCallSkipFrame = "tailscale"

// ServiceLogFuncCallSkipPC is the log func call skip pc of the tailscale client deamon service.
const ServiceLogFuncCallSkipPC = "tailscale"

// ServiceLogFuncCallSkipFile is the log func call skip file of the tailscale client deamon service.
const ServiceLogFuncCallSkipFile = "tailscale"

// ServiceLogFuncCallSkipLine is the log func call skip line of the tailscale client deamon service.
const ServiceLogFuncCallSkipLine = "tailscale"

// ServiceLogFuncCallSkipFunc is the log func call skip func of the tailscale client deamon service.
const ServiceLogFuncCallSkipFunc = "tailscale"

// ServiceLogFuncCallSkipName is the log func call skip name of the tailscale client deamon service.
const ServiceLogFuncCallSkipName = "tailscale"

// ServiceLogFuncCallSkipEntry is the log func call skip entry of the tailscale client deamon service.
const ServiceLogFuncCallSkipEntry = "tailscale"

// ServiceLogFuncCallSkipAll is the log func call skip all of the tailscale client deamon service.
const ServiceLogFuncCallSkipAll = "tailscale"

// ServiceLogFuncCallSkipNone is the log func call skip none of the tailscale client deamon service.
const ServiceLogFuncCallSkipNone = "tailscale"

// ServiceLogFuncCallSkipDefault is the log func call skip default of the tailscale client deamon service.
const ServiceLogFuncCallSkipDefault = "tailscale"

// Service is the tailscale client deamon service.
type Service struct {
	// Name is the name of the service.
	Name string

	// Description is the description of the service.
	Description string

	// Dependencies is the dependencies of the service.
	Dependencies []string

	// User is the user of the service.
	User string

	// Password is the password of the service.
	Password string

	// Interactive is the interactive of the service.
	Interactive bool

	// StartType is the start type of the service.
	StartType StartType

	// Log is the log of the service.
	Log *log.Logger

	// LogFile is the log file of the service.
	LogFile string

	// LogPrefix is the log prefix of the service.
	LogPrefix string

	// LogFlags is the log flags of the service.
	LogFlags int

	// LogMaxSize is the log max size of the service.
	LogMaxSize int

	// LogMaxBackups is the log max backups of the service.
	LogMaxBackups int

	// LogMaxAge is the log max age of the service.
	LogMaxAge int

	// LogCompress is the log compress of the service.
	LogCompress bool

	// LogLocalTime is the log local time of the service.
	LogLocalTime bool

	// LogUTC is the log utc of the service.
	LogUTC bool

	// LogShortFile is the log short file of the service.
	LogShortFile bool

	// LogLongFile is the log long file of the service.
	LogLongFile bool

	// LogMicroseconds is the log microseconds of the service.
	LogMicroseconds bool

	// LogMilliseconds is the log milliseconds of the service.
	LogMilliseconds bool

	// LogUTCDate is the log utc date of the service.
	LogUTCDate bool

	// LogFullTimestamp is the log full timestamp of the service.
	LogFullTimestamp bool

	// LogCallerPrettyfier is the log caller prettyfier of the service.
	LogCallerPrettyfier func(f *runtime.Frame) (function string, file string)

	// LogLevel is the log level of the service.
	LogLevel log.Level

	// LogFunc is the log func of the service.
	LogFunc func(args ...interface{})

	// LogFuncCallDepth is the log func call depth of the service.
	LogFuncCallDepth int

	// LogFuncCallSkip is the log func call skip of the service.
	LogFuncCallSkip int

	// LogFuncCallSkipFrame is the log func call skip frame of the service.
	LogFuncCallSkipFrame int

	// LogFuncCallSkipPC is the log func call skip pc of the service.
	LogFuncCallSkipPC int

	// LogFuncCallSkipFile is the log func call skip file of the service.
	LogFuncCallSkipFile int

	// LogFuncCallSkipLine is the log func call skip line of the service.
	LogFuncCallSkipLine int

	// LogFuncCallSkipFunc is the log func call skip func of the service.
	LogFuncCallSkipFunc int

	// LogFuncCallSkipName is the log func call skip name of the service.
	LogFuncCallSkipName int

	// LogFuncCallSkipEntry is the log func call skip entry of the service.
	LogFuncCallSkipEntry int

	// LogFuncCallSkipAll is the log func call skip all of the service.
	LogFuncCallSkipAll int

	// LogFuncCallSkipNone is the log func call skip none of the service.
	LogFuncCallSkipNone int

	// LogFuncCallSkipDefault is the log func call skip default of the service.
	LogFuncCallSkipDefault int

	// Service is the service of the service.
	Service *service.Service

	// ServiceConfig is the service config of the service.
	ServiceConfig *service.Config

	// ServiceLogger is the service logger of the service.
	ServiceLogger *service.Logger

	// ServiceLoggerConfig is the service logger config of the service.
	ServiceLoggerConfig *service.LoggerConfig

	// ServiceLoggerConfigFile is the service logger config file of the service.
	ServiceLoggerConfigFile *service.LoggerConfigFile

	// ServiceLoggerConfigSyslog is the service logger config syslog of the service.
	ServiceLoggerConfigSyslog *service.LoggerConfigSyslog

	// ServiceLoggerConfigEventlog is the service logger config eventlog of the service.
	ServiceLoggerConfigEventlog *service.LoggerConfigEventlog

	// ServiceLoggerConfigJournal is the service logger config journal of the service.
	ServiceLoggerConfigJournal *service.LoggerConfigJournal

	// ServiceLoggerConfigStdout is the service logger config stdout of the service.
	ServiceLoggerConfigStdout *service.LoggerConfigStdout

	// ServiceLoggerConfigStderr is the service logger config stderr of the service.
	ServiceLoggerConfigStderr *service.LoggerConfigStderr

	// ServiceLoggerConfigFileWriter is the service logger config file writer of the service.
	ServiceLoggerConfigFileWriter *service.LoggerConfigFileWriter

	// ServiceLoggerConfigSyslogWriter is the service logger config syslog writer of the service.
	ServiceLoggerConfigSyslogWriter *service.LoggerConfigSyslogWriter

	// ServiceLoggerConfigEventlogWriter is the service logger config eventlog writer of the service.
	ServiceLoggerConfigEventlogWriter *service.LoggerConfigEventlogWriter

	// ServiceLoggerConfigJournalWriter is the service logger config journal writer of the service.
	ServiceLoggerConfigJournalWriter *service.LoggerConfigJournalWriter

	// ServiceLoggerConfigStdoutWriter is the service logger config stdout writer of the service.
	ServiceLoggerConfigStdoutWriter *service.LoggerConfigStdoutWriter

	// ServiceLoggerConfigStderrWriter is the service logger config stderr writer of the service.
	ServiceLoggerConfigStderrWriter *service.LoggerConfigStderrWriter
}

// NewService returns a new service.
func NewService() *Service {
	errorLog := log.New(os.Stderr, "", log.LstdFlags)
	errorLog.SetOutput(os.Stderr)
	errorLog.SetFlags(log.LstdFlags)
	errorLog.SetPrefix("")
	errorLog.SetLevel(log.ErrorLevel)
	errorLog.SetFormatter(&log.TextFormatter{
		DisableColors:          false,
		DisableTimestamp:       false,
		DisableLevelTruncation: false,
		FullTimestamp:          false,
		TimestampFormat:        time.RFC3339,
		ForceColors:            false,
		ForceQuote:             false,
		DisableSorting:         false,
		DisableQuote:           false,
		QuoteEmptyFields:       false,
		IsTerminal:             false,
		DisableLevelTruncation: false,
		PadLevelText:           false,
		TrimMessages:           false,
		CallerPrettyfier:       nil,
	})
	errorLog.SetReportCaller(false)
	errorLog.SetNoLock()
	errorLog.SetExitFunc(os.Exit)
	errorLog.SetHook(&log.LvlFilterHandler{
		MinLevel: log.ErrorLevel,
		MaxLevel: log.FatalLevel,
		Handler:  log.DiscardHandler(),
	})
	errorLog.SetOutput(&log.LvlFilterHandler{
		MinLevel: log.ErrorLevel,
		MaxLevel: log.FatalLevel,
		Handler:  log.DiscardHandler(),
	})
	errorLog.SetOutput(&log.LvlFilterHandler{
		MinLevel: log.ErrorLevel,
		MaxLevel: log.FatalLevel,
		Handler:  log.DiscardHandler(),
	})

	infoLog := log.New(os.Stdout, "", log.LstdFlags)
	infoLog.SetOutput(os.Stdout)
	infoLog.SetFlags(log.LstdFlags)
	infoLog.SetPrefix("")
	infoLog.SetLevel(log.InfoLevel)
	infoLog.SetFormatter(&log.TextFormatter{
		DisableColors:          false,
		DisableTimestamp:       false,
		DisableLevelTruncation: false,
		FullTimestamp:          false,
		TimestampFormat:        time.RFC3339,
		ForceColors:            false,
		ForceQuote:             false,
		DisableSorting:         false,
		DisableQuote:           false,
		QuoteEmptyFields:       false,
		IsTerminal:             false,
		DisableLevelTruncation: false,
		PadLevelText:           false,
		TrimMessages:           false,
		CallerPrettyfier:       nil,
	})
	infoLog.SetReportCaller(false)
	infoLog.SetNoLock()
	infoLog.SetExitFunc(os.Exit)
	infoLog.SetHook(&log.LvlFilterHandler{
		MinLevel: log.InfoLevel,
		MaxLevel: log.InfoLevel,
		Handler:  log.DiscardHandler(),
	})
	infoLog.SetOutput(&log.LvlFilterHandler{
		MinLevel: log.InfoLevel,
		MaxLevel: log.InfoLevel,
		Handler:  log.DiscardHandler(),
	})
	infoLog.SetOutput(&log.LvlFilterHandler{
		MinLevel: log.InfoLevel,
		MaxLevel: log.InfoLevel,
		Handler:  log.DiscardHandler(),
	})

	debugLog := log.New(os.Stdout, "", log.LstdFlags)
	debugLog.SetOutput(os.Stdout)
	debugLog.SetFlags(log.LstdFlags)
	debugLog.SetPrefix("")
	debugLog.SetLevel(log.DebugLevel)
	debugLog.SetFormatter(&log.TextFormatter{
		DisableColors:          false,
		DisableTimestamp:       false,
		DisableLevelTruncation: false,
		FullTimestamp:          false,
		TimestampFormat:        time.RFC3339,
		ForceColors:            false,
		ForceQuote:             false,
		DisableSorting:         false,
		DisableQuote:           false,
		QuoteEmptyFields:       false,
		IsTerminal:             false,
		DisableLevelTruncation: false,
		PadLevelText:           false,
		TrimMessages:           false,
		CallerPrettyfier:       nil,
	})
	debugLog.SetReportCaller(false)
	debugLog.SetNoLock()
	debugLog.SetExitFunc(os.Exit)
	debugLog.SetHook(&log.LvlFilterHandler{
		MinLevel: log.DebugLevel,
		MaxLevel: log.DebugLevel,
		Handler:  log.DiscardHandler(),
	})
	debugLog.SetOutput(&log.LvlFilterHandler{
		MinLevel: log.DebugLevel,
		MaxLevel: log.DebugLevel,
		Handler:  log.DiscardHandler(),
	})
	debugLog.SetOutput(&log.LvlFilterHandler{
		MinLevel: log.DebugLevel,
		MaxLevel: log.DebugLevel,
		Handler:  log.DiscardHandler(),
	})

	return &Service{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		DebugLog: debugLog,
	}
}

// Run runs the service.
func (s *Service) Run() error {
	s.DebugLog.Debug("service run")

	return nil
}

// Stop stops the service.
func (s *Service) Stop() error {
	s.DebugLog.Debug("service stop")

	return nil
}

// ServiceLoggerConfigStdoutWriter is the service logger config stdout writer of the service.
type ServiceLoggerConfigStdoutWriter struct {
	// MinLevel is the minimum level of the service logger config stdout writer of the service.
	MinLevel log.Level

	// MaxLevel is the maximum level of the service logger config stdout writer of the service.
	MaxLevel log.Level

	// Handler is the handler of the service logger config stdout writer of the service.
	Handler log.Handler

	// Formatter is the formatter of the service logger config stdout writer of the service.
	Formatter log.Formatter

	// Output is the output of the service logger config stdout writer of the service.
	Output io.Writer

	// ExitFunc is the exit function of the service logger config stdout writer of the service.
	ExitFunc func(int)

	// ReportCaller is the report caller of the service logger config stdout writer of the service.
	ReportCaller bool

	// Hooks is the hooks of the service logger config stdout writer of the service.
	Hooks log.LevelHooks

	// Level is the level of the service logger config stdout writer of the service.
	Level log.Level

	// NoLock is the no lock of the service logger config stdout writer of the service.
	NoLock bool

	// Flags is the flags of the service logger config stdout writer of the service.
	Flags int

	// Prefix is the prefix of the service logger config stdout writer of the service.
	Prefix string
}

// ServiceLoggerConfigStderrWriter is the service logger config stderr writer of the service.
type ServiceLoggerConfigStderrWriter struct {
	// MinLevel is the minimum level of the service logger config stderr writer of the service.
	MinLevel log.Level

	// MaxLevel is the maximum level of the service logger config stderr writer of the service.
	MaxLevel log.Level

	// Handler is the handler of the service logger config stderr writer of the service.
	Handler log.Handler

	// Formatter is the formatter of the service logger config stderr writer of the service.
	Formatter log.Formatter

	// Output is the output of the service logger config stderr writer of the service.
	Output io.Writer

	// ExitFunc is the exit function of the service logger config stderr writer of the service.
	ExitFunc func(int)

	// ReportCaller is the report caller of the service logger config stderr writer of the service.
	ReportCaller bool

	// Hooks is the hooks of the service logger config stderr writer of the service.
	Hooks log.LevelHooks

	// Level is the level of the service logger config stderr writer of the service.
	Level log.Level

	// NoLock is the no lock of the service logger config stderr writer of the service.
	NoLock bool

	// Flags is the flags of the service logger config stderr writer of the service.
	Flags int

	// Prefix is the prefix of the service logger config stderr writer of the service.
	Prefix string
}

// ServiceLoggerConfig is the service logger config of the service.
type ServiceLoggerConfig struct {
	// StdoutWriter is the stdout writer of the service logger config of the service.
	StdoutWriter ServiceLoggerConfigStdoutWriter

	// StderrWriter is the stderr writer of the service logger config of the service.
	StderrWriter ServiceLoggerConfigStderrWriter
}

// ServiceConfig is the service config of the service.
type ServiceConfig struct {
	// Logger is the logger of the service config of the service.
	Logger ServiceLoggerConfig
}

// Service is the service of the service.
type Service struct {
	// ErrorLog is the error log of the service of the service.
	ErrorLog *log.Logger

	// InfoLog is the info log of the service of the service.
	InfoLog *log.Logger

	// DebugLog is the debug log of the service of the service.
	DebugLog *log.Logger
}

// NewService creates a new service.
func NewService(config ServiceConfig) *Service {
	errorLog := log.New()
	errorLog.SetOutput(os.Stderr)
	errorLog.SetFlags(log.LstdFlags)
	errorLog.SetPrefix("")
	errorLog.SetLevel(log.ErrorLevel)
	errorLog.SetFormatter(&log.TextFormatter{
		DisableColors:          false,
		DisableTimestamp:       false,
		DisableLevelTruncation: false,
		FullTimestamp:          false,
		TimestampFormat:        time.RFC3339,
		ForceColors:            false,
		ForceQuote:             false,
		DisableSorting:         false,
		DisableQuote:           false,
		QuoteEmptyFields:       false,
		IsTerminal:             false,
		DisableLevelTruncation: false,
		PadLevelText:           false,
		TrimMessages:           false,
		CallerPrettyfier:       nil,
	})
	errorLog.SetReportCaller(false)
	errorLog.SetNoLock()
	errorLog.SetExitFunc(os.Exit)
	errorLog.SetHook(&log.LvlFilterHandler{
		MinLevel: log.ErrorLevel,
		MaxLevel: log.ErrorLevel,
		Handler:  log.DiscardHandler(),
	})
	errorLog.SetOutput(&log.LvlFilterHandler{
		MinLevel: log.ErrorLevel,
		MaxLevel: log.ErrorLevel,
		Handler:  log.DiscardHandler(),
	})
	errorLog.SetOutput(&log.LvlFilterHandler{
		MinLevel: log.ErrorLevel,
		MaxLevel: log.ErrorLevel,
		Handler:  log.DiscardHandler(),
	})

	infoLog := log.New()
	infoLog.SetOutput(os.Stdout)
	infoLog.SetFlags(log.LstdFlags)
	infoLog.SetPrefix("")
	infoLog.SetLevel(log.InfoLevel)
	infoLog.SetFormatter(&log.TextFormatter{
		DisableColors:          false,
		DisableTimestamp:       false,
		DisableLevelTruncation: false,
		FullTimestamp:          false,
		TimestampFormat:        time.RFC3339,
		ForceColors:            false,
		ForceQuote:             false,
		DisableSorting:         false,
		DisableQuote:           false,
		QuoteEmptyFields:       false,
		IsTerminal:             false,
		DisableLevelTruncation: false,
		PadLevelText:           false,
		TrimMessages:           false,
		CallerPrettyfier:       nil,
	})
	infoLog.SetReportCaller(false)
	infoLog.SetNoLock()
	infoLog.SetExitFunc(os.Exit)
	infoLog.SetHook(&log.LvlFilterHandler{
		MinLevel: log.InfoLevel,
		MaxLevel: log.InfoLevel,
		Handler:  log.DiscardHandler(),
	})
	infoLog.SetOutput(&log.LvlFilterHandler{
		MinLevel: log.InfoLevel,
		MaxLevel: log.InfoLevel,
		Handler:  log.DiscardHandler(),
	})
	infoLog.SetOutput(&log.LvlFilterHandler{
		MinLevel: log.InfoLevel,
		MaxLevel: log.InfoLevel,
		Handler:  log.DiscardHandler(),
	})

	debugLog := log.New()
	debugLog.SetOutput(os.Stdout)
	debugLog.SetFlags(log.LstdFlags)
	debugLog.SetPrefix("")
	debugLog.SetLevel(log.DebugLevel)
	debugLog.SetFormatter(&log.TextFormatter{
		DisableColors:          false,
		DisableTimestamp:       false,
		DisableLevelTruncation: false,
		FullTimestamp:          false,
		TimestampFormat:        time.RFC3339,
		ForceColors:            false,
		ForceQuote:             false,
		DisableSorting:         false,
		DisableQuote:           false,
		QuoteEmptyFields:       false,
		IsTerminal:             false,
		DisableLevelTruncation: false,
		PadLevelText:           false,
		TrimMessages:           false,
		CallerPrettyfier:       nil,
	})
	debugLog.SetReportCaller(false)
	debugLog.SetNoLock()
	debugLog.SetExitFunc(os.Exit)
	debugLog.SetHook(&log.LvlFilterHandler{
		MinLevel: log.DebugLevel,
		MaxLevel: log.DebugLevel,
		Handler:  log.DiscardHandler(),
	})
	debugLog.SetOutput(&log.LvlFilterHandler{
		MinLevel: log.DebugLevel,
		MaxLevel: log.DebugLevel,
		Handler:  log.DiscardHandler(),
	})
	debugLog.SetOutput(&log.LvlFilterHandler{
		MinLevel: log.DebugLevel,
		MaxLevel: log.DebugLevel,
		Handler:  log.DiscardHandler(),
	})

	return &Service{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		DebugLog: debugLog,
	}
}

// ServiceLoggerConfigStdoutWriter is the stdout writer of the service logger config of the service.
type ServiceLoggerConfigStdoutWriter struct {
	// Enabled is the enabled of the stdout writer of the service logger config of the service.
	Enabled bool
}

// ServiceLoggerConfigStderrWriter is the stderr writer of the service logger config of the service.
type ServiceLoggerConfigStderrWriter struct {
	// Enabled is the enabled of the stderr writer of the service logger config of the service.
	Enabled bool
}

// ServiceLoggerConfigFileWriter is the file writer of the service logger config of the service.
type ServiceLoggerConfigFileWriter struct {
	// Enabled is the enabled of the file writer of the service logger config of the service.
	Enabled bool
	// Filename is the filename of the file writer of the service logger config of the service.
	Filename string
	// MaxSize is the max size of the file writer of the service logger config of the service.
	MaxSize int
	// MaxBackups is the max backups of the file writer of the service logger config of the service.
	MaxBackups int
	// MaxAge is the max age of the file writer of the service logger config of the service.
	MaxAge int
	// Compress is the compress of the file writer of the service logger config of the service.
	Compress bool
}

// ServiceLoggerConfig is the logger config of the service.
type ServiceLoggerConfig struct {
	// Level is the level of the service logger config of the service.
	Level string
	// StdoutWriter is the stdout writer of the service logger config of the service.
	StdoutWriter ServiceLoggerConfigStdoutWriter
	// StderrWriter is the stderr writer of the service logger config of the service.
	StderrWriter ServiceLoggerConfigStderrWriter
	// FileWriter is the file writer of the service logger config of the service.
	FileWriter ServiceLoggerConfigFileWriter
}

// ServiceConfig is the config of the service.
type ServiceConfig struct {
	// Logger is the logger of the service config of the service.
	Logger ServiceLoggerConfig
}

// Service is the service.
type Service struct {
	// ErrorLog is the error log of the service.
	ErrorLog *log.Logger
	// InfoLog is the info log of the service.
	InfoLog *log.Logger
	// DebugLog is the debug log of the service.
	DebugLog *log.Logger
}

// NewService is the constructor of the service.
func NewService(config ServiceConfig) *Service {
	service := NewServiceDefault()

	if config.Logger.Level == "error" {
		service.ErrorLog.SetOutput(os.Stderr)
		service.ErrorLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.ErrorLevel,
			MaxLevel: log.ErrorLevel,
			Handler:  log.StreamHandler(os.Stderr, log.TerminalFormat(false)),
		})
		service.InfoLog.SetOutput(os.Stdout)
		service.InfoLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.InfoLevel,
			MaxLevel: log.InfoLevel,
			Handler:  log.StreamHandler(os.Stdout, log.TerminalFormat(false)),
		})
		service.DebugLog.SetOutput(os.Stdout)
		service.DebugLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.DebugLevel,
			MaxLevel: log.DebugLevel,
			Handler:  log.StreamHandler(os.Stdout, log.TerminalFormat(false)),
		})
	} else if config.Logger.Level == "info" {
		service.ErrorLog.SetOutput(os.Stderr)
		service.ErrorLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.ErrorLevel,
			MaxLevel: log.ErrorLevel,
			Handler:  log.StreamHandler(os.Stderr, log.TerminalFormat(false)),
		})
		service.InfoLog.SetOutput(os.Stderr)
		service.InfoLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.InfoLevel,
			MaxLevel: log.InfoLevel,
			Handler:  log.StreamHandler(os.Stderr, log.TerminalFormat(false)),
		})
		service.DebugLog.SetOutput(os.Stdout)
		service.DebugLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.DebugLevel,
			MaxLevel: log.DebugLevel,
			Handler:  log.StreamHandler(os.Stdout, log.TerminalFormat(false)),
		})
	} else if config.Logger.Level == "debug" {
		service.ErrorLog.SetOutput(os.Stderr)
		service.ErrorLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.ErrorLevel,
			MaxLevel: log.ErrorLevel,
			Handler:  log.StreamHandler(os.Stderr, log.TerminalFormat(false)),
		})
		service.InfoLog.SetOutput(os.Stderr)
		service.InfoLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.InfoLevel,
			MaxLevel: log.InfoLevel,
			Handler:  log.StreamHandler(os.Stderr, log.TerminalFormat(false)),
		})
		service.DebugLog.SetOutput(os.Stderr)
		service.DebugLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.DebugLevel,
			MaxLevel: log.DebugLevel,
			Handler:  log.StreamHandler(os.Stderr, log.TerminalFormat(false)),
		})
	}

	if config.Logger.StdoutWriter.Enabled {
		service.ErrorLog.SetOutput(os.Stdout)
		service.ErrorLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.ErrorLevel,
			MaxLevel: log.ErrorLevel,
			Handler:  log.StreamHandler(os.Stdout, log.TerminalFormat(false)),
		})
		service.InfoLog.SetOutput(os.Stdout)
		service.InfoLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.InfoLevel,
			MaxLevel: log.InfoLevel,
			Handler:  log.StreamHandler(os.Stdout, log.TerminalFormat(false)),
		})
		service.DebugLog.SetOutput(os.Stdout)
		service.DebugLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.DebugLevel,
			MaxLevel: log.DebugLevel,
			Handler:  log.StreamHandler(os.Stdout, log.TerminalFormat(false)),
		})
	}

	if config.Logger.StderrWriter.Enabled {
		service.ErrorLog.SetOutput(os.Stderr)
		service.ErrorLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.ErrorLevel,
			MaxLevel: log.ErrorLevel,
			Handler:  log.StreamHandler(os.Stderr, log.TerminalFormat(false)),
		})
		service.InfoLog.SetOutput(os.Stderr)
		service.InfoLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.InfoLevel,
			MaxLevel: log.InfoLevel,
			Handler:  log.StreamHandler(os.Stderr, log.TerminalFormat(false)),
		})
		service.DebugLog.SetOutput(os.Stderr)
		service.DebugLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.DebugLevel,
			MaxLevel: log.DebugLevel,
			Handler:  log.StreamHandler(os.Stderr, log.TerminalFormat(false)),
		})
	}

	if config.Logger.FileWriter.Enabled {
		file, err := os.OpenFile(config.Logger.FileWriter.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}

		service.ErrorLog.SetOutput(file)
		service.ErrorLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.ErrorLevel,
			MaxLevel: log.ErrorLevel,
			Handler:  log.StreamHandler(file, log.TerminalFormat(false)),
		})
		service.InfoLog.SetOutput(file)
		service.InfoLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.InfoLevel,
			MaxLevel: log.InfoLevel,
			Handler:  log.StreamHandler(file, log.TerminalFormat(false)),
		})
		service.DebugLog.SetOutput(file)
		service.DebugLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.DebugLevel,
			MaxLevel: log.DebugLevel,
			Handler:  log.StreamHandler(file, log.TerminalFormat(false)),
		})
	}

	if config.Logger.SyslogWriter.Enabled {
		w, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL0, config.Logger.SyslogWriter.Tag)
		if err != nil {
			log.Fatalf("Failed to open syslog: %v", err)
		}

		service.ErrorLog.SetOutput(w)
		service.ErrorLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.ErrorLevel,
			MaxLevel: log.ErrorLevel,
			Handler:  log.StreamHandler(w, log.TerminalFormat(false)),
		})
		service.InfoLog.SetOutput(w)
		service.InfoLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.InfoLevel,
			MaxLevel: log.InfoLevel,
			Handler:  log.StreamHandler(w, log.TerminalFormat(false)),
		})
		service.DebugLog.SetOutput(w)
		service.DebugLog.SetOutput(&log.LvlFilterHandler{
			MinLevel: log.DebugLevel,
			MaxLevel: log.DebugLevel,
			Handler:  log.StreamHandler(w, log.TerminalFormat(false)),
		})
	}

	return service
}

func (s *Service) Run() {
	s.InfoLog.Info("Starting service")
	s.InfoLog.Info("Service started")
}

func main() {
	service := NewService()
	service.Run()
}

// THIS FILE IS AUTOMATICALLY GENERATED. I AM BEING TAUGHT TO LEARN HOW TO USE GOLANG WITH THE SUPPORT OF AI AGENTS.
// I AM NOT CONFIDENT THAT THIS WILL WORK AS EXPECTED. PLEASE REVIEW AND TEST BEFORE USING IN PRODUCTION..
