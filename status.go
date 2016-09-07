package servicestatus

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
)

type status struct {
	Name      string   `json:"name"`
	Version   string   `json:"version"`
	Hostname  string   `json:"hostname"`
	Timestamp string   `json:"timestamp"`
	Status    string   `json:"status"`
	Checks    []string `json:"checks"`
	Errors    []string `json:"errors"`
	Uptime    string   `json:"uptime"`
	DiskUsage string   `json:"diskUsage"`
}

var startTime time.Time

func init() {
	startTime = time.Now()
}

func diskUsage() string {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs("/", &fs)
	if err != nil {
		return "Unknown"
	}

	all := float64(fs.Blocks * uint64(fs.Bsize))
	free := float64(fs.Bfree * uint64(fs.Bsize))
	used := all - free
	return fmt.Sprintf("%d%%", int((used/all)*100))
}

func uptime() string {
	uptime := time.Since(startTime)
	days := int(uptime.Hours()) / 24
	hours := int(uptime.Hours()) % 24
	minutes := int(uptime.Minutes()) - (hours * 60)
	seconds := int32(uptime.Seconds()) - int32(minutes*60)
	return fmt.Sprintf("%dd:%d:%d:%d", days, hours, minutes, seconds)
}

// Check - to be performed when status info is requested.
type Check func() bool

// ServiceStatus - Provides status for service in TBP format.
type ServiceStatus struct {
	name    string
	version string
	checks  map[string]Check
}

// NewServiceStatus - create a ServiceStatus
func NewServiceStatus(name string, version string) ServiceStatus {
	return ServiceStatus{
		name:    name,
		version: version,
		checks:  make(map[string]Check),
	}
}

// AddCheck - register a check to be performed when status info is requested.
func (ss ServiceStatus) AddCheck(name string, check Check) {
	ss.checks[name] = check
}

// Status - returns a JSON status string
func (ss ServiceStatus) Status() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Unknown"
	}

	statusMsg := "Online"

	checkNames := []string{}
	failedChecks := []string{}

	for checkName, nextCheck := range ss.checks {
		checkNames = append(checkNames, checkName)
		if !nextCheck() {
			failedChecks = append(failedChecks, checkName)
			statusMsg = "Offline"
		}
	}

	status := status{
		Name:      ss.name,
		Version:   ss.version,
		Hostname:  hostname,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Status:    statusMsg,
		Checks:    checkNames,
		Errors:    failedChecks,
		Uptime:    uptime(),
		DiskUsage: diskUsage(),
	}

	b, err := json.MarshalIndent(status, "", "\t")
	if err != nil {
		log.Println(err)
	}
	return string(b)
}
