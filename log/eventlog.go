// +build windows

package log

import (
	"syscall"
	"unsafe"
)

const (
	null = 0

	eventlogErrorType       = 0x0001
	eventlogWarningType     = 0x0002
	eventlogInformationType = 0x0004

	severitySuccess       = 0x0
	severityInformational = 0x1
	severityWarning       = 0x2
	severityError         = 0x3

	customerCode = 1
)

var (
	advapi32            = syscall.NewLazyDLL("advapi32.dll")
	registerEventSource = advapi32.NewProc("RegisterEventSourceW")
	reportEvent         = advapi32.NewProc("ReportEventW")
)

type eventlogWriter struct {
	hEvent   uintptr
	facility uintptr
}

func (w *eventlogWriter) report(wType, severity, code uintptr, m string) error {
	wCategory := uintptr(0)
	dwEventID := severity<<30 | customerCode<<29 | w.facility<<16 | code
	lpString0 := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(m)))
	lpStrings := uintptr(unsafe.Pointer(&lpString0))
	r, _, err := reportEvent.Call(w.hEvent, wType, wCategory, dwEventID, null, 1, 0, lpStrings, null)
	if r == 0 {
		return err
	}
	return nil
}

func (w *eventlogWriter) Info(m string) error {
	return w.report(eventlogInformationType, severityInformational, 0, m)
}

func (w *eventlogWriter) Warning(m string) error {
	return w.report(eventlogWarningType, severityWarning, 0, m)
}

func (w *eventlogWriter) Err(m string) error {
	return w.report(eventlogErrorType, severityError, 0, m)
}

// Open the eventlog writer
func Open(facility Facility, tag string) error {
	hEvent, _, err := registerEventSource.Call(null, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(tag))))
	if hEvent == 0 {
		return err
	}
	theWriter = &eventlogWriter{
		hEvent:   hEvent,
		facility: uintptr(facility),
	}
	return nil
}
