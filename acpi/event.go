package acpi

type EventType = int

const (
	// The lid has been physically opened
	EventLidOpened = iota
	// The lid has been physically closed
	EventLidClosed = iota
	// The ThinkPad has been docked into a UltraDock/UltraBase
	EventDocked = iota
	// The ThinkPad has been docked into a UltraDock/UltraBase
	EventUndocked = iota
	// The power button on the ThinkPad or the Dock has been pressed
	EventButtonPower = iota
)
