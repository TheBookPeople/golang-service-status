package servicestatus

import (
	"encoding/json"
	"os"
	"testing"
)

func TestServiceStatus(t *testing.T) {

	ss := NewServiceStatus("MyApp", "0.2.1")

	ss.AddCheck("passing_check", "passing check description", func() bool { return true })
	ss.AddCheck("failing_check", "failing check description", func() bool { return false })

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
		actual   interface{}
		expected interface{}
	}{
		{"name", s.Name, "MyApp"},
		{"version", s.Version, "0.2.1"},
		{"hostname", s.Hostname, hostname},
		{"check size", len(s.Checks), 2},
		{"passing check name", s.Checks[0].Name, "passing_check"},
		{"failing check name", s.Checks[1].Name, "failing_check"},
		{"passing check description", s.Checks[0].Description, "passing check description"},
		{"failing check description", s.Checks[1].Description, "failing check description"},
		{"check size", len(s.Errors), 1},
		{"failing check name", s.Errors[0].Name, "failing_check"},
		{"failing check description", s.Errors[0].Description, "failing check description"},
		{"status", s.Status, "Offline"},
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
