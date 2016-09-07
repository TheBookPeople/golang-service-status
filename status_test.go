package servicestatus

import (
	"encoding/json"
	"os"
	"testing"
)

func TestServiceStatus(t *testing.T) {

	ss := NewServiceStatus("MyApp", "0.2.1")

	statusJSON := ss.Status()

	var s status
	err := json.Unmarshal([]byte(statusJSON), &s)
	if err != nil {
		t.Error(err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name     string
		actual   string
		expected string
	}{
		{"name", s.Name, "MyApp"},
		{"version", s.Version, "0.2.1"},
		{"hostname", s.Hostname, hostname},
		{"status", s.Status, "Online"},
		/*{"timestamp", s.Timestamp, "actual"},
		{"uptime", s.Uptime, "actual"},
		{"diskUsage", s.DiskUsage, "actual"},*/
	}

	for _, test := range tests {
		if test.actual != test.expected {
			t.Errorf("Expected %v to be %q but was %q.\n", test.name, test.expected, test.actual)
		}
	}
}
