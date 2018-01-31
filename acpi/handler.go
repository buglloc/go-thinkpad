package acpi

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

const (
	AcpidSock = "/var/run/acpid.socket"
)

var Debug = false

var AcpiEventsMap = map[string]EventType{
	"button/power PBTN":                       EventButtonPower,
	"button/lid LID open":                     EventLidOpened,
	"button/lid LID close":                    EventLidClosed,
	"ibm/hotkey LEN0068:00 00000080 00004010": EventDocked,
	"ibm/hotkey LEN0268:00 00000080 00004010": EventDocked,
	"ibm/hotkey LEN0068:00 00000080 00004011": EventUndocked,
	"ibm/hotkey LEN0268:00 00000080 00004011": EventUndocked,
}

type HandleMultiEventFn func(eventType EventType) error
type HandleEventFn func() error

func HandleMulti(cb HandleMultiEventFn) error {
	c, err := net.Dial("unix", AcpidSock)
	if err != nil {
		return fmt.Errorf("failed to dial acpid: %s", err.Error())
	}
	defer c.Close()

	reader := bufio.NewReader(c)
	for {
		event, err := reader.ReadString('\n')
		if Debug {
			log.Printf("received ACPI event: %s\n", event)
		}

		if err != nil {
			return fmt.Errorf("failed to read from acpid: %s", err.Error())
		}

		for eventPrefix, eventType := range AcpiEventsMap {
			if strings.HasPrefix(event, eventPrefix) {
				if err := cb(eventType); err != nil {
					return err
				}
				break
			}
		}
	}
}

func Handle(eventType EventType, cb HandleEventFn) error {
	return HandleMulti(func(handledType EventType) error {
		if eventType != handledType {
			return nil
		}

		return cb()
	})
}
