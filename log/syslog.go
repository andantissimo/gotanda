// +build !windows

package log

import (
	"log/syslog"
)

func (f Facility) toPriority() syslog.Priority {
	switch f {
	case LOCAL0: return syslog.LOG_LOCAL0
	case LOCAL1: return syslog.LOG_LOCAL1
	case LOCAL2: return syslog.LOG_LOCAL2
	case LOCAL3: return syslog.LOG_LOCAL3
	case LOCAL4: return syslog.LOG_LOCAL4
	case LOCAL5: return syslog.LOG_LOCAL5
	case LOCAL6: return syslog.LOG_LOCAL6
	case LOCAL7: return syslog.LOG_LOCAL7
	}
	panic("unexpected facility")
}

// Open the syslog writer
func Open(facility Facility, tag string) error {
	writer, err := syslog.New(facility.toPriority()|syslog.LOG_DEBUG, tag)
	if err != nil {
		return err
	}
	theWriter = writer
	return nil
}
